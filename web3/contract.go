package main

import (
	"context"
	"log"
	"strings"
	"time"
	"web3/contract/defi"
	"web3/contract/nft"
	"web3/contract/store"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ethUri        = "wss://sepolia.infura.io/ws/v3/b0daa49c16d7466cbdf68176ba2a243a"
	aaveAddress   = common.HexToAddress("0x316B1198D72c782Dd44440fCfB6d03Fc98b5c160")
	nftAddress    = common.HexToAddress("0xb8cD5FC286922c54AB32A8AaBF583E382b44F050")
	storeAddress  = common.HexToAddress("0xfeadcf82070998D19A215C91E19638Bfcd1Ab854")
	walletAddress = common.HexToAddress("0xe3D3a9a1111872990e0f5a1351D7876162A40Fa6")
)

type Contract struct {
	eventBus   *cqrs.EventBus
	commandBus *cqrs.CommandBus
	store      *store.Store
	nft        *nft.Nft
	aave       *defi.Defi
	client     *ethclient.Client
}

func (c *Contract) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	nonce, err := c.client.PendingNonceAt(ctx, address)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func (c *Contract) Watch(ctx context.Context) error {
	// ──────────────── 1. NFT 事件监听 ────────────────
	nftCh := make(chan types.Log)
	nftAbi, _ := abi.JSON(strings.NewReader(nft.NftABI))
	nftSub, err := c.client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{nftAddress},
	}, nftCh)
	if err != nil {
		return err
	}
	// ──────────────── 2. Store 事件监听 ────────────────
	itemSetChan := make(chan *store.StoreItemSet)
	isSub, err := c.store.WatchItemSet(&bind.WatchOpts{Context: ctx}, itemSetChan)
	if err != nil {
		return err
	}
	// ──────────────── 3. Aave 事件监听（新增）───────────────
	aaveCh := make(chan types.Log)
	aaveAbi, _ := abi.JSON(strings.NewReader(defi.DefiABI))
	aaveSub, err := c.client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{aaveAddress},
	}, aaveCh)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			/// ---- NFT 事件 ----
			case err := <-nftSub.Err():
				log.Fatal(err)
			case nc := <-nftCh:
				event, err := nftAbi.EventByID(nc.Topics[0])
				if err != nil {
					log.Println("Unknown event:", nc.Topics[0].Hex())
					continue
				}
				eventName := event.Name
				ev := EventLog{
					TxHash:          nc.TxHash.Hex(),
					LogIndex:        int(nc.Index),
					BlockNumber:     int64(nc.BlockNumber),
					BlockTime:       time.Now(), // 建议通过区块时间获取并填充
					ContractAddress: nc.Address.Hex(),
					EventSignature:  nc.Topics[0].Hex(),
					EventName:       eventName,
					CreatedAt:       time.Now(),
				}
				if len(nc.Topics) > 0 {
					ev.Topic0 = nc.Topics[0].Hex()
				}
				if len(nc.Topics) > 1 {
					ev.Topic1 = nc.Topics[1].Hex()
				}
				if len(nc.Topics) > 2 {
					ev.Topic2 = nc.Topics[2].Hex()
				}
				if len(nc.Topics) > 3 {
					ev.Topic3 = nc.Topics[3].Hex()
				}
				ev.Data = nc.Data
				c.eventBus.Publish(ctx, ev)
			/// ---- Aave 事件 ----
			case err := <-aaveSub.Err():
				log.Fatal("aaveSub err:", err)
			case ac := <-aaveCh:
				event, err := aaveAbi.EventByID(ac.Topics[0])
				if err != nil {
					log.Println("Unknown Aave event:", ac.Topics[0].Hex())
					continue
				}
				eventName := event.Name
				log.Println("Aave event:", eventName)
			/// ---- Store 事件 ----
			case ev := <-itemSetChan:
				cmd := ItemSetCommand{}
				cmd.Key = ev.Key
				cmd.Value = ev.Value
				c.commandBus.Send(ctx, cmd)
			case err := <-isSub.Err():
				log.Fatal("isSub err:", err)
			}
		}
	}()
	return nil
}

func NewContract(routers *Routers) *Contract {
	c := Contract{}
	client, err := ethclient.Dial(ethUri)
	if err != nil {
		panic(err)
	}
	store, err := store.NewStore(storeAddress, client)
	if err != nil {
		panic(err)
	}
	nft, err := nft.NewNft(nftAddress, client)
	if err != nil {
		panic(err)
	}
	aave, err := defi.NewDefi(aaveAddress, client)
	if err != nil {
		panic(err)
	}
	c.aave = aave
	c.nft = nft
	c.store = store
	c.eventBus = routers.EventBus
	c.commandBus = routers.CommandBus
	c.client = client
	return &c
}

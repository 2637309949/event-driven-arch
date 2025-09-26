package main

import (
	"context"
	"log"
	"strings"
	"time"
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
	ethUri       = "wss://sepolia.infura.io/ws/v3/b0daa49c16d7466cbdf68176ba2a243a"
	nftAddress   = common.HexToAddress("0xb8cD5FC286922c54AB32A8AaBF583E382b44F050")
	storeAddress = common.HexToAddress("0xfeadcf82070998D19A215C91E19638Bfcd1Ab854")
)

type Contract struct {
	eventBus   *cqrs.EventBus
	commandBus *cqrs.CommandBus
	store      *store.Store
	nft        *nft.Nft
	client     *ethclient.Client
}

func (c *Contract) Watch(ctx context.Context) error {
	nftCh := make(chan types.Log)
	contractAbi, _ := abi.JSON(strings.NewReader(nft.NftABI))
	nftSub, err := c.client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{nftAddress},
	}, nftCh)
	if err != nil {
		return err
	}
	itemSetChan := make(chan *store.StoreItemSet)
	isSub, err := c.store.WatchItemSet(&bind.WatchOpts{Context: ctx}, itemSetChan)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case err := <-nftSub.Err():
				log.Fatal(err)
			case nc := <-nftCh:
				event, err := contractAbi.EventByID(nc.Topics[0])
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
				ev.Data = string(nc.Data)
				c.eventBus.Publish(ctx, ev)
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

func NewContract(eventBus *cqrs.EventBus, commandBus *cqrs.CommandBus) *Contract {
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
	c.nft = nft
	c.store = store
	c.eventBus = eventBus
	c.commandBus = commandBus
	c.client = client
	return &c
}

package main

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"
	"web3/contract/store"
	"web3/contract/swap"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// https://sepolia.etherscan.io/
	ethUri        = "wss://sepolia.infura.io/ws/v3/b0daa49c16d7466cbdf68176ba2a243a"
	swapAddress   = common.HexToAddress("0x44506214dc0ABBbF07cBa4b38689A2aBfaF82861")
	storeAddress  = common.HexToAddress("0xfeadcf82070998D19A215C91E19638Bfcd1Ab854")
	walletAddress = common.HexToAddress("0xe3D3a9a1111872990e0f5a1351D7876162A40Fa6")
)

type Contract struct {
	eventBus   *cqrs.EventBus
	commandBus *cqrs.CommandBus
	store      *store.Store
	swap       *swap.Swap
	client     *ethclient.Client
}

// 二分法定位区块高度
func (c *Contract) GetBlockByTime(ctx context.Context, targetTime int64) (uint64, error) {
	latest, err := c.client.BlockByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}

	low := uint64(0)
	high := latest.NumberU64()
	var blockNum uint64

	for low <= high {
		mid := (low + high) / 2
		block, err := c.client.BlockByNumber(ctx, big.NewInt(int64(mid)))
		if err != nil {
			return 0, err
		}

		if int64(block.Time()) < targetTime {
			low = mid + 1
		} else {
			blockNum = mid
			if mid == 0 {
				break
			}
			high = mid - 1
		}
	}
	return blockNum, nil
}

// 查询时间段内指定合约事件日志
func (c *Contract) FilterLogsByTime(ctx context.Context, contractAddr common.Address, startTime, endTime int64) ([]types.Log, error) {
	var startBlock, endBlock *big.Int
	var err error

	if startTime != 0 {
		b, err := c.GetBlockByTime(ctx, startTime)
		if err != nil {
			return nil, err
		}
		startBlock = big.NewInt(int64(b))
	}
	if endTime != 0 {
		b, err := c.GetBlockByTime(ctx, endTime)
		if err != nil {
			return nil, err
		}
		endBlock = big.NewInt(int64(b))
	}

	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []common.Address{contractAddr},
	}

	logs, err := c.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// 实时监控
func (c *Contract) Watch(ctx context.Context) {
	// ──────────────── 1. Swap 事件监听 ────────────────
	swapCh := make(chan types.Log)
	swapAbi, _ := abi.JSON(strings.NewReader(swap.SwapABI))
	swapSub, err := c.client.SubscribeFilterLogs(ctx, ethereum.FilterQuery{
		Addresses: []common.Address{swapAddress},
	}, swapCh)
	if err != nil {
		panic(err)
	}
	// ──────────────── 2. Store 事件监听 ────────────────
	itemSetChan := make(chan *store.StoreItemSet)
	itemSub, err := c.store.WatchItemSet(&bind.WatchOpts{Context: ctx}, itemSetChan)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			/// ---- Swap 事件 ----
			case err := <-swapSub.Err():
				log.Fatal("Watch.swapSub fatal:", err)
			case nc := <-swapCh:
				ev, err := c.ParseEventLog(ctx, swapAbi, nc)
				if err != nil {
					log.Println("Unknown event:", nc.Topics[0].Hex())
					continue
				}
				c.eventBus.Publish(ctx, ev)
			/// ---- Store 事件 ----
			case ev := <-itemSetChan:
				cmd := ItemSetCommand{}
				cmd.Key = ev.Key
				cmd.Value = ev.Value
				c.commandBus.Send(ctx, cmd)
			case err := <-itemSub.Err():
				log.Fatal("Watch.itemSub fatal:", err)
			}
		}
	}()
}

// 解析日志
func (c *Contract) ParseEventLog(ctx context.Context, abi abi.ABI, nc types.Log) (EventLog, error) {
	ev := EventLog{}
	event, err := abi.EventByID(nc.Topics[0])
	if err != nil {
		log.Println("Unknown event:", nc.Topics[0].Hex())
		return ev, err
	}
	eventName := event.Name
	ev.TxHash = nc.TxHash.Hex()
	ev.LogIndex = int(nc.Index)
	ev.BlockNumber = int64(nc.BlockNumber)
	ev.BlockTime = time.Now()
	ev.ContractAddress = nc.Address.Hex()
	ev.EventSignature = nc.Topics[0].Hex()
	ev.EventName = eventName
	ev.CreatedAt = time.Now()
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
	return ev, nil
}

// 查找交易Nonce
func (c *Contract) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	nonce, err := c.client.PendingNonceAt(ctx, address)
	if err != nil {
		return 0, err
	}
	return nonce, nil
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
	swap, err := swap.NewSwap(swapAddress, client)
	if err != nil {
		panic(err)
	}
	c.swap = swap
	c.store = store
	c.eventBus = routers.EventBus
	c.commandBus = routers.CommandBus
	c.client = client
	return &c
}

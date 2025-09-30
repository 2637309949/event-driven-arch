package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/robfig/cron"
)

type Timed struct {
	eventBus *cqrs.EventBus
	contract *Contract
	repo     *Repository
	cr       *cron.Cron
}

// 模拟交易
func (t *Timed) TimedFlushTest() {
	er := EventOccurred{
		EventName: "PendingNonce",
		Occurred:  time.Now(),
	}
	err := t.eventBus.Publish(ctx, &er)
	if err != nil {
		log.Println("eventBus.Publish err:", err)
		return
	}
}

func (t *Timed) TimedFlushStat() {
	// 刷新交易总数
	nonce, err := t.contract.PendingNonceAt(ctx, walletAddress)
	if err != nil {
		fmt.Println(err)
	}
	ep := EventStats{
		EventLabel: "交易总数",
		EventName:  "PendingNonce",
		EventCount: int(nonce),
	}
	err = t.repo.UpdateEventStats(ctx, &ep)
	if err != nil {
		fmt.Println(err)
		return
	}
	//刷新活跃用户
	au, err := t.repo.QueryActiveUsers(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	ep = EventStats{
		EventLabel: "活跃用户",
		EventName:  "ActiveUsers",
		EventCount: int(au),
	}
	err = t.repo.UpdateEventStats(ctx, &ep)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (t *Timed) Start(ctx context.Context) {
	// t.cr.AddFunc("0/2 * * * * ?", t.TimedFlushTest)  // 每2秒钟刷新一次
	t.cr.AddFunc("0 0/10 * * * ?", t.TimedFlushStat) // 每10分钟刷新一次stat
	go t.cr.Start()
}

func NewTimed(cr *cron.Cron, contract *Contract, repo *Repository, eventBus *cqrs.EventBus) *Timed {
	c := Timed{}
	c.eventBus = eventBus
	c.contract = contract
	c.repo = repo
	c.cr = cr
	return &c
}

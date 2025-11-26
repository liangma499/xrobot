package wallet

import (
	"context"
	"xbase/eventbus"
	"xbase/log"
	"xbase/utils/xconv"
)

const (
	withdrawWaterTopic = "tron:wallet:withdraw:water" // 提现流水
	balanceChangeTopic = "tron:wallet:balance:change" // 钱包余额变动
)

// PublishBalanceChange 发布钱包余额变动的事件
func PublishBalanceChange(payload *BalanceChangePayload) {
	err := eventbus.Publish(context.Background(), balanceChangeTopic, payload)
	if err != nil {
		log.Errorf("pub event failed, payload = %s err = %v", xconv.Json(payload), err)
	}
}

// SubscribeBalanceChange 订阅钱包余额变动的事件
func SubscribeBalanceChange(handler func(uuid string, payload *BalanceChangePayload)) {
	err := eventbus.Subscribe(context.Background(), balanceChangeTopic, func(event *eventbus.Event) {
		payload := &BalanceChangePayload{}

		err := event.Payload.Scan(payload)
		if err != nil {
			return
		}

		handler(event.ID, payload)
	})
	if err != nil {
		log.Errorf("subscribe event failed: %v", err)
	}
}

// PublishWithdrawWater 用户提现是成功发布
func PublishWithdrawWater(payload *WithdrawWaterEventPayload) {
	err := eventbus.Publish(context.Background(), withdrawWaterTopic, payload)
	if err != nil {
		log.Errorf("pub event failed, payload = %s err = %v", xconv.Json(payload), err)
	}
}

// SubscribeWithdrawWater 订阅钱包余额变动的事件
func SubscribeWithdrawWater(handler func(uuid string, payload *WithdrawWaterEventPayload)) {
	err := eventbus.Subscribe(context.Background(), withdrawWaterTopic, func(event *eventbus.Event) {
		payload := &WithdrawWaterEventPayload{}

		err := event.Payload.Scan(payload)
		if err != nil {
			return
		}

		handler(event.ID, payload)
	})
	if err != nil {
		log.Errorf("subscribe event failed: %v", err)
	}
}

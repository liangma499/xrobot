package cryptocurrencyevent

import (
	"context"
	"xbase/eventbus"
	"xbase/log"
	"xbase/utils/xconv"
)

// PublishCryptoCurrency 发布有充值成功用户
func PublishCryptoCurrency(payload *CryptoCurrencyMsg) {
	err := eventbus.Publish(context.Background(), cryptocurrencySuccessTopic, payload)
	if err != nil {
		log.Errorf("pub event failed, payload = %s err = %v", xconv.Json(payload), err)
	}
}

// SubscribeCryptoCurrency 订阅有充值成功用户
func SubscribeCryptoCurrency(handler func(uuid string, payload *CryptoCurrencyMsg)) {
	err := eventbus.Subscribe(context.Background(), cryptocurrencySuccessTopic, func(event *eventbus.Event) {
		payload := &CryptoCurrencyMsg{}

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

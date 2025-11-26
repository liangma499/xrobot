package user

import (
	"context"
	"xbase/eventbus"
	"xbase/log"
	"xbase/utils/xconv"
)

const (
	registerTopic   = "user:register"   // 用户注册成功
	loginTopic      = "user:login"      // 用户登录成功
	infoChangeTopic = "user:infoChange" // 用户信息变动
	offlineTopic    = "user:offline"    // 用户下线事件
)

// PublishRegister 发布注册事件
func PublishRegister(payload *RegisterPayload) {
	err := eventbus.Publish(context.Background(), registerTopic, payload)
	if err != nil {
		log.Errorf("pub event failed, payload = %s err = %v", xconv.Json(payload), err)
	}
}

// SubscribeRegister 订阅注册事件
func SubscribeRegister(handler func(payload *RegisterPayload)) {
	err := eventbus.Subscribe(context.Background(), registerTopic, func(event *eventbus.Event) {
		payload := &RegisterPayload{}

		err := event.Payload.Scan(payload)
		if err != nil {
			return
		}

		handler(payload)
	})
	if err != nil {
		log.Errorf("subscribe event failed: %v", err)
	}
}

// PublishLogin 发布登录事件
func PublishLogin(payload *LoginPayload) {
	err := eventbus.Publish(context.Background(), loginTopic, payload)
	if err != nil {
		log.Errorf("publish event failed, payload = %s err = %v", xconv.Json(payload), err)
	}
}

// SubscribeLogin 订阅登录事件
func SubscribeLogin(handler func(payload *LoginPayload)) {
	err := eventbus.Subscribe(context.Background(), loginTopic, func(event *eventbus.Event) {
		payload := &LoginPayload{}

		err := event.Payload.Scan(payload)
		if err != nil {
			return
		}

		handler(payload)
	})
	if err != nil {
		log.Errorf("subscribe event failed: %v", err)
	}
}

// PublishInfoChange 发布用户信息变动事件
func PublishInfoChange(uid int64) {
	err := eventbus.Publish(context.Background(), infoChangeTopic, uid)
	if err != nil {
		log.Errorf("publish user's info change event failed, uid = %d err = %v", uid, err)
	}
}

// SubscribeInfoChange 订阅用户信息变动事件
func SubscribeInfoChange(handler func(uid int64)) {
	err := eventbus.Subscribe(context.Background(), infoChangeTopic, func(event *eventbus.Event) {
		handler(event.Payload.Int64())
	})
	if err != nil {
		log.Errorf("subscribe user's info change event failed: %v", err)
	}
}

// PublishOffline 发布用户下线事件
func PublishOffline(uid int64) {
	err := eventbus.Publish(context.Background(), offlineTopic, uid)
	if err != nil {
		log.Errorf("publish user offline event failed, uid = %d err = %v", uid, err)
	}
}

// SubscribeOffline 订阅用户下线事件
func SubscribeOffline(handler func(uid int64)) {
	err := eventbus.Subscribe(context.Background(), offlineTopic, func(event *eventbus.Event) {
		handler(event.Payload.Int64())
	})
	if err != nil {
		log.Errorf("subscribe user offline event failed: %v", err)
	}
}

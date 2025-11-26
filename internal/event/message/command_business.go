package message

import (
	"context"
	"xbase/eventbus"
	"xbase/log"
	"xbase/utils/xconv"
)

type MessageBusiness struct {
	MessageCommon
	UserName   string `json:"userName"`   //用户名
	InviteCode string `json:"inviteCode"` //邀请码
	FirstName  string `json:"firstName"`  //姓
	LastName   string `json:"lastName"`   //名
}

func (mb *MessageBusiness) Clone() *MessageBusiness {
	if mb == nil {
		return nil
	}
	return &MessageBusiness{
		MessageCommon: mb.MessageCommon.Clone(),
		UserName:      mb.UserName,   //用户名
		InviteCode:    mb.InviteCode, //邀请码
		FirstName:     mb.FirstName,  //姓
		LastName:      mb.LastName,   //名
	}
}

// PublishMessageStart 发布用户Start命令
func PublishMessageBusiness(payload *MessageBusiness) {
	err := eventbus.Publish(context.Background(), MessageBusinessTopic, payload)
	if err != nil {
		log.Errorf("pub event failed, payload = %s err = %v", xconv.Json(payload), err)
	}
}

// SubscribeMessageStart 订阅读发布用户Start命令
func SubscribeMessageBusiness(handler func(uuid string, payload *MessageBusiness)) {
	err := eventbus.Subscribe(context.Background(), MessageBusinessTopic, func(event *eventbus.Event) {
		payload := &MessageBusiness{}

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

package message

import (
	"context"
	"xbase/eventbus"
	"xbase/log"
	"xbase/utils/xconv"
)

type MessageStart struct {
	MessageCommon
	UserName   string `json:"userName"`   //用户名
	InviteCode string `json:"inviteCode"` //邀请码
	FirstName  string `json:"firstName"`  //姓
	LastName   string `json:"lastName"`   //名
}

func (mb *MessageStart) Clone() *MessageStart {
	if mb == nil {
		return nil
	}
	return &MessageStart{
		MessageCommon: mb.MessageCommon.Clone(),
		UserName:      mb.UserName,   //用户名
		InviteCode:    mb.InviteCode, //邀请码
		FirstName:     mb.FirstName,  //姓
		LastName:      mb.LastName,   //名
	}
}

// PublishMessageStart 发布用户Start命令
func PublishMessageStart(payload *MessageStart) {
	err := eventbus.Publish(context.Background(), messageStartTopic, payload)
	if err != nil {
		log.Errorf("pub event failed, payload = %s err = %v", xconv.Json(payload), err)
	}
}

// SubscribeMessageStart 订阅读发布用户Start命令
func SubscribeMessageStart(handler func(uuid string, payload *MessageStart)) {
	err := eventbus.Subscribe(context.Background(), messageStartTopic, func(event *eventbus.Event) {
		payload := &MessageStart{}

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

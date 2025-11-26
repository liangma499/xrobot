package message

import (
	tgtypes "xrobot/internal/xtelegram/tg-types"
)

const (
	messageStartTopic    = "message:command:start" // 钱包余额变动
	MessageBusinessTopic = "message:command:business"
)

type MessageType string

const (
	MessageType_Private    MessageType = "private"    //私聊
	MessageType_Group      MessageType = "group"      //普通群组
	MessageType_Supergroup MessageType = "supergroup" //超级群组
	MessageType_Channel    MessageType = "channel"    //频道
)

// 公共字段
type MessageCommon struct {
	Button          tgtypes.XTelegramButton `json:"button"`
	Type            MessageType             `json:"type"`            // 聊天的类型，可以是 “private”（私聊）、“group”（普通群组）、“supergroup”（超级群组）或 “channel”（频道）。
	ChatID          int64                   `json:"chatID"`          //聊天ID
	UserID          int64                   `json:"userID"`          //用户Telegram ID
	ChannelCode     string                  `json:"channelCode"`     //渠道码
	ClientIP        string                  `json:"clientIP"`        //前端IP
	ChatInstance    string                  `json:"chatInstance"`    //消息号
	WaitforinputMsg WaitforinputMsg         `json:"waitforinputMsg"` //处理输入型消息
	OrderID         string                  `json:"orderID"`
}

func (mc MessageCommon) Clone() MessageCommon {

	return MessageCommon{
		Button:          mc.Button,
		Type:            mc.Type,
		ChatID:          mc.ChatID,
		UserID:          mc.UserID,
		ChannelCode:     mc.ChannelCode,
		ClientIP:        mc.ClientIP,
		ChatInstance:    mc.ChatInstance,
		WaitforinputMsg: mc.WaitforinputMsg.Clone(),
		OrderID:         mc.OrderID,
	}
}

type WaitforinputMsg struct {
	InPutMsg bool                  `json:"inPutMsg"`
	PutMsg   string                `json:"putMsg,omitempty"` //接收到的消息
	Extended *WaitforinputExtended `json:"extended,omitempty"`
}

func (we WaitforinputMsg) Clone() WaitforinputMsg {

	return WaitforinputMsg{
		InPutMsg: we.InPutMsg,
		PutMsg:   we.PutMsg, //接收到的消息
		Extended: we.Extended.Clone(),
	}
}

type WaitforinputExtended struct {
	LastCallbackData string `json:"lastCallbackData,omitempty"` //过期时间
}

func (we *WaitforinputExtended) Clone() *WaitforinputExtended {
	if we == nil {
		return nil
	}
	return &WaitforinputExtended{
		LastCallbackData: we.LastCallbackData,
	}
}

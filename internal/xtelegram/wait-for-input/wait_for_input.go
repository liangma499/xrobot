package waitforinput

import (
	"context"
	"encoding/json"
	"fmt"
	rediswaitForInput "tron_robot/internal/component/redis/redis-wait-for-input"
	"tron_robot/internal/event/message"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"tron_robot/internal/xtypes"
	"xbase/log"
	"xbase/utils/xconv"
)

type WaitforinputInfo struct {
	UserID        int64                         `json:"-"` //用户TG
	InPutMsg      bool                          `json:"-"`
	Button        tgtypes.XTelegramButton       `json:"-"`
	UserBottonKey string                        `json:"userBottonKey,omitempty"`
	Extended      *message.WaitforinputExtended `json:"extended,omitempty"`
}

func (wipi *WaitforinputInfo) Clone() *WaitforinputInfo {
	if wipi == nil {
		return nil
	}
	return &WaitforinputInfo{
		UserID:        wipi.UserID,
		UserBottonKey: wipi.UserBottonKey,
		Extended:      wipi.Extended.Clone(),
	}
}

func SetWaitForinputKey(wipi *WaitforinputInfo) {
	if wipi == nil {
		return
	}
	//保存最后操作的回调信息
	if wipi.Button.IsSaveLastButton() {
		Instance().SetLastCallBack(wipi.UserID, wipi.Button.CallbackData())
	}

	// 保存需要输入的信息
	key := fmt.Sprintf(xtypes.CacheWaitForInputKey, wipi.UserID)
	if wipi.InPutMsg {
		if err := rediswaitForInput.Instance().Del(context.Background(), key).Err(); err != nil {
			log.Errorf("del err:%v,tgUserID:%s", err, wipi.UserID)
		}
		return
	}

	if wipi.UserBottonKey == "" {
		if err := rediswaitForInput.Instance().Del(context.Background(), key).Err(); err != nil {
			log.Errorf("del err:%v,tgUserID:%s", err, wipi.UserID)
		}
		return
	}
	rst, err := json.Marshal(wipi)
	if err != nil {
		log.Errorf("set err:%v,tgUserID:%d", err, wipi.UserID)
	}
	if err := rediswaitForInput.Instance().Set(context.Background(), key, xconv.String(rst), xtypes.CacheWaitForInputExpiration).Err(); err != nil {
		log.Errorf("set err:%v,tgUserID:%d,userKey:%s", err, wipi.UserID, xconv.String(rst))
	}
}

func GetWaitForinputKey(userID int64) *WaitforinputInfo {
	key := fmt.Sprintf(xtypes.CacheWaitForInputKey, userID)
	res, err := rediswaitForInput.Instance().Get(context.Background(), key).Bytes()
	if err != nil {
		return nil
	}
	if res != nil {
		if len(res) > 0 {
			rst := &WaitforinputInfo{}
			if err = json.Unmarshal(res, rst); err != nil {
				log.Errorf("get err:%v,userKey:%s", err, userID)
			} else {
				return rst
			}

		}

	}
	return nil
}

func DelWaitForinputKey(tgUserID int64) {
	key := fmt.Sprintf(xtypes.CacheWaitForInputKey, tgUserID)
	if err := rediswaitForInput.Instance().Del(context.Background(), key).Err(); err != nil {
		log.Errorf("del err:%v,tgUserID:%d", err, tgUserID)
	}
}

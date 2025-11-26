package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"tron_robot/internal/code"
	redisdefault "tron_robot/internal/component/redis/redis-default"
	"xbase/errors"
)

const (
	thirdGameameUid = "thirdGame:user:%s:currency%s:key:%s" // 三方地址
	expiration      = 60 * time.Minute
)

type ThirdGameInfo struct {
	CurrencyID   int      `json:"currency_id" redis:"currency_id"`
	Currency     string   `json:"currency" redis:"currency"`
	GameID       int64    `json:"game_id" redis:"game_id"`
	GameName     string   `json:"game_name" redis:"game_name"`
	LobbyUrl     string   `json:"lobbyUrl" redis:"lobbyUrl"`
	UCode        string   `json:"u_code" redis:"u_code"`
	UID          int64    `json:"uid" redis:"uid"`
	PGameID      string   `json:"p_gameID" redis:"p_gameID"`
	ProviderCode string   `json:"provider_code" redis:"provider_code"`
	ProviderIp   []string `json:"provider_ip" redis:"provider_ip"` // 供应商白名单IP
	CheckIp      int32    `json:"check_ip" redis:"check_ip"`       // 1强制检查IP 2不检查
	UserType     int32    `json:"user_type" redis:"user_type"`
}

// 三方游戏币种设置
func SetThirdGameInfo(uCode, md5Key string, data *ThirdGameInfo) error {
	currency := strings.ToUpper(data.Currency)
	key := fmt.Sprintf(thirdGameameUid, uCode, currency, md5Key)
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = redisdefault.Instance().Set(context.TODO(), key, string(byteData), expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
func GetThirdGameInfo(uCode, md5Key, currency string) (*ThirdGameInfo, error) {
	currency = strings.ToUpper(currency)
	key := fmt.Sprintf(thirdGameameUid, uCode, currency, md5Key)

	str := redisdefault.Instance().Get(context.TODO(), key).Val()
	if str == "" {
		return nil, errors.NewError(code.AccountNotExist)
	}
	data := &ThirdGameInfo{}
	err := json.Unmarshal([]byte(str), data)
	if err != nil {
		return nil, errors.NewError(code.AccountNotExist)
	}
	redisdefault.Instance().Expire(context.TODO(), key, expiration)
	return data, err
}

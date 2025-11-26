package waitforinput_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"tron_robot/internal/xtypes"
	"xbase/log"

	"github.com/go-redis/redis/v8"
)

const saveLastButtonScript = `
	if #KEYS ~= 2 then
	    return {"-1"}
	end
	local redisKey = KEYS[1]
	local cmd = KEYS[2]
	
	if cmd == "get" then
		local lastCallbackData = ""
		local exist = redis.call('EXISTS', redisKey)
		if exist ~= 0 then
			lastCallbackData = redis.call('GET', redisKey)
		end
		return {"1", lastCallbackData}
	elseif cmd == "set" then
		redis.call('SET', redisKey,ARGV[1])
		-- 设置key 一天过期
		redis.call('Expire', redisKey, 86400)
		return {"1"}
	end
`

type saveLastButton struct {
	saveLastButtonScript *redis.Script
	redisClient          redis.UniversalClient
}

var (
	once     sync.Once
	instance *saveLastButton
)

func Instance() *saveLastButton {
	once.Do(func() {
		instance = &saveLastButton{
			saveLastButtonScript: redis.NewScript(saveLastButtonScript),
			redisClient: redis.NewUniversalClient(&redis.UniversalOptions{
				Addrs:      []string{"127.0.0.1:6379"},
				DB:         1,
				Username:   "",
				Password:   "",
				MaxRetries: 3,
			}),
		}
	})

	return instance
}

func (slb saveLastButton) SetLastCallBack(userID int64, callbackData string) bool {
	key := fmt.Sprintf(xtypes.UserLastButtonKey, userID)
	rst, err := slb.saveLastButtonScript.Run(context.Background(), slb.redisClient, []string{key, "set"}, callbackData).StringSlice()
	if err != nil {
		log.Errorf("get err:%v,user:%v,callbackData:%v", err, userID, callbackData)
		return false
	}
	lenght := len(rst)
	if lenght == 0 {
		log.Warnf("set rst is nil user:%v,callbackData:%v", userID, callbackData)
		return false
	}
	if lenght == 1 {
		return rst[0] == "1"
	}
	return false
}
func (slb saveLastButton) GetLastCallBack(userID int64) string {
	key := fmt.Sprintf(xtypes.UserLastButtonKey, userID)
	rst, err := slb.saveLastButtonScript.Run(context.Background(), slb.redisClient, []string{key, "get"}).StringSlice()
	if err != nil {
		log.Errorf("get err:%v,user:%v", err, userID)
		return ""
	}
	lenght := len(rst)
	if lenght == 0 {
		log.Warnf("get rst is nil user:%v,callbackData:%v", userID)
		return ""
	}
	if lenght < 2 {
		log.Warnf("get rst:%v,user:%v,callbackData:%v", rst[0], userID)
		return ""
	}
	return rst[1]
}

func TestClient_ResdisSet(t *testing.T) {

	Instance().SetLastCallBack(1, "cccccccccccccccc")

	key := Instance().GetLastCallBack(1)
	log.Warnf("%v", key)

}
func TestClient_GetOrderID(t *testing.T) {
	callData := "XB_EFR_ROABP1890447709966434304"
	button := tgtypes.StringToXTelegramButton(callData)

	log.Warnf("%s", button.GetOrderID(callData))

}

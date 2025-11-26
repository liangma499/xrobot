package waitforinput

import (
	"context"
	"fmt"
	"sync"
	rediswaitForInput "tron_robot/internal/component/redis/redis-wait-for-input"
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
}

var (
	once     sync.Once
	instance *saveLastButton
)

func Instance() *saveLastButton {
	once.Do(func() {
		instance = &saveLastButton{
			saveLastButtonScript: redis.NewScript(saveLastButtonScript),
		}
	})

	return instance
}

func (slb saveLastButton) SetLastCallBack(userID int64, callbackData string) bool {
	key := fmt.Sprintf(xtypes.UserLastButtonKey, userID)
	rst, err := slb.saveLastButtonScript.Run(context.Background(), rediswaitForInput.Instance(), []string{key, "set"}, callbackData).StringSlice()
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
	rst, err := slb.saveLastButtonScript.Run(context.Background(), rediswaitForInput.Instance(), []string{key, "get"}).StringSlice()
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

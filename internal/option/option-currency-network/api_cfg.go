package optionCurrencyNetworkCfg

import (
	"context"
	"fmt"
	"sync"
	"time"
	redispayment "tron_robot/internal/component/redis/redis-crypto-currencies"
	"tron_robot/internal/model"
	"tron_robot/internal/xtypes"
	"xbase/log"
	"xbase/utils/xconv"
	"xbase/utils/xrand"
	"xbase/utils/xtime"

	"github.com/go-redis/redis/v8"
)

type apiConfig struct {
	mux sync.Mutex
}

var (
	onceApi  sync.Once
	instance *apiConfig
)

func Instance() *apiConfig {
	onceApi.Do(func() {
		instance = &apiConfig{}

	})
	return instance
}
func (dao *apiConfig) doDurationKindRedisKey(dk xtypes.DurationKind, apiKind xtypes.APIKind) string {
	switch dk {
	case xtypes.DurationKind_Daily: // 日榜
		{
			nowStr := xtime.Now().Format("20060102")
			return fmt.Sprintf(xtypes.ApiRankCfgRedisKey, dk, apiKind, nowStr)
		}
	case xtypes.DurationKind_Monthly: // 月榜
		{
			nowStr := xtime.Now().Format("200601")
			return fmt.Sprintf(xtypes.ApiRankCfgRedisKey, dk, apiKind, nowStr)
		}
	}
	return fmt.Sprintf(xtypes.ApiRankCfgRedisKey, dk, apiKind, "none")
}
func (dao *apiConfig) doDuration(dk xtypes.DurationKind) int64 {
	switch dk {
	case xtypes.DurationKind_Daily: // 日榜
		{
			return xtime.DayHead(1).Unix() - xtime.Now().Unix() + xrand.Int64(600, 3600)
		}
	case xtypes.DurationKind_Monthly: // 月榜
		{
			return xtime.MonthHead(1).Unix() - xtime.Now().Unix() + xrand.Int64(1200, 7200)
		}
	}
	return -1
}
func (dao *apiConfig) InitApiCfg(platformCfg *model.OptionCurrencyNetwork) error {

	if platformCfg == nil {
		return nil
	}
	ctx := context.Background()
	for key, item := range platformCfg.ApiCfg {
		redisKey := dao.doDurationKindRedisKey(item.DurationKind, key)

		//没有这个Key 把所有的初始化的redis中

		addMembers := make([]*redis.Z, 0)
		for k := range item.Cfg {
			addMembers = append(addMembers, &redis.Z{
				Score:  0,
				Member: k,
			})
		}
		redispayment.Instance().ZAddNX(ctx, redisKey, addMembers...)
		durationTime := dao.doDuration(item.DurationKind)
		if durationTime > 0 {
			redispayment.Instance().Expire(ctx, redisKey, time.Duration(durationTime)*time.Second)
		}
	}

	return nil
}
func (dao *apiConfig) doApiCfg(channelType xtypes.NetWorkChannelType, APIKind xtypes.APIKind) (*xtypes.ApiCfg, xtypes.KeyToApiCfgInfo) {
	ctx := context.Background()
	channelcfg := GetOptByChannelType(channelType)
	if channelcfg == nil {
		return nil, nil
	}
	if channelcfg.ApiCfg == nil {
		return nil, nil
	}

	cfg, ok := channelcfg.ApiCfg[APIKind]
	if !ok {
		return nil, nil
	}
	if cfg == nil {
		return nil, nil
	}

	cfgMap := cfg.Cfg.Clone()

	redisKey := dao.doDurationKindRedisKey(cfg.DurationKind, APIKind)

	list, err := redispayment.Instance().ZRangeWithScores(ctx, redisKey, 0, 1000).Result()
	if err != nil {
		log.Errorf("fetch score amount failed, key = %s start = %d stop = %d err = %v", redisKey, 0, 100, err)
		return nil, nil
	}
	if len(list) == 0 && len(cfgMap) > 0 {
		dao.InitApiCfg(channelcfg)
		for key, item := range cfgMap {
			redispayment.Instance().ZIncrBy(ctx, redisKey, 1, key)
			return item, cfgMap
		}
	} else {
		for _, item := range list {
			member := xconv.String(item.Member)
			if info, ok := cfgMap[member]; ok && info != nil {
				if info.DurationMax > 0 && info.DurationMax > int64(item.Score) {
					redispayment.Instance().ZRem(ctx, redisKey, member)
					continue
				}

				redispayment.Instance().ZIncrBy(ctx, redisKey, 1, member)
				return info, cfgMap
			} else {
				redispayment.Instance().ZRem(ctx, redisKey, member)
			}

		}
	}

	return nil, nil
}
func (dao *apiConfig) GetApiByChannelType(channelType xtypes.NetWorkChannelType, APIKind xtypes.APIKind) (*xtypes.ApiCfg, xtypes.KeyToApiCfgInfo) {
	dao.mux.Lock()
	defer dao.mux.Unlock()

	rst, apiCfgMap := dao.doApiCfg(xtypes.NetWorkChannelType_TRON, APIKind)
	return rst, apiCfgMap
}

// 主要用于验证
func (dao *apiConfig) GetNeedApiByChannelType(channelType xtypes.NetWorkChannelType, APIKind xtypes.APIKind) *xtypes.ApiCfg {
	dao.mux.Lock()
	defer dao.mux.Unlock()

	rst, _ := dao.doApiCfg(channelType, APIKind)
	return rst
}

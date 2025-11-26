package xtypes

import "time"

const (
	CacheUserBaseKey          = "baseuser:%d"          // 缓存用户KEY
	CacheUserBaseCodeToUIDKey = "baseuser:code:%s:uid" // 用户编号对应的用户ID
)

const (
	CacheWaitForInputKey        = "waitForInput:%d:code" // 用户等待输入按钮
	CacheWaitForInputExpiration = 1 * time.Hour          // 过期时间
)

const (
	UserLastButtonKey = "userLastButton:%d:code" // 用户最后用的button
)

const (
	UniqueAmountKey = "uniqueAmount:%s:currency"
)

const (
	ApiRankCfgRedisKey = "rankApiCfg:%d:kind:%d:api:%s:date"
)
const (
	CacheApiCfgKey        = "apiCfg:%d:kind"    // api类型
	CacheApiCfgExpiration = 15 * 24 * time.Hour // 过期时间
)
const (
	ListenerAddrKey = "listenerAddr:%s:network" // 监听用户地址
)
const (
	CacheChannelCfgKey        = "channelCfg:%s:code" // 渠道编号
	CacheChannelCfgExpiration = 15 * 24 * time.Hour  // 过期时间
)
const (
	CacheNetworkAddresToCfgKey         = "networkAddresToCfgKey:%s:network:%s"           // api类型
	CacheNetworkChannelCodeToCfgKey    = "networkChannelCodeToCfg:%s:network:%s:kind:%d" // api类型
	CacheNetworkAddressToCfgExpiration = 15 * 24 * time.Hour                             // 过期时间
)
const (
	ListenerTrxIDKey = "listenerTrxID:%s:network" // 监听交易Hash
)

const (
	LastMessgeIDKey        = "lastMessgeIDKey:%d:%s:msg" // 监听交易Hash
	LastMessgeIDExpiration = 15 * time.Minute
)

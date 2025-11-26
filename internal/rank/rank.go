package rank

import (
	"context"
	"fmt"
	"math"
	"xbase/codes"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xconv"
	"xbase/utils/xtime"
	redisdefault "xrobot/internal/component/redis/redis-default"

	"github.com/go-redis/redis/v8"
)

// 增加榜单积分脚本
const incrScript = `
	local cmdRun = KEYS[#KEYS]
	if not cmdRun then
		return {}
	end
	if not (cmdRun == 'ZINCRBY' or cmdRun == 'ZADD') then
		return {}
	end
	for i = 1, #KEYS - 1 do
		if redis.call('EXISTS', KEYS[i]) == 1 then
			redis.call(cmdRun, KEYS[i], ARGV[1], ARGV[2])
		else
			redis.call(cmdRun, KEYS[i], ARGV[1], ARGV[2])
			local ttl = tonumber(ARGV[i+2])
			if ttl > 0 then
				redis.call('EXPIRE', KEYS[i], ttl)
			end
		end
	end
	return {}
`

const (
	dailyRankKey   = "daily:rank:%s"   // zset; 日榜
	weeklyRankKey  = "weekly:rank:%s"  // zset; 周榜
	monthlyRankKey = "monthly:rank:%s" // zset; 月榜
	overallRankKey = "overall:rank"    // zset; 总榜
)

const (
	DailyRank   Kind = 1 // 日榜
	WeeklyRank  Kind = 2 // 周榜
	MonthlyRank Kind = 3 // 月榜
	OverallRank Kind = 4 // 总榜
)

const (
	Asc  Order = 1 // 从低到高
	Desc Order = 2 // 从高到低
)

// Kind 排行榜类型
type Kind int

// Order 排序方式
type Order int

type Rank struct {
	opts   *Options
	redis  redis.UniversalClient
	script *redis.Script
}

type rank struct {
	rank   *Rank
	prefix string
}

type Item struct {
	Rank   int     `json:"rank"`
	Member string  `json:"member"`
	Score  float64 `json:"score"`
}

// FetchRankListArgs 拉取排行榜（参数）
type FetchRankListArgs struct {
	Kind   Kind  // 榜单类型
	Offset int   // 榜单偏移
	Page   int   // 当前页
	Limit  int   // 每页大小
	Order  Order // 排序方式
}

// FetchRankListRst 拉取排行榜（结果）
type FetchRankListRst struct {
	StartTime int64   // 开始时间
	EndTime   int64   // 结束时间
	HasMore   bool    // 是否还有更多
	List      []*Item // 榜单列表
}

// GetMemberRankArgs 获取成员排名（参数）
type GetMemberRankArgs struct {
	Kind   Kind   // 榜单类型
	Offset int    // 榜单偏移
	Order  Order  // 排序方式
	Member string // 成员
}

// GetMemberRankRst 获取成员排名（结果）
type GetMemberRankRst struct {
	Rank  int64
	Score float64
}
type DelMemberRankArgs struct {
	Kind   Kind   // 榜单类型
	Offset int    // 榜单偏移
	Member string // 成员
}

type Options struct {
	Prefix string        // key前缀模板
	Cycles map[Kind]uint // 榜单周期
}

func NewRank(opts *Options) *Rank {
	return &Rank{
		opts:   opts,
		redis:  redisdefault.Instance(),
		script: redis.NewScript(incrScript),
	}
}

// Prefix 生成前缀
func (r *Rank) Prefix(args ...any) *rank {
	return &rank{
		rank:   r,
		prefix: fmt.Sprintf(r.opts.Prefix, args...),
	}
}

// FetchRankList 拉取排行榜（不携带前缀key）
func (r *Rank) FetchRankList(ctx context.Context, args *FetchRankListArgs) (*FetchRankListRst, error) {
	return (&rank{rank: r}).FetchRankList(ctx, args)
}

// Incr 增加榜单积分
func (r *rank) incr(ctx context.Context, member string, score float64, bUpdate ...bool) error {
	if len(r.rank.opts.Cycles) == 0 {
		return nil
	}

	var (
		now    = xtime.Now()
		keys   = make([]string, 0, len(r.rank.opts.Cycles))
		values = make([]any, 0, 2+len(r.rank.opts.Cycles))
	)

	values = append(values, score, member)
	for kind, cycle := range r.rank.opts.Cycles {
		switch kind {
		case DailyRank:
			if cycle > 0 {
				label := now.Format("20060102")
				if r.prefix != "" {
					keys = append(keys, r.prefix+":"+fmt.Sprintf(dailyRankKey, label))
				} else {
					keys = append(keys, fmt.Sprintf(dailyRankKey, label))
				}
				values = append(values, cycle*24*3600)
			}
		case WeeklyRank:
			if cycle > 0 {
				year, week := now.ISOWeek()
				label := fmt.Sprintf("%d%d", year, week)
				if r.prefix != "" {
					keys = append(keys, r.prefix+":"+fmt.Sprintf(weeklyRankKey, label))
				} else {
					keys = append(keys, fmt.Sprintf(weeklyRankKey, label))
				}
				values = append(values, cycle*7*24*3600)
			}
		case MonthlyRank:
			if cycle > 0 {
				label := now.Format("200601")
				if r.prefix != "" {
					keys = append(keys, r.prefix+":"+fmt.Sprintf(monthlyRankKey, label))
				} else {
					keys = append(keys, fmt.Sprintf(monthlyRankKey, label))
				}
				values = append(values, cycle*31*24*3600)
			}
		case OverallRank:
			if r.prefix != "" {
				keys = append(keys, r.prefix+":"+overallRankKey)
			} else {
				keys = append(keys, overallRankKey)
			}
			values = append(values, -1)
		}
	}

	//if cmd ~= 'ZINCRBY' or cmd ~= 'ZADD' then
	if len(bUpdate) > 0 && bUpdate[0] {
		//更新数据
		keys = append(keys, "ZADD")
	} else {
		//增加数据
		keys = append(keys, "ZINCRBY")
	}
	err := r.rank.script.Run(ctx, r.rank.redis, keys, values...).Err()
	if err != nil {
		log.Errorf("incr score failed, prefix = %s member = %s score = %f err = %v", r.prefix, member, score, err)
		return errors.NewError(err, codes.InternalError)
	}
	return nil
}

// Incr 增加榜单积分
func (r *rank) Incr(ctx context.Context, member string, score float64) error {
	return r.incr(ctx, member, score)
}

// Decr 扣减榜单积分
func (r *rank) Decr(ctx context.Context, member string, score float64) error {
	return r.incr(ctx, member, 0-score)
}

// Update 更新榜单积分
func (r *rank) Update(ctx context.Context, member string, score float64) error {
	return r.incr(ctx, member, score, true)
}

// FetchRankList 拉取排行榜
func (r *rank) FetchRankList(ctx context.Context, args *FetchRankListArgs) (*FetchRankListRst, error) {
	var (
		err   error
		key   string
		list  []redis.Z
		start = int64((args.Page - 1) * args.Limit)
		stop  = start + int64(args.Limit) + 1
		rst   = &FetchRankListRst{}
	)

	switch args.Kind {
	case DailyRank: // 日榜
		dayHead := xtime.DayHead(args.Offset)
		label := dayHead.Format("20060102")
		rst.StartTime = dayHead.Unix()
		rst.EndTime = xtime.DayTail(args.Offset).Unix()
		if r.prefix != "" {
			key = r.prefix + ":" + fmt.Sprintf(dailyRankKey, label)
		} else {
			key = fmt.Sprintf(dailyRankKey, label)
		}
	case WeeklyRank: // 周榜
		weekHead := xtime.WeekHead(args.Offset)
		year, week := weekHead.ISOWeek()
		label := fmt.Sprintf("%d%d", year, week)
		rst.StartTime = weekHead.Unix()
		rst.EndTime = xtime.WeekTail(args.Offset).Unix()
		if r.prefix != "" {
			key = r.prefix + ":" + fmt.Sprintf(weeklyRankKey, label)
		} else {
			key = fmt.Sprintf(weeklyRankKey, label)
		}
	case MonthlyRank: // 月榜
		monthHead := xtime.MonthHead(args.Offset)
		label := monthHead.Format("200601")
		rst.StartTime = monthHead.Unix()
		rst.EndTime = xtime.MonthTail(args.Offset).Unix()
		if r.prefix != "" {
			key = r.prefix + ":" + fmt.Sprintf(monthlyRankKey, label)
		} else {
			key = fmt.Sprintf(monthlyRankKey, label)
		}
	default: // 总榜
		if r.prefix != "" {
			key = r.prefix + ":" + overallRankKey
		} else {
			key = overallRankKey
		}
	}

	if args.Order == Asc {
		list, err = r.rank.redis.ZRangeWithScores(ctx, key, start, stop).Result()
	} else {
		list, err = r.rank.redis.ZRevRangeWithScores(ctx, key, start, stop).Result()
	}
	if err != nil {
		log.Errorf("fetch score amount failed, key = %s start = %d stop = %d err = %v", key, start, stop, err)
		return nil, errors.NewError(err, codes.InternalError)
	}

	rst.HasMore = len(list) > args.Limit
	rst.List = make([]*Item, 0, int(math.Min(float64(len(list)), float64(args.Limit))))
	for i, item := range list {
		rst.List = append(rst.List, &Item{
			Rank:   int(start) + i + 1,
			Member: xconv.String(item.Member),
			Score:  item.Score,
		})

		if len(rst.List) == args.Limit {
			break
		}
	}

	return rst, nil
}

// GetMemberRank 获取成员排名
func (r *rank) GetMemberRank(ctx context.Context, args *GetMemberRankArgs) (*GetMemberRankRst, error) {
	var (
		err error
		key string
		val int64
	)

	switch args.Kind {
	case DailyRank: // 日榜
		label := xtime.DayHead(args.Offset).Format("20060102")
		if r.prefix != "" {
			key = r.prefix + ":" + fmt.Sprintf(dailyRankKey, label)
		} else {
			key = fmt.Sprintf(dailyRankKey, label)
		}
	case WeeklyRank: // 周榜
		year, week := xtime.WeekHead(args.Offset).ISOWeek()
		label := fmt.Sprintf("%d%d", year, week)
		if r.prefix != "" {
			key = r.prefix + ":" + fmt.Sprintf(weeklyRankKey, label)
		} else {
			key = fmt.Sprintf(weeklyRankKey, label)
		}
	case MonthlyRank: // 月榜
		label := xtime.MonthHead(args.Offset).Format("200601")
		if r.prefix != "" {
			key = r.prefix + ":" + fmt.Sprintf(monthlyRankKey, label)
		} else {
			key = fmt.Sprintf(monthlyRankKey, label)
		}
	default: // 总榜
		if r.prefix != "" {
			key = r.prefix + ":" + overallRankKey
		} else {
			key = overallRankKey
		}
	}

	if args.Order == Asc {
		val, err = r.rank.redis.ZRank(ctx, key, args.Member).Result()
	} else {
		val, err = r.rank.redis.ZRevRank(ctx, key, args.Member).Result()
	}
	if err != nil {
		if err != redis.Nil {
			log.Errorf("get member's rank failed, key = %s member = %s err = %v", key, args.Member, err)
			return nil, errors.NewError(err, codes.InternalError)
		}

		return &GetMemberRankRst{}, nil
	}

	score, err := r.rank.redis.ZScore(ctx, key, args.Member).Result()
	if err != nil && err != redis.Nil {
		log.Errorf("get member's score failed, key = %s member = %s err = %v", key, args.Member, err)
		return nil, errors.NewError(err, codes.InternalError)
	}

	return &GetMemberRankRst{
		Rank:  val + 1,
		Score: score,
	}, nil
}

// DelMemberRank 删除成员排名
func (r *rank) DelMemberRank(ctx context.Context, args *DelMemberRankArgs) error {
	var (
		err error
		key string
	)

	switch args.Kind {
	case DailyRank: // 日榜
		label := xtime.DayHead(args.Offset).Format("20060102")
		if r.prefix != "" {
			key = r.prefix + ":" + fmt.Sprintf(dailyRankKey, label)
		} else {
			key = fmt.Sprintf(dailyRankKey, label)
		}
	case WeeklyRank: // 周榜
		year, week := xtime.WeekHead(args.Offset).ISOWeek()
		label := fmt.Sprintf("%d%d", year, week)
		if r.prefix != "" {
			key = r.prefix + ":" + fmt.Sprintf(weeklyRankKey, label)
		} else {
			key = fmt.Sprintf(weeklyRankKey, label)
		}
	case MonthlyRank: // 月榜
		label := xtime.MonthHead(args.Offset).Format("200601")
		if r.prefix != "" {
			key = r.prefix + ":" + fmt.Sprintf(monthlyRankKey, label)
		} else {
			key = fmt.Sprintf(monthlyRankKey, label)
		}
	default: // 总榜
		if r.prefix != "" {
			key = r.prefix + ":" + overallRankKey
		} else {
			key = overallRankKey
		}
	}

	_, err = r.rank.redis.ZRem(ctx, key, args.Member).Result()
	if err != nil && err != redis.Nil {

		return errors.NewError(err, codes.InternalError)
	}

	return nil
}

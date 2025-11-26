package rank_test

import (
	"context"
	"fmt"
	"testing"
	"xbase/utils/xconv"
	ranks "xrobot/internal/rank"
)

var rank *ranks.Rank

func init() {
	rank = ranks.NewRank(&ranks.Options{
		Prefix: "test:%d:recharge",
		Cycles: map[ranks.Kind]uint{
			ranks.DailyRank:   3, // 保存3天的数据
			ranks.WeeklyRank:  3, // 保存3周的数据
			ranks.MonthlyRank: 3, // 保存3月的数据
			ranks.OverallRank: 0, // 永久保存
		},
	})
}

func TestRank_Incr(t *testing.T) {
	r := rank.Prefix(12)

	for i := 0; i < 20; i++ {
		err := r.Update(context.Background(), xconv.String(i), 200.0)
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Log("OK")
}

func TestRank_FetchRankList(t *testing.T) {
	r := rank.Prefix(11)

	rst, err := r.FetchRankList(context.Background(), &ranks.FetchRankListArgs{
		Kind:  ranks.DailyRank,
		Page:  1,
		Limit: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(xconv.Json(rst))
}

func TestRank_GetMemberRank(t *testing.T) {
	r := rank.Prefix(11)

	rst, err := r.GetMemberRank(context.Background(), &ranks.GetMemberRankArgs{
		Kind:   ranks.DailyRank,
		Member: "39",
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(xconv.Json(rst))
}

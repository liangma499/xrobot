package rate_test

import (
	"testing"
	"time"
	"tron_robot/internal/service/wallet/rate"
	"xbase/log"
	"xbase/utils/xtime"
)

func TestClient_Timer(t *testing.T) {
	for {
		tmr := time.After(10 * time.Second)
		<-tmr
		log.Warnf("%v", xtime.Now())
	}
}

func TestClient_rate(t *testing.T) {
	rate.Instance().InitTimer()
	select {}
}

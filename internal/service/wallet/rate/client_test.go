package rate_test

import (
	"testing"
	"time"
	"xbase/log"
	"xbase/utils/xtime"
	"xrobot/internal/service/wallet/rate"
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

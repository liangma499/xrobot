package walletsdk_test

import (
	"testing"
	walletsdk "xrobot/tools/wallet/wallet-sdk"
)

func TestClient_GetBalance(t *testing.T) {
	u, err := walletsdk.NewClient().GetSlot()

	t.Logf("u:%v err:%v", u, err)

	data, err := walletsdk.NewClient().GetBlock(u)

	t.Logf("data:%v err:%v", data, err)
}

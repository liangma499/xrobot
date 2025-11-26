package basic_test

import (
	"context"
	"testing"
	"xbase/config"
	"xbase/config/file"
	"xbase/utils/xconv"

	"xbase/registry/consul"
	"xbase/transport/rpcx"
	"xrobot/internal/service/basic"
	basepb "xrobot/internal/service/basic/pb"
)

var transporter = rpcx.NewTransporter(
	rpcx.WithClientDiscovery(consul.NewRegistry()),
)

func init() {
	// 设置配置中心
	config.SetConfigurator(config.NewConfigurator(config.WithSources(file.NewSource(file.WithPath("E:/workspace/src/u2bet_server/mesh/config")))))
}
func TestClient_SendEmailCode(t *testing.T) {
	client, err := basic.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	reply, err := client.SendEmailCode(context.Background(), &basepb.SendEmailCodeArgs{
		Email: "libasi167@gmail.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(xconv.Json(reply))
}

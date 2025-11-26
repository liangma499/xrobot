package cryptocurrency_test

import (
	"context"
	"testing"
	"xbase/registry/consul"
	"xbase/transport/rpcx"
	cryptocurrencysvc "xrobot/internal/service/cryptocurrency"
	cryptocurrencypb "xrobot/internal/service/cryptocurrency/pb"
)

var transporter = rpcx.NewTransporter(
	rpcx.WithClientDiscovery(consul.NewRegistry()),
)

func TestClient_CryptoCurrency(t *testing.T) {
	client, err := cryptocurrencysvc.NewClient(transporter.NewClient)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.CryptoCurrency(context.Background(), &cryptocurrencypb.CryptoCurrencyArgs{
		UID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

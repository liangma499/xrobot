package xeth_test

import (
	"context"
	"testing"
	"xrobot/internal/utils/xeth"
)

func Test_HeaderNumber(t *testing.T) {
	number, err := xeth.HeaderNumber(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(number)

	hash, err := xeth.HeaderHash(context.Background(), number+10)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(hash)
}

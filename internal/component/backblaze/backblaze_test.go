package backblaze_test

import (
	"context"
	"testing"
	"xrobot/internal/component/backblaze"
)

var ins = backblaze.NewInstance(&backblaze.Config{
	Account:  "ad9fc922950a",
	Key:      "005ae24f2177b8c8ab870d528fd0128464ff77e6ec",
	Bucket:   "xrobot-test",
	BasePath: "avatar/",
})

func TestBackBlaze_PutObject(t *testing.T) {
	object, err := ins.PutObject(context.Background(), "admin.png")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
	t.Log(object)
}

func TestBackBlaze_DeleteObject(t *testing.T) {
	err := ins.DeleteObject(context.Background(), "1745659011564306432.png")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}

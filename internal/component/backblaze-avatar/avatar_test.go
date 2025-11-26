package backblazeAvatar_test

import (
	"context"
	"testing"
	backblazeAvatar "tron_robot/internal/component/backblaze-avatar"
	"tron_robot/internal/xtelegram/telegram/types"
)

var ins = backblazeAvatar.NewInstance(&backblazeAvatar.Config{
	BasePath:    "../avatar/",
	BasePathUrl: "avatar/",
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
	err := ins.DeleteObject(context.Background(), "1788769319069417472.png")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("OK")
}
func Test_FileTgAvarl(t *testing.T) {
	data, err := ins.PutFileTgAvarl(&types.File{
		FilePath: "photos/file_4.jpg",
	}, "6867997452:AAFYZXHAC_TDvcfBiYto2ShutRSiUcboa04")

	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

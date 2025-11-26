package main_test

import (
	"testing"
	"xbase/log"
	"xrobot/internal/utils/xresource"
)

func TestClient_Config(t *testing.T) {
	str := xresource.ToResourceUrl("aaaaaaaaaaaa")
	log.Warnf(str)
}

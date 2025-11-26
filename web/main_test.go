package main_test

import (
	"testing"
	"tron_robot/internal/utils/xresource"
	"xbase/log"
)

func TestClient_Config(t *testing.T) {
	str := xresource.ToResourceUrl("aaaaaaaaaaaa")
	log.Warnf(str)
}

package xresource

import (
	"fmt"
	"os"
	"strings"
	optionbaseconfig "tron_robot/internal/option/option-base-config"
	"tron_robot/internal/xtypes"
	"xbase/utils/xrand"
	"xbase/utils/xvalidate"
)

// ToAvatarUrl 转换成头像地址
func ToAvatarUrl(avatar string) string {
	if xvalidate.IsDigit(avatar) {
		return avatar
	}

	if avatar == "" {
		return avatar
	}
	if strings.Contains(avatar, "http://") || strings.Contains(avatar, "https://") {
		return avatar
	}
	baseUrl := optionbaseconfig.GetValue(xtypes.BaseUrlKey)

	if baseUrl == "" {
		return avatar
	}

	return baseUrl + "/" + strings.TrimPrefix(avatar, "/")

}

// ToResourceUrl 转换为资源地址
func ToResourceUrl(image string) string {
	if image == "" {
		return image
	}
	if strings.Contains(image, "http://") || strings.Contains(image, "https://") {
		return image
	}
	baseUrl := optionbaseconfig.GetValue(xtypes.BaseUrlKey)

	if baseUrl == "" {
		return image
	}

	return baseUrl + "/" + strings.TrimPrefix(image, "/")
}

func RandAvatarUrl(avatarCount int32) string {

	return fmt.Sprintf("/resource/upload/avatar/%d.png", xrand.Int32(1, avatarCount))
}
func InviteCodeUrl(userCode, channelCode string) string {
	url := optionbaseconfig.GetValue(xtypes.InviteCodeUrlKey)
	if url == "" {
		return url
	}
	replaces := make(map[string]string)

	replaces["code"] = userCode
	replaces["channelCode"] = channelCode
	body := os.Expand(url, func(s string) string {
		return replaces[s]
	})
	return body
}

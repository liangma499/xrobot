package xtypes

const (
	BaseUrlKey              = "base_url"
	AvatarUrlKey            = "avatar_url"
	InviteCodeUrlKey        = "invite_code_url"
	WebLobbyUrlKey          = "web_lobby_url"
	TelegramWebhookUrlKey   = "telegram_webhook_url"
	TelegramApiUrlKey       = "telegram_api_url"
	CommissionRatioKey      = "commission_ratio"       // 佣金比例
	CommissionCreateCardKey = "commission_create_card" // 开卡奖励
	PointsMineRatioKey      = "points_mine_ratio"      // 自己充值的积分比例
	PointsChildRatioKey     = "points_child_ratio"     // 子级充值的积分比例
	ServerStatusKey         = "server_status"          // 服务器状态
	WhiteListKey            = "white_list"             // 维护白名单
)

type ServerStatus string

const (
	Opened   ServerStatus = "1" //正常
	Maintain ServerStatus = "2" //维护
	Closed   ServerStatus = "3" //关闭
	Error    ServerStatus = "4" //错误
)

func (s ServerStatus) String() string {
	return string(s)
}

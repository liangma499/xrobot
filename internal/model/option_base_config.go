package model

import (
	"time"
	"tron_robot/internal/xtypes"
)

// 卡片配置表 用户类型
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionBaseConfig -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionBaseConfig struct {
	Key         string              `gorm:"column:card_type;primarykey;size:64;comment:配置Key"`
	Value       string              `gorm:"column:value;type:text;not null;comment:卡名"` //
	Memo        string              `gorm:"column:memo;size:512;comment:说明"`            //
	Status      xtypes.OptionStatus `gorm:"column:status;size:32;comment:状态( 1启用,2禁用)"`
	OperateUid  int64               `gorm:"column:operate_uid;size:64;comment:操作用户ID"`      //
	OperateUser string              `gorm:"column:operate_user;size:64;comment:操作用户名"`      //
	CreateAt    time.Time           `gorm:"column:created_at;type:timestamp;comment:创建时间戳"` //
	UpdateAt    time.Time           `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"` //
}

// `gorm:"column:login_at;size:64"`
func (c *OptionBaseConfig) TableName() string {
	return "option_base_config"
}

/*
EnergyPricesU:   decimal.NewFromFloat(65000),  //有U交易需要的能量
EnergyPricesNou: decimal.NewFromFloat(131000), //无U交易需要的能量
TrxPriceU:       decimal.NewFromFloat(3),      //有U的转账价格
TrxPriceNoU:     decimal.NewFromFloat(6),      //无U的转账价格
*/

/*
用户ID: 5619143861
用户昵称: TRX能量供应
闪兑利润: 4%
闪租供价: 2.4 TRX
笔数供价: 3.3 TRX
星星供价: 0.017U/个
会员供价: 14.5U/3月、21U/6月、37U/1年
靓号分成：40%
推广分成: 120U/机器人
我的下级：0/直推 0/间推
已用笔数: 0 笔
剩余笔数: 0 笔
当前余额：291.03 TRX + 4.5 USDT

🅰️直推收益
🔸机器人：30U
🔸闪租：0.1 TRX/笔
🔸笔数：0.1 TRX/笔
🔸闪兑：1%
🔸靓号：10%
🔸星星：0.0005U/个
🔸会员：0.5U/3月、1U/6月、2U/1年

🅱️间推收益
🔹机器人：20U
🔹闪租：0.1 TRX/笔
🔹笔数：0.1 TRX/笔
🔹闪兑：0.5%
🔹靓号：5%
🔹星星：0.0005U/个
🔹会员：0.5U/3月、0.5U/6月、1U/1年
*/

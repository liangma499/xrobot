package model

import (
	"time"
	"tron_robot/internal/xtypes"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=OptionCurrencyNetwork -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type OptionCurrencyNetwork struct {
	ID            int64                     `gorm:"column:id;size:64;primarykey;autoIncrement;comment:自增ID"` //序号
	Type          xtypes.NetWorkChannelType `gorm:"column:type;unique" json:"type"`                          //类型 1:EVM 2:Solana 3:Ton
	NetWork       xtypes.NetWork            `gorm:"column:net_work;primarykey;size:32" json:"netWork"`       //充值渠道名
	ApiCfg        xtypes.ApiCfgList         `gorm:"column:api_cfg;type:json" json:"api_cfg"`                 //连接地址
	PrivateKey    xtypes.EncryptByte        `gorm:"column:private_key;type:blob" json:"privateKey"`          //加密数据
	DecimalPlaces int8                      `gorm:"column:decimal_places;size:8" json:"decimal_places"`      //加密数据
	Status        xtypes.OptionStatus       `gorm:"column:status;size:8" json:"status"`                      //状态（1=启用;2=停用）
	CreateAt      time.Time                 `gorm:"column:create_at;autoCreateTime" json:"-"`                //创建时间
	UpdateAt      time.Time                 `gorm:"column:update_at;autoUpdateTime" json:"-"`                //更新时间

}

// TableName 表名称
func (*OptionCurrencyNetwork) TableName() string {
	return "option_currency_network"
}
func (cp *OptionCurrencyNetwork) Clone() *OptionCurrencyNetwork {
	if cp == nil {
		return nil
	}
	return &OptionCurrencyNetwork{
		ID:            cp.ID,                 //序号
		Type:          cp.Type,               //类型 1:EVM 2:Solana 3:Ton
		NetWork:       cp.NetWork,            //充值渠道名
		ApiCfg:        cp.ApiCfg.Clone(),     //连接地址
		PrivateKey:    cp.PrivateKey.Clone(), //加密数据
		DecimalPlaces: cp.DecimalPlaces,
		Status:        cp.Status,   //状态（1=启用;2=停用）
		CreateAt:      cp.CreateAt, //创建时间
		UpdateAt:      cp.UpdateAt, //更新时间
	}

}

package model

import (
	"time"
	"tron_robot/internal/xtypes"

	"github.com/shopspring/decimal"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=PaymentCryptoTransaction -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type PaymentCryptoTransaction struct {
	ID              int64                    `gorm:"column:id;size:64;primarykey;autoIncrement"` // 主键
	ChannelCode     string                   `gorm:"column:channel_code;index;size:100;comment:渠道名称" json:"channelCode"`
	AddressKind     xtypes.AddressKind       `gorm:"column:address_kind;size:32;comment:地址类型" json:"addressKind"`
	NetWork         xtypes.NetWork           `gorm:"column:network;size:32;uniqueIndex:uinque_net_tid;index:idx_ns;comment:网络"`
	TransactionHash string                   `gorm:"column:transaction_hash;size:128;uniqueIndex:uinque_net_tid;comment:交易hash"`
	BlockHash       string                   `gorm:"column:block_hash;size:128;comment:区块hash"`
	BlockNum        int64                    `gorm:"column:block_num;size:64;comment:区块编号"`
	Protocol        string                   `gorm:"column:protocol;size:128;comment:协议名"`
	ToAddress       string                   `gorm:"column:to_address;size:128;comment:收款地址"`
	FromAddress     string                   `gorm:"column:from_address;size:128;comment:转出地址"`
	Amount          decimal.Decimal          `gorm:"column:amount;type:decimal(32,9);default:0;comment:金额"`
	RealAmount      decimal.Decimal          `gorm:"column:real_amount;type:decimal(32,9);default:0;comment:原始金额"`
	Contract        string                   `gorm:"column:contract;size:128;comment:合约名 trx 为-"`
	Currency        xtypes.Currency          `gorm:"column:currency;size:32;comment:币种名"`
	EnergyFee       int64                    `gorm:"column:energy_fee;size:64;default:0;comment:gasFree"`
	NetUsage        int64                    `gorm:"column:net_usage;size:64;default:0;comment:net_usage"`
	EnergyUsage     int64                    `gorm:"column:energy_usage;size:64;default:0;comment:energy_usage"`
	VerifyCount     int32                    `gorm:"column:verify_count;size:32;default:0;comment:验证次数"`
	TransactionKind xtypes.TransactionKind   `gorm:"column:transaction_kind;size:64;default:0;comment: 0无 1充值 2提现 3能量转入 4能量转出"`
	Stauts          xtypes.TransactionStatus `gorm:"column:stauts;size:64;default:1;index:idx_ns;comment:1待验证 2交易成功 3交易失败 4完成"`
	CreateAt        time.Time                `gorm:"column:created_at;type:timestamp;comment:创建时间戳" json:"-"`
	UpdateAt        time.Time                `gorm:"column:updated_at;type:timestamp;comment:修改时间戳" json:"-"`
}

func (u *PaymentCryptoTransaction) TableName() string {
	return "payment_crypto_transaction"
}

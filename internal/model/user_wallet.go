package model

import (
	"xrobot/internal/xtypes"

	"github.com/shopspring/decimal"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=UserWallet -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserWallet struct {
	ID               int64                      `gorm:"column:id;size:64;primarykey;autoIncrement;comment:自增ID"`                               // 主键
	UID              int64                      `gorm:"column:uid;uniqueIndex:uin_uid_currency;size:64" json:"uid" redis:"uid"`                // 用户ID
	Currency         xtypes.Currency            `gorm:"column:currency;uniqueIndex:uin_uid_currency;size:32" json:"currency" redis:"currency"` // 币种
	WalletAmountKind xtypes.WalletAmountKind    `gorm:"column:wallet_amount_kind;uniqueIndex:uin_uid_currency;size:32" json:"walletAmountKind" redis:"walletAmountKind"`
	Cash             decimal.Decimal            `gorm:"column:cash;type:decimal(32,10);default:0;" json:"cash" redis:"cash"` // 现金
	Used             decimal.Decimal            `gorm:"column:used;type:decimal(32,10);default:0;" json:"used" redis:"used"` // 奖金
	Def              xtypes.WalletDefaultStatus `gorm:"column:def;size:8;default:0" json:"def" redis:"def"`                  // 是否默认
}

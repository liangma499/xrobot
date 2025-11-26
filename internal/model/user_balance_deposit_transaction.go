package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// 用户交易记录
//
//go:generate xgorm-dao-generator -model-dir=. -model-names=UserBalanceDepositTransaction -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type UserBalanceDepositTransaction struct {
	ID                     int64           `gorm:"column:id;size:64;primarykey;autoIncrement;"`          // 主键
	ThirdUID               string          `gorm:"column:third_uid;size:64;index:index_thirduid_cardid"` // 三方用户ID
	Symbol                 string          `gorm:"column:symbol;size:16;comment:USDT USDC"`              // 卡号
	Network                string          `gorm:"column:net_work;size:32;Polygon,Tron"`                 // 币种
	Amount                 decimal.Decimal `gorm:"column:amount;type:decimal(32,4)"`                     // amount
	DayZeroTime            int64           `gorm:"column:day_zero_time;size:64;index;"`                  //记录的当天
	DepositTransactionTime string          `gorm:"column:deposit_transaction_time;size:64"`              //交易时间
	CreateAt               time.Time       `gorm:"column:created_at;type:timestamp;comment:创建时间戳"`
	UpdateAt               time.Time       `gorm:"column:updated_at;type:timestamp;comment:修改时间戳"`
}

func (u *UserBalanceDepositTransaction) TableName() string {
	return "user_balance_deposit_transaction"
}

/*
type BalanceDepositTransaction struct {
	BusinessType           string `json:"businessType"`           //BalanceDepositTransaction
	Symbol                 string `json:"symbol"`                 //USDT USDC
	Network                string `json:"network"`                //
	Amount                 string `json:"amount"`                 //金额	9.87
	DepositTransactionTime string `json:"depositTransactionTime"` //createTime	入金时间
	Uid                    string `json:"uid"`                    //用户唯一id	“48763843”
}
*/

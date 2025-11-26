package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
	"xbase/log"
	"xrobot/internal/xtypes"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=PaymentAmountUser -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type PaymentAmountUser struct {
	ID             int64           `gorm:"column:id;size:64;primarykey;autoIncrement"` // 主键
	OrderID        string          `gorm:"column:order_id;uniqueIndex:uin_order_currency;not null;size:32"`
	UID            int64           `gorm:"column:uid;not null;size:64;index;comment:用户ID"`
	Usage          xtypes.Usage    `gorm:"column:usage;not null;size:32;comment:用途"`
	TelegramUid    int64           `gorm:"column:telegram_uid;not null;size:64;index;comment:telegram_userid"` // 账号
	ChannelName    string          `gorm:"column:channel_name;size:64"`                                        // 渠道名
	ChannelCode    string          `gorm:"column:channel_code;size:32"`
	Currency       xtypes.Currency `gorm:"column:currency;uniqueIndex:uin_amount_currency;uniqueIndex:uin_order_currency;size:32;comment:币种,USDT,USDC,大写"`
	Amount         string          `gorm:"column:amount;size:64;uniqueIndex:uin_amount_currency;"` // 编号
	Extend         *AmountUserInfo `gorm:"column:extend;type:json"`
	ExpirationTime int64           `gorm:"column:expiration_time;not null;index;size:64"`
	CreateAt       time.Time       `gorm:"column:created_at;type:timestamp;comment:创建时间戳" json:"-"`
	UpdateAt       time.Time       `gorm:"column:updated_at;type:timestamp;comment:修改时间戳" json:"-"`
}

func (u *PaymentAmountUser) TableName() string {
	return "payment_amount_user"
}

type AddressInfo struct {
	Address   string `json:"address,omitempty"`
	Activated bool   `json:"activated"`
}
type AddressInfoType []*AddressInfo

func (at AddressInfoType) Clone() AddressInfoType {
	if at == nil {
		return nil
	}
	rst := make(AddressInfoType, 0)
	for _, item := range at {
		rst = append(rst, &AddressInfo{
			Address:   item.Address,
			Activated: item.Activated,
		})
	}
	return rst
}

type AmountUserInfo struct {
	AddressInfo AddressInfoType `json:"address,omitempty"`
	MessageID   int             `json:"messageID,omitempty"`
	BiShu       int64           `json:"biShu,omitempty"`
}

func (ot *AmountUserInfo) Clone() *AmountUserInfo {
	if ot == nil {
		return nil
	}
	rst := &AmountUserInfo{
		AddressInfo: ot.AddressInfo.Clone(),
		MessageID:   ot.MessageID,
		BiShu:       ot.BiShu,
	}
	return rst
}
func (ot AmountUserInfo) Value() (driver.Value, error) {

	s, err := json.Marshal(ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil, err
	}
	return s, nil
}

func (ot *AmountUserInfo) Scan(value any) error {
	if value == nil {
		return nil
	}
	if ot == nil {
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	err := json.Unmarshal(s, ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil
	}
	return nil
}

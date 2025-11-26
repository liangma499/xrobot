package xtypes

import (
	"database/sql/driver"
	"encoding/json"
	tgtypes "tron_robot/internal/xtelegram/tg-types"
	"xbase/errors"
	"xbase/log"

	"github.com/shopspring/decimal"
)

const (
	OfficialChannelCode = "c8888888"
)

type ChannelType int

const (
	ChannelTypeNormal ChannelType = 1 //  普通
	ChannelTypeTG     ChannelType = 2 // telegram
	ChannelTypeAll    ChannelType = 2 // telegram
)

type OfficialTypeKind int32

const (
	Official    OfficialTypeKind = 1 //  官方
	OfficialNot OfficialTypeKind = 2 //
	OfficialAll OfficialTypeKind = 3 // 非官方
)

type OptionChannelTelegram struct {
	MainRobotLink   string `json:"main_robot_link"`
	MainRobotToken  string `json:"main_robot_token"`
	PushRobotLink   string `json:"push_robot_link"`
	PushRobotToken  string `json:"push_robot_token"`
	CustomerRobot   string `json:"customer_robot"`
	CommunityLink   string `json:"community_link"`
	MainChannel     string `json:"main_channel"`
	MainChannelID   int64  `json:"main_channel_id"`
	MainChannelLink string `json:"main_channel_link"`
	GroupID         int64  `json:"group_id"`
	GroupLink       string `json:"group_link"`
}

func (ot *OptionChannelTelegram) Clone() *OptionChannelTelegram {
	if ot == nil {
		return nil
	}
	return &OptionChannelTelegram{
		MainRobotLink:   ot.MainRobotLink,
		MainRobotToken:  ot.MainRobotToken,
		PushRobotLink:   ot.PushRobotLink,
		PushRobotToken:  ot.PushRobotToken,
		CustomerRobot:   ot.CustomerRobot,
		CommunityLink:   ot.CommunityLink,
		MainChannel:     ot.MainChannel,
		MainChannelID:   ot.MainChannelID,
		MainChannelLink: ot.MainChannelLink,
		GroupID:         ot.GroupID,
		GroupLink:       ot.GroupLink,
	}
}
func (ot OptionChannelTelegram) Value() (driver.Value, error) {

	s, err := json.Marshal(ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil, err
	}
	return s, nil
}

func (ot *OptionChannelTelegram) Scan(value any) error {
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

type BiShuCombo struct {
	Price    decimal.Decimal `json:"price"`    //价格
	Currency string          `json:"currency"` //币种
}

func (ot *BiShuCombo) Clone() *BiShuCombo {
	if ot == nil {
		return nil
	}
	return &BiShuCombo{
		Price:    ot.Price,
		Currency: ot.Currency,
	}
}

type ComboKindInfo struct {
	ComboKind ComboKind `json:"comboKind"` //套餐类型
	Duration  int       `json:"duration"`  //套餐时长
}
type MapComboKindInfo map[tgtypes.XTelegramButton]*BiShuCombo

func (mc MapComboKindInfo) Clone() MapComboKindInfo {
	if mc == nil {
		return nil
	}
	rst := make(MapComboKindInfo)
	for k, v := range mc {
		rst[k] = v.Clone()
	}
	return rst
}

type OptionChannelCfg struct {
	BiShuCfg                   MapComboKindInfo `json:"biShuCfg"`                   //笔数配置
	PriceBiShuMax              int32            `json:"priceBiShuMax"`              //最大值
	Customer                   string           `json:"customer"`                   //客服配置
	EnergySavings              string           `json:"energySavings"`              //节省配置
	ComboKindEnergyFlashRental ComboKindInfo    `json:"comboKindEnergyFlashRental"` //能量闪租套餐
	ActivationFee              decimal.Decimal  `json:"activationFee"`              //激活单价

}

func (ot *OptionChannelCfg) Clone() *OptionChannelCfg {
	if ot == nil {
		return nil
	}
	rst := &OptionChannelCfg{
		BiShuCfg:      ot.BiShuCfg.Clone(),
		PriceBiShuMax: ot.PriceBiShuMax,
		Customer:      ot.Customer,
		EnergySavings: ot.EnergySavings,
		ComboKindEnergyFlashRental: ComboKindInfo{
			ComboKind: ot.ComboKindEnergyFlashRental.ComboKind,
			Duration:  ot.ComboKindEnergyFlashRental.Duration,
		},
		ActivationFee: ot.ActivationFee,
	}

	return rst
}
func (ot OptionChannelCfg) Value() (driver.Value, error) {

	s, err := json.Marshal(ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil, err
	}
	return s, nil
}

func (ot *OptionChannelCfg) Scan(value any) error {
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

type Price struct {
	EnergyPricesU    decimal.Decimal `json:"energyPricesU,omitempty"`   //有U交易需要的能量
	EnergyPricesNou  decimal.Decimal `json:"energyPricesNou,omitempty"` //无U交易需要的能量
	TrxPriceU        decimal.Decimal `json:"erxPriceU,omitempty"`       //有U的转账价格
	TrxPriceNoU      decimal.Decimal `json:"erxPriceNoU,omitempty"`     //无U的转账价格
	CustomizeBalance decimal.Decimal `json:"customizeBalance"`
}

func (ot *Price) Clone() *Price {
	if ot == nil {
		return nil
	}
	rst := &Price{
		EnergyPricesU:    ot.EnergyPricesU.Copy(), //Block
		EnergyPricesNou:  ot.EnergyPricesNou.Copy(),
		TrxPriceU:        ot.TrxPriceU.Copy(),
		TrxPriceNoU:      ot.TrxPriceNoU.Copy(),
		CustomizeBalance: ot.CustomizeBalance.Copy(),
	}
	return rst
}
func (ot Price) Value() (driver.Value, error) {

	s, err := json.Marshal(ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil, err
	}
	return s, nil
}

func (ot *Price) Scan(value any) error {
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

package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
	"xbase/log"
	"xrobot/internal/xtypes"
)

//go:generate xgorm-dao-generator -model-dir=. -model-names=PaymentNetwork -dao-dir=../dao/ -sub-pkg-enable=true -mysql-pkg-path=/internal/component/mysql/mysql-default
type PaymentNetwork struct {
	NetWork     xtypes.NetWork `gorm:"column:network;primaryKey;size:32;" json:"uid" redis:"uid"` // 主键
	NetworkInfo *NetworkInfo   `gorm:"column:network_info;type:json"`
	CreateAt    time.Time      `gorm:"column:created_at;type:timestamp;comment:创建时间戳" json:"-"`
	UpdateAt    time.Time      `gorm:"column:updated_at;type:timestamp;comment:修改时间戳" json:"-"`
}

func (u *PaymentNetwork) TableName() string {
	return "payment_network"
}

type BlockExtend struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
	Now   int64 `json:"now,omitempty"`
	Key   int64 `json:"key,omitempty"`
}

func (ot *BlockExtend) Clone() *BlockExtend {
	if ot == nil {
		return nil
	}
	return &BlockExtend{
		Start: ot.Start,
		End:   ot.End,
		Now:   ot.Now,
		Key:   ot.Key,
	}
}

type BlockExtendMap map[int64]*BlockExtend

func (ot BlockExtendMap) Clone() BlockExtendMap {
	if ot == nil {
		return nil
	}
	rst := make(BlockExtendMap)
	for k, v := range ot {
		rst[k] = v.Clone()
	}
	return rst
}

type NetworkInfo struct {
	Block       int64          `json:"block,omitempty"` //Start number. Default 0
	BlockExtend BlockExtendMap `json:"block_extend,omitempty"`
}

func (ot *NetworkInfo) Clone() *NetworkInfo {
	if ot == nil {
		return nil
	}
	rst := &NetworkInfo{
		Block:       ot.Block, //Block
		BlockExtend: ot.BlockExtend.Clone(),
	}
	return rst
}
func (ot NetworkInfo) Value() (driver.Value, error) {

	s, err := json.Marshal(ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil, err
	}
	return s, nil
}

func (ot *NetworkInfo) Scan(value any) error {
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

package xtypes

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const (
	MaxSupervisorLevel = 2
)

type SupervisorInfo map[int64]int64

func (sup SupervisorInfo) Value() (driver.Value, error) {
	byteData, err := json.Marshal(sup)
	if err != nil {
		return nil, err
	}
	return string(byteData), nil
}
func (sup *SupervisorInfo) Scan(value any) error {
	if value == nil {
		*sup = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	return json.Unmarshal(s, sup)
}
func (sup SupervisorInfo) Clone() SupervisorInfo {
	info := make(map[int64]int64)
	if sup == nil {
		return info
	}
	if len(sup) == 0 {
		return info
	}
	for id, item := range sup {
		info[id] = item
	}

	return info
}

// 上级推广了下级
func (sup SupervisorInfo) Supervisor(UID int64) SupervisorInfo {
	supervisor := make(map[int64]int64)

	supervisor[1] = UID

	level := MaxSupervisorLevel + 1
	supLen := len(sup) + 1

	if supLen < level {
		level = supLen
	}

	for i := 2; i < level; i++ {
		index := int64(i - 1)
		if uid, ok := sup[index]; ok {
			supervisor[int64(i)] = uid
		}

	}
	return supervisor
}

type SupervisorChannel map[int64]string

func (sup SupervisorChannel) Value() (driver.Value, error) {
	byteData, err := json.Marshal(sup)
	if err != nil {
		return nil, err
	}
	return string(byteData), nil
}
func (sup *SupervisorChannel) Scan(value any) error {
	if value == nil {
		*sup = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	return json.Unmarshal(s, sup)
}
func (sup SupervisorChannel) Clone() SupervisorChannel {
	info := make(map[int64]string)
	if sup == nil {
		return info
	}
	if len(sup) == 0 {
		return info
	}
	for id, item := range sup {
		info[id] = item
	}

	return info
}

// 上级推广了下级
func (sup SupervisorChannel) Supervisor(channelCode string) SupervisorChannel {
	supervisor := make(SupervisorChannel)

	supervisor[1] = channelCode

	level := MaxSupervisorLevel + 1
	supLen := len(sup) + 1

	if supLen < level {
		level = supLen
	}

	for i := 2; i < level; i++ {
		index := int64(i - 1)
		if uid, ok := sup[index]; ok {
			supervisor[int64(i)] = uid
		}

	}
	return supervisor
}

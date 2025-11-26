package tgnotice

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type FrequencyInfo struct {
	Time     string `json:"times"`    //几点发送就是几(24小时制)
	Interval int    `json:"interval"` //间隔时间(无间隔时间则为0) 到点只发一次
}

func (j FrequencyInfo) Value() (driver.Value, error) {
	byteData, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return string(byteData), nil
}
func (j *FrequencyInfo) Scan(value any) error {
	if value == nil {
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	return json.Unmarshal(s, j)
}

//===========================

type SendToType []int64

func (j SendToType) Value() (driver.Value, error) {
	byteData, err := json.Marshal(j)
	if err != nil {
		return nil, err
	}
	return string(byteData), nil
}
func (j *SendToType) Scan(value any) error {
	if value == nil {
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	return json.Unmarshal(s, j)
}

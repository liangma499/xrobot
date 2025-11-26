package xtypes

import (
	"database/sql/driver"
	"encoding/json"
	"xbase/errors"
	"xbase/log"
)

type APIKind int32

const (
	APITrongrid   APIKind = 1
	APITronscan   APIKind = 2
	APIGetblockIO APIKind = 3
	APIEtherscan  APIKind = 4
	APIBscscan    APIKind = 5
	APISolana     APIKind = 6
)

type DurationKind int32

const (
	DurationKind_None    DurationKind = 0 // 无
	DurationKind_Daily   DurationKind = 1 // 日榜
	DurationKind_Monthly DurationKind = 2 // 月榜
)

type ApiCfg struct {
	Url         string `json:"url,omitempty"`         //三方地址
	AppID       string `json:"appID,omitempty"`       //appID
	Secret      string `json:"secret,omitempty"`      //secret
	DurationMax int64  `json:"durationMax,omitempty"` //周期内可以最大数量
}

func (ot *ApiCfg) Clone() *ApiCfg {
	if ot == nil {
		return nil
	}
	return &ApiCfg{
		Url:         ot.Url,    //API接口基础地址 或者网络(主要针对区块)
		AppID:       ot.AppID,  //appID
		Secret:      ot.Secret, //SignatureKey
		DurationMax: ot.DurationMax,
	}

}

type KeyToApiCfgInfo map[string]*ApiCfg

func (ot KeyToApiCfgInfo) Clone() KeyToApiCfgInfo {
	if ot == nil {
		return nil
	}
	if len(ot) == 0 {
		return nil
	}
	rst := make(KeyToApiCfgInfo)

	for key, item := range ot {
		rst[key] = item.Clone()
	}
	return rst
}

type ApiCfgListInfo struct {
	Cfg          KeyToApiCfgInfo `json:"cfg,omitempty"`
	DurationKind DurationKind    `json:"durationKind,omitempty"`
}

func (ot *ApiCfgListInfo) Clone() *ApiCfgListInfo {
	if ot == nil {
		return nil
	}
	rst := &ApiCfgListInfo{
		DurationKind: ot.DurationKind,
		Cfg:          ot.Cfg.Clone(),
	}
	return rst
}

type ApiCfgList map[APIKind]*ApiCfgListInfo

func (ot ApiCfgList) Clone() ApiCfgList {
	if ot == nil {
		return nil
	}
	if len(ot) == 0 {
		return nil
	}
	rst := make(ApiCfgList)
	for key, item := range ot {
		rst[key] = item.Clone()
	}
	return rst
}

func (ot ApiCfgList) Value() (driver.Value, error) {

	s, err := json.Marshal(ot)
	if err != nil {
		log.Warnf("unmarshal:%v", err)
		return nil, err
	}
	return s, nil
}

func (ot *ApiCfgList) Scan(value any) error {
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

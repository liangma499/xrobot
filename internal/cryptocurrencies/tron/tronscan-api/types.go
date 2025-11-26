package tronscanapi

import "github.com/shopspring/decimal"

type ContractData struct {
	Amount       decimal.Decimal `json:"amount,omitempty"`
	OwnerAddress string          `json:"owner_address,omitempty"`
	ToAddress    string          `json:"to_address,omitempty"`
	AssetName    string          `json:"asset_name,omitempty"`
	TokenInfo    *TokenInfo      `json:"tokenInfo,omitempty"`
}

type Cost struct {
	NetFee             int64 `json:"net_fee,omitempty"`              //"net_fee": 268000,
	EnergyPenaltyTotal int64 `json:"energy_penalty_total,omitempty"` //"energy_penalty_total": 0,
	EnergyUsage        int64 `json:"energy_usage,omitempty"`         //"energy_usage": 0,
	Fee                int64 `json:"fee,omitempty"`                  //"fee": 268000,
	EnergyFee          int64 `json:"energy_fee,omitempty"`           //"energy_fee": 0,
	EnergyUsageTotal   int64 `json:"energy_usage_total,omitempty"`   //"energy_usage_total": 0,
	OriginEnergyUsage  int64 `json:"origin_energy_usage,omitempty"`  //"origin_energy_usage": 0,
	NetUsage           int64 `json:"net_usage,omitempty"`            //"net_usage": 0
}
type TokenInfo struct {
	TokenId      string `json:"tokenId,omitempty"`      //"tokenId": "_",
	TokenAbbr    string `json:"tokenAbbr,omitempty"`    //"tokenAbbr": "trx",
	TokenName    string `json:"tokenName,omitempty"`    //"tokenName": "trx",
	TokenDecimal int    `json:"tokenDecimal,omitempty"` //"tokenDecimal": 6,
	TokenCanShow int    `json:"tokenCanShow,omitempty"` //"tokenCanShow": 1,
	TokenType    string `json:"tokenType,omitempty"`    //"tokenType": "trc10",
	TokenLogo    string `json:"tokenLogo,omitempty"`    //"tokenLogo": "https://static.tronscan.org/production/logo/trx.png",
	TokenLevel   string `json:"tokenLevel,omitempty"`   //"tokenLevel": "2",
	IssuerAddr   string `json:"issuerAddr,omitempty"`   //"TB19pTknBYg2Ew6g7LeEo76dsirnswxawn"
	Vip          bool   `json:"vip,omitempty"`          //"vip": true

}

type TransactionData struct {
	Block         int64           `json:"block,omitempty"`
	Hash          string          `json:"hash,omitempty"`
	Timestamp     int64           `json:"timestamp,omitempty"`
	OwnerAddress  string          `json:"ownerAddress,omitempty"`
	ToAddressList []string        `json:"toAddressList,omitempty"`
	ToAddress     string          `json:"toAddress,omitempty"`
	ContractType  int32           `json:"contractType,omitempty"`
	Confirmed     bool            `json:"confirmed,omitempty"`
	Revert        bool            `json:"revert,omitempty"`
	ContractData  *ContractData   `json:"contractData,omitempty"`
	SmartCalls    string          `json:"SmartCalls,omitempty"`
	Events        string          `json:"Events,omitempty"`
	Id            string          `json:"id,omitempty"`
	Data          string          `json:"data,omitempty"`
	Fee           string          `json:"fee,omitempty"`
	ContractRet   string          `json:"contractRet,omitempty"`
	Result        string          `json:"result,omitempty"`
	Amount        decimal.Decimal `json:"amount,omitempty"`
	Cost          *Cost           `json:"cost,omitempty"`
	TokenInfo     *TokenInfo      `json:"tokenInfo,omitempty"`
	TokenType     string          `json:"tokenType,omitempty"`
}

type ConfirmList struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
	Block   int64  `json:"block,omitempty"`
	Url     string `json:"url,omitempty"`
}

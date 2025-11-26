package tronscanapi

import (
	"xrobot/internal/cryptocurrencies/tron/internal"

	"github.com/shopspring/decimal"
)

type AccountDetailV2Req struct {
	Address string `json:"address"` //Start index，default is 0
}

func (t *AccountDetailV2Req) ToMap() map[string]any {
	if t == nil {
		return nil
	}
	return map[string]any{
		"address": t.Address,
	}
}

// https://apilist.tronscanapi.com/api/accountv2?address=TEj8NgQM37dABXxVQRZo2b7nTkoCFM2qCQ
type Keys struct {
	Address string          `json:"address,omitempty"`
	Weight  decimal.Decimal `json:"weight,omitempty"`
}
type OwnerPermission struct {
	Keys           []*Keys         `json:"keys,omitempty"`
	Threshold      decimal.Decimal `json:"threshold,omitempty"`
	PermissionName string          `json:"permission_name,omitempty"`
}

type WithPriceTokens struct {
	Amount           decimal.Decimal `json:"amount,omitempty"`
	TokenPriceInTrx  decimal.Decimal `json:"tokenPriceInTrx,omitempty"`
	TokenId          string          `json:"tokenId,omitempty"`
	Balance          decimal.Decimal `json:"balance,omitempty"`
	TokenName        string          `json:"tokenName,omitempty"`
	TokenDecimal     int32           `json:"tokenDecimal,omitempty"`
	TokenAbbr        string          `json:"tokenAbbr,omitempty"`
	TokenCanShow     int32           `json:"tokenCanShow,omitempty"`
	TokenType        string          `json:"tokenType,omitempty"`
	Vip              bool            `json:"vip,omitempty"`
	TokenLogo        string          `json:"tokenLogo,omitempty"`
	NrOfTokenHolders decimal.Decimal `json:"nrOfTokenHolders,omitempty"`
	TransferCount    decimal.Decimal `json:"transferCount,omitempty"`
}
type Representative struct {
	LastWithDrawTime int64           `json:"lastWithDrawTime,omitempty"`
	Allowance        decimal.Decimal `json:"allowance,omitempty"`
	Enabled          bool            `json:"enabled,omitempty"`
	Url              string          `json:"url,omitempty"`
}
type Assets struct {
	NetPercentage decimal.Decimal `json:"netPercentage,omitempty"`
	NetLimit      decimal.Decimal `json:"netLimit,omitempty"`
	NetRemaining  decimal.Decimal `json:"netRemaining,omitempty"`
	NetUsed       decimal.Decimal `json:"netUsed,omitempty"`
}
type Bandwidth struct {
	EnergyRemaining   decimal.Decimal    `json:"energyRemaining,omitempty"`
	TotalEnergyLimit  decimal.Decimal    `json:"totalEnergyLimit,omitempty"`
	TotalEnergyWeight decimal.Decimal    `json:"totalEnergyWeight,omitempty"`
	NetUsed           decimal.Decimal    `json:"netUsed,omitempty"`
	StorageLimit      decimal.Decimal    `json:"storageLimit,omitempty"`
	StoragePercentage decimal.Decimal    `json:"storagePercentage,omitempty"`
	Assets            map[string]*Assets `json:"assets,omitempty"`
	NetPercentage     decimal.Decimal    `json:"netPercentage,omitempty"`
	StorageUsed       decimal.Decimal    `json:"storageUsed,omitempty"`
	StorageRemaining  decimal.Decimal    `json:"storageRemaining,omitempty"`
	FreeNetLimit      decimal.Decimal    `json:"freeNetLimit,omitempty"`
	EnergyUsed        decimal.Decimal    `json:"energyUsed,omitempty"`
	FreeNetRemaining  decimal.Decimal    `json:"freeNetRemaining,omitempty"`
	NetLimit          decimal.Decimal    `json:"netLimit,omitempty"`
	NetRemaining      decimal.Decimal    `json:"netRemaining,omitempty"`
	EnergyLimit       decimal.Decimal    `json:"energyLimit,omitempty"`
	FreeNetUsed       decimal.Decimal    `json:"freeNetUsed,omitempty"`
	TotalNetWeight    decimal.Decimal    `json:"totalNetWeight,omitempty"`
	FreeNetPercentage decimal.Decimal    `json:"freeNetPercentage,omitempty"`
	EnergyPercentage  decimal.Decimal    `json:"energyPercentage,omitempty"`
	TotalNetLimit     decimal.Decimal    `json:"totalNetLimit,omitempty"`
}
type Frozen struct {
	Total    decimal.Decimal `json:"total,omitempty"`
	Balances any             `json:"balances,omitempty"`
}
type AccountResource struct {
	Frozen_balance_for_energy any `json:"frozen_balance_for_energy,omitempty"`
}
type ActivePermissions struct {
	ID              int64           `json:"id,omitempty"`
	Operations      string          `json:"operations,omitempty"`
	Keys            []*Keys         `json:"keys,omitempty"`
	Threshold       decimal.Decimal `json:"threshold,omitempty"`
	Typ             string          `json:"type,omitempty"`
	Permission_name string          `json:"permission_name,omitempty"`
}
type AccountDetailV2Resp struct {
	TransactionsOut                    decimal.Decimal      `json:"transactions_out,omitempty"`
	AcquiredDelegateFrozenForBandWidth decimal.Decimal      `json:"acquiredDelegateFrozenForBandWidth,omitempty"`
	RewardNum                          decimal.Decimal      `json:"rewardNum,omitempty"`
	GreyTag                            string               `json:"greyTag,omitempty"`
	OwnerPermission                    *OwnerPermission     `json:"ownerPermission,omitempty"`
	RedTag                             string               `json:"redTag,omitempty"`
	PublicTag                          string               `json:"publicTag,omitempty"`
	WithPriceTokens                    []*WithPriceTokens   `json:"WithPriceTokens,omitempty"`
	DelegateFrozenForEnergy            decimal.Decimal      `json:"delegateFrozenForEnergy,omitempty"`
	Balance                            decimal.Decimal      `json:"balance,omitempty"`
	FeedbackRisk                       bool                 `json:"feedbackRisk,omitempty"`
	VoteTotal                          decimal.Decimal      `json:"voteTotal,omitempty"`
	TotalFrozen                        decimal.Decimal      `json:"totalFrozen,omitempty"`
	Delegated                          any                  `json:"delegated,omitempty"`
	Transactions_in                    int64                `json:"transactions_in,omitempty"`
	Latest_operation_time              int64                `json:"latest_operation_time,omitempty"`
	TotalTransactionCount              int64                `json:"totalTransactionCount,omitempty"`
	Representative                     *Representative      `json:"representative,omitempty"`
	FrozenForBandWidth                 decimal.Decimal      `json:"frozenForBandWidth,omitempty"`
	Announcement                       string               `json:"announcement,omitempty"`
	Reward                             decimal.Decimal      `json:"reward,omitempty"`
	AddressTagLogo                     string               `json:"addressTagLogo,omitempty"`
	AllowExchange                      any                  `json:"allowExchange,omitempty"`
	Address                            string               `json:"address,omitempty"`
	Frozen_supply                      any                  `json:"frozen_supply,omitempty"`
	Bandwidth                          *Bandwidth           `json:"bandwidth,omitempty"`
	Date_created                       int64                `json:"date_created,omitempty"`
	AccountType                        int32                `json:"accountType,omitempty"`
	Exchanges                          any                  `json:"exchanges,omitempty"`
	Frozen                             *Frozen              `json:"frozen,omitempty"`
	AccountResource                    *AccountResource     `json:"accountResource,omitempty"`
	Transactions                       int64                `json:"transactions,omitempty"`
	BlueTag                            string               `json:"blueTag,omitempty"`
	Witness                            decimal.Decimal      `json:"witness,omitempty"`
	DelegateFrozenForBandWidth         decimal.Decimal      `json:"delegateFrozenForBandWidth,omitempty"`
	Name                               string               `json:"name,omitempty"`
	FrozenForEnergy                    decimal.Decimal      `json:"frozenForEnergy,omitempty"`
	Activated                          bool                 `json:"activated,omitempty"`
	AcquiredDelegateFrozenForEnergy    decimal.Decimal      `json:"acquiredDelegateFrozenForEnergy,omitempty"`
	ActivePermissions                  []*ActivePermissions `json:"activePermissions,omitempty"`
}

func GetAccountDetailV2(baseUrl, apiKey string, args *AccountDetailV2Req) (*AccountDetailV2Resp, error) {
	c := internal.NewClient(baseUrl, apiKey, false)
	resp := &AccountDetailV2Resp{
		WithPriceTokens: make([]*WithPriceTokens, 0),
	}
	err := c.Get("/api/accountv2", args.ToMap(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (adr *AccountDetailV2Resp) GetWithPriceTokensByTokenAbbr(tokenAbbr string) *WithPriceTokens {
	if adr == nil {
		return nil
	}
	for _, item := range adr.WithPriceTokens {
		if item.TokenAbbr == tokenAbbr {
			return item
		}
	}
	return nil
}

// 合约地址
func (adr *AccountDetailV2Resp) GetWithPriceTokensByTokenId(tokenId string) *WithPriceTokens {
	if adr == nil {
		return nil
	}
	for _, item := range adr.WithPriceTokens {
		if item.TokenId == tokenId || ((tokenId == "-" || tokenId == "_") && (item.TokenId == "-" || item.TokenId == "_")) {
			return item
		}
	}
	return nil
}

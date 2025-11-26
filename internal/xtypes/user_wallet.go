package xtypes

type WalletDefaultStatus int32

// 默认货币状态
const (
	WalletDefaultNo  WalletDefaultStatus = 0
	WalletDefaultYes WalletDefaultStatus = 1
)

type WalletAmountKind int32

const (
	WalletAmountKindCash WalletAmountKind = 1 // true

)

func (w WalletAmountKind) Int32() int32 {
	return int32(w)
}

var WalletAmountKindUsed = []WalletAmountKind{WalletAmountKindCash} // true

const (
	UserWalletBalanceKey    = "user:%d:wallet:%s:kind:%d:balance" // 用户钱包余额
	UserWalletFieldid       = "id"                                // 用户ID
	UserWalletFieldUID      = "uid"                               // 用户ID
	UserWalletFieldCurrency = "currency_id"                       // 货币类型
	UserWalletFieldCash     = "cash"                              // 现金
	UserWalletFieldUsed     = "used"                              // 已使用
	UserWalletFieldDef      = "def"                               // 是否默认货币
	UserWalletAmountKindKey = "walletAmountKind"                  // 用户钱包类型
)

func (t WalletAmountKind) IsValid() bool {
	switch t {
	case WalletAmountKindCash:
		{
			return true
		}
	}
	return false
}

type UserControlKind int32

const (
	UserNone UserControlKind = 0
	UserWin  UserControlKind = 1
	UserLose UserControlKind = 2
)

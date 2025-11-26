package cryptocurrencyevent

const (
	cryptocurrencySuccessTopic = "cryptocurrency:success" // 钱包余额变动
)

// 公共字段
type CryptoCurrencyMsg struct {
	TimeUninx int64 `json:"timeUninx"`
}

func (mc CryptoCurrencyMsg) Clone() CryptoCurrencyMsg {

	return CryptoCurrencyMsg{
		TimeUninx: mc.TimeUninx,
	}
}

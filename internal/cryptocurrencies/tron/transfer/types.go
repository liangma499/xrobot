package transfer

import (
	"time"
	"tron_robot/internal/xtypes"

	"github.com/shopspring/decimal"
)

const (
	timeOut = 10 * time.Second
)

type CollectionTransferOutRes struct {
	Amount     float64
	TxID       string
	OutAddress string
	BeforTrx   int64
}

type transferContract struct {
	Protocol    string
	ToAddress   string          //转入地址
	FromAddress string          //转出地址
	RealAmount  decimal.Decimal //真实金额
	Contract    string          //合约名 trx 为-
	Currency    xtypes.Currency //币种名
	IsEngergy   bool
}

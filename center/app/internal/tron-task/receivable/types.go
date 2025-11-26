package receivable

import "xrobot/internal/xtypes"

type addressCheckRes struct {
	transactionKind xtypes.TransactionKind
	channelCode     string
	addressKind     xtypes.AddressKind
}

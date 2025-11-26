package receivable

import "tron_robot/internal/xtypes"

type addressCheckRes struct {
	transactionKind xtypes.TransactionKind
	channelCode     string
	addressKind     xtypes.AddressKind
}

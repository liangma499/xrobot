package solana

import (
	"tron_robot/internal/code"

	"fmt"
	httpClient "tron_robot/internal/cryptocurrencies/solana/internal/http-client"
	optionCurrencyNetworkCfg "tron_robot/internal/option/option-currency-network"
	"tron_robot/internal/xtypes"
	"xbase/errors"

	"tron_robot/internal/cryptocurrencies/solana/internal"
	"tron_robot/internal/cryptocurrencies/solana/internal/rpc"

	"github.com/shopspring/decimal"
)

const (
	netWorkChannelType = xtypes.NetWorkChannelType_Solana
	BlockNumUrl        = "https://api.mainnet-beta.solana.com"
	JsonRpcVersion     = "2.0"
	GetBlockIo         = "getblock.io"
)

type CommonReq struct {
	ID      string `json:"id"`      // "id": "getblock.io"
	JsonRpc string `json:"jsonrpc"` //"jsonrpc": "2.0",
	Method  string `json:"method"`  //"method": "getblockcount",
	Params  []any  `json:"params"`  //"params": [],

}
type toTransaction struct {
	TxID            string
	Contract        string
	From            string
	To              string
	Protocol        string
	Amount          decimal.Decimal
	Fee             decimal.Decimal
	Type            string
	RecentBlockhash string
	IsDecimals      bool
}
type solanaInfo struct {
	client     *rpc.Client
	httpClient *httpClient.HttpClient
	httpToken  string
}
type TransactionRet struct {
	TxID     string
	Contract string
	From     string
	To       string
	IsCreate bool
}

func NewModuleProxyFetch() (*solanaInfo, error) {

	apiCfg := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(netWorkChannelType, xtypes.APISolana)
	if apiCfg == nil {
		return nil, errors.NewError(code.OptionNotFound)
	}
	if apiCfg.AppID == "" {
		return &solanaInfo{
			client: rpc.New(apiCfg.Url),
		}, nil

	} else {
		return &solanaInfo{
			httpClient: httpClient.NewClient(apiCfg.Url, true),
			httpToken:  apiCfg.AppID,
		}, nil
	}

}

func NewModuleProxyByAppIDUrlFetch(url, httpToken string) (*solanaInfo, error) {

	if httpToken == "" {
		return &solanaInfo{
			client: rpc.New(url),
		}, nil
	} else {
		return &solanaInfo{
			httpClient: httpClient.NewClient(url, true),
			httpToken:  httpToken,
		}, nil
	}

}

func (sol *solanaInfo) getLastinstructionsTransfer(tx rpc.ParsedTransactionWithMeta) *rpc.ParsedInstruction {
	var info *rpc.InstructionInfo
	rst := make([]*rpc.ParsedInstruction, 0)
	for _, item := range tx.Transaction.Message.Instructions {
		if item.Parsed == nil {
			continue
		}
		info = item.Parsed.GetInstructionInfo()
		if info == nil {
			continue
		}
		if info.InstructionType == "transfer" {
			rst = append(rst, item)
		}
	}
	lenght := len(rst)
	if lenght <= 0 {
		return nil
	}
	return rst[lenght-1]
}
func (sol *solanaInfo) GetDecimals(tx rpc.ParsedTransactionWithMeta, authority, programId string, places int8) int8 {

	if tx.Meta == nil {
		return places
	}

	for _, item := range tx.Meta.PreTokenBalances {
		if item.Owner.String() == authority &&
			item.ProgramId.String() == programId {
			if item.UiTokenAmount != nil {
				return int8(item.UiTokenAmount.Decimals)
			}
		}

	}
	return places
}

func (sol *solanaInfo) Contract(tx rpc.ParsedTransactionWithMeta) (*toTransaction, error) {
	if tx.Meta == nil {
		return nil, fmt.Errorf("meta  is nil")
	}
	if tx.Meta.Err != nil {
		return nil, fmt.Errorf("err:%v", tx.Meta.Err)
	}
	if tx.Transaction == nil {
		return nil, fmt.Errorf("transaction is nil")
	}
	if tx.Transaction.Message.AccountKeys == nil {
		return nil, fmt.Errorf("transaction is nil")
	}
	if len(tx.Transaction.Signatures) == 0 {
		return nil, fmt.Errorf("tx.Transaction.Signatures len is zero")
	}
	txID := tx.Transaction.Signatures[0].String()

	//判断instructions
	trx := sol.getLastinstructionsTransfer(tx)
	if trx == nil {
		return nil, nil
	}
	if trx.Parsed == nil {
		return nil, nil
	}

	parsedInfo := trx.Parsed.GetInstructionInfo()
	if parsedInfo == nil {
		return nil, nil
	}
	contractInfo := &toTransaction{
		TxID:            txID,
		Contract:        trx.ProgramId.String(),
		Protocol:        parsedInfo.Info.Mint.String(),
		Type:            parsedInfo.InstructionType,
		Fee:             decimal.NewFromUint64(tx.Meta.Fee),
		RecentBlockhash: tx.Transaction.Message.RecentBlockHash,
	}
	//系统转sol
	if trx.Program == "system" {
		contractInfo.From = parsedInfo.Info.Source.String()
		contractInfo.To = parsedInfo.Info.Destination.String()
		contractInfo.Amount = parsedInfo.Info.Lamports.Copy()
		contractInfo.IsDecimals = false
		return contractInfo, nil
	} else if trx.Program == "spl-token" { //合约转USDT等
		contractInfo.From = parsedInfo.Info.Authority.String()
		contractInfo.To = parsedInfo.Info.Destination.String()
		amount, err := decimal.NewFromString(parsedInfo.Info.Amount)
		if err != nil {
			return nil, fmt.Errorf("err:%v", err)
		}
		contractInfo.Amount = amount.Copy()
		contractInfo.IsDecimals = true
		return contractInfo, nil
	}

	return nil, nil
}

type GetBlockOpts struct {
	// Encoding for each returned Transaction, either "json", "jsonParsed", "base58" (slow), "base64".
	// If parameter not provided, the default encoding is "json".
	// - "jsonParsed" encoding attempts to use program-specific instruction parsers to return
	//   more human-readable and explicit data in the transaction.message.instructions list.
	// - If "jsonParsed" is requested but a parser cannot be found, the instruction falls back
	//   to regular JSON encoding (accounts, data, and programIdIndex fields).
	//
	// This parameter is optional.
	Encoding internal.EncodingType `json:"encoding,omitempty"`

	// Level of transaction detail to return.
	// If parameter not provided, the default detail level is "full".
	//
	// This parameter is optional.
	TransactionDetails rpc.TransactionDetailsType `json:"transactionDetails,omitempty"`

	// Whether to populate the rewards array.
	// If parameter not provided, the default includes rewards.
	//
	// This parameter is optional.
	Rewards *bool `json:"rewards,omitempty"`

	// "processed" is not supported.
	// If parameter not provided, the default is "finalized".
	//
	// This parameter is optional.
	Commitment rpc.CommitmentType `json:"commitment,omitempty"`

	// Max transaction version to return in responses.
	// If the requested block contains a transaction with a higher version, an error will be returned.
	MaxSupportedTransactionVersion *uint64 `json:"maxSupportedTransactionVersion,omitempty"`
}
type GetParsedTransactionOpts struct {
	// Encoding for each returned Transaction, either "json", "jsonParsed", "base58" (slow), "base64".
	// If parameter not provided, the default encoding is "json".
	// - "jsonParsed" encoding attempts to use program-specific instruction parsers to return
	//   more human-readable and explicit data in the transaction.message.instructions list.
	// - If "jsonParsed" is requested but a parser cannot be found, the instruction falls back
	//   to regular JSON encoding (accounts, data, and programIdIndex fields).
	//
	// This parameter is optional.
	Encoding internal.EncodingType `json:"encoding,omitempty"`
	// Desired commitment. "processed" is not supported. If parameter not provided, the default is "finalized".
	Commitment rpc.CommitmentType `json:"commitment,omitempty"`

	// Max transaction version to return in responses.
	// If the requested block contains a transaction with a higher version, an error will be returned.
	MaxSupportedTransactionVersion *uint64 `json:"maxSupportedTransactionVersion,omitempty"`
}

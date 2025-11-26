package bnbscan

import (
	"tron_robot/internal/code"
	"tron_robot/internal/xtypes"

	"fmt"
	"math/big"
	"strconv"
	"strings"
	"tron_robot/internal/cryptocurrencies/bnbscan/internal"
	optionCurrencyNetworkCfg "tron_robot/internal/option/option-currency-network"
	"xbase/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type BnbMethod string

const (
	netWorkChannelType            = xtypes.NetWorkChannelType_BNB
	method_BNB          BnbMethod = "0x"
	method_transfer     BnbMethod = "0xa9059cbb" //: "transfer(address,uint256)",
	method_transferFrom BnbMethod = "0x23b872dd" //: "transferFrom(address,address,uint256)",
	method_approve      BnbMethod = "0x095ea7b3" //: "approve(address,uint256)",
	method_allowance    BnbMethod = "0xdd62ed3e" //: "allowance(address,address)",
	method_balanceOf    BnbMethod = "0x70a08231" //: "balanceOf(address)",
	method_totalSupply  BnbMethod = "0x18160ddd" //: "totalSupply()",
	method_mint         BnbMethod = "0x40c10f19" //: "mint(address,uint256)",
	method_burn         BnbMethod = "0x42966c68" //: "burn(uint256)",
	method_getReserves  BnbMethod = "0xdb507c29" //: "getReserves()",
	method_setFeeTo     BnbMethod = "0x91477cac"
)

func (em BnbMethod) IsTransfer() bool {
	switch em {
	case method_BNB, method_transfer, method_transferFrom:
		{
			return true
		}
	}
	return false
}
func (em BnbMethod) Name() string {
	switch em {
	case method_BNB:
		{
			return "-"
		}
	case method_transfer:
		{
			return "transfer(address,uint256)"
		}
	case method_transferFrom:
		{
			return "transferFrom(address,address,uint256)"
		}
	case method_approve:
		{
			return "approve(address,uint256)"
		}
	case method_allowance:
		{
			return "allowance(address,address)"
		}
	case method_balanceOf:
		{
			return "balanceOf(address)"
		}
	case method_totalSupply:
		{
			return "totalSupply()"
		}
	case method_mint:
		{
			return "mint(address,uint256)"
		}
	case method_burn:
		{
			return "burn(uint256)"
		}
	case method_getReserves:
		{
			return "getReserves()"
		}
	case method_setFeeTo:
		{
			return "setFeeTo()"
		}
	}
	return ""
}
func (em BnbMethod) Contract(gb *ethInfo, tx *Transaction) (*toTransaction, error) {
	switch em {
	case method_transfer:
		{
			to, err := gb.ToAddress(tx.Input[10:74])
			if err != nil {
				return nil, fmt.Errorf("to:%v", err)
			}
			amout, err := gb.HexToInit64(tx.Input[74:])
			if err != nil {
				return nil, err
			}
			return &toTransaction{
				Contract: common.HexToAddress(tx.To).String(),
				From:     common.HexToAddress(tx.From).String(),
				To:       to,
				Protocol: em.Name(),
				Amount:   amout.Copy(),
				Type:     tx.Typ,
			}, nil
		}
	case method_transferFrom:
		{
			from, err := gb.ToAddress(tx.Input[10:74])
			if err != nil {
				return nil, fmt.Errorf("from:%v", err)
			}
			// 提取接收者地址

			to, err := gb.ToAddress(tx.Input[74:138])
			if err != nil {
				return nil, fmt.Errorf("to:%v", err)
			}

			amout, err := gb.HexToInit64(tx.Input[138:])
			if err != nil {
				return nil, err
			}
			return &toTransaction{
				Contract: common.HexToAddress(tx.To).String(),
				From:     from,
				To:       to,
				Protocol: em.Name(),
				Amount:   amout.Copy(),
				Type:     tx.Typ,
			}, nil
		}
	}
	return nil, fmt.Errorf("not  transfer")
}

type BnbCommon struct {
	Module string `json:"module"`
	Action string `json:"action"`
	Apikey string `json:"apikey"`
}

type ethInfo struct {
	Apikey string
	Client *internal.Client
}

func NewModuleProxy() (*ethInfo, error) {

	apiCfg := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(netWorkChannelType, xtypes.APIBscscan)
	if apiCfg == nil {
		return nil, errors.NewError(code.OptionNotFound)
	}
	return &ethInfo{
		Apikey: apiCfg.AppID,
		Client: internal.NewClient(apiCfg.Url, false),
	}, nil

}

func NewModuleProxyByAppIDUrl(appID, url string) (*ethInfo, error) {

	return &ethInfo{
		Apikey: appID,
		Client: internal.NewClient(url, false),
	}, nil

}

func (gb *ethInfo) removePrefix(hexStr string) string {
	if strings.HasPrefix(hexStr, "0x") || strings.HasPrefix(hexStr, "0X") {
		return hexStr[2:] // 去掉前两个字符
	}
	return hexStr
}
func (gb *ethInfo) HexToInit64(hexStr string) (decimal.Decimal, error) {
	hexStr = gb.removePrefix(hexStr)
	ethAmount := new(big.Int)
	_, b := ethAmount.SetString(hexStr, 16)

	if !b {
		return decimal.Zero, fmt.Errorf("can not to big int:%s", hexStr)
	}
	return decimal.NewFromBigInt(ethAmount, 0), nil
}

func (gb *ethInfo) Init64Hex(num int64) string {
	return fmt.Sprintf("0x%s", strconv.FormatInt(num, 16))
}

type toTransaction struct {
	Contract string
	From     string
	To       string
	Protocol string
	Amount   decimal.Decimal
	Type     string
}

func (gb *ethInfo) Contract(tx *Transaction) (*toTransaction, error) {

	//这个是ETH转ETH不需要合约
	if tx.Input == "0x" {

		amout, err := gb.HexToInit64(tx.Value)
		if err != nil {
			return nil, err
		}
		return &toTransaction{
			Contract: "-",
			From:     common.HexToAddress(tx.From).String(),
			To:       common.HexToAddress(tx.To).String(),
			Protocol: "-",
			Amount:   amout.Copy(),
			Type:     tx.Typ,
		}, nil

	} else if len(tx.Input) >= 138 {
		method := BnbMethod(tx.Input[:10])
		if !method.IsTransfer() {
			return nil, nil
		}
		return method.Contract(gb, tx)
	}

	return nil, nil
}

func (gb *ethInfo) ToAddress(address string) (string, error) {
	fromBig := new(big.Int)
	_, b := fromBig.SetString(address, 16)
	if !b {
		return "", fmt.Errorf("address is err")
	}
	return common.HexToAddress("0x" + fromBig.Text(16)).String(), nil
}

type BlockInfo struct {
	BaseFeePerGas    string         `json:"baseFeePerGas"`
	Difficulty       string         `json:"difficulty"`
	ExtraData        string         `json:"extraData"`
	GasLimit         string         `json:"gasLimit"`
	GasUsed          string         `json:"gasUsed"`
	Hash             string         `json:"hash"`
	LogsBloom        string         `json:"logsBloom"`
	Miner            string         `json:"miner"`
	MixHash          string         `json:"mixHash"`
	Nonce            string         `json:"nonce"`
	Number           string         `json:"number"`
	ParentHash       string         `json:"parentHash"`
	ReceiptsRoot     string         `json:"receiptsRoot"`
	Sha3Uncles       string         `json:"sha3Uncles"`
	Size             string         `json:"size"`
	StateRoot        string         `json:"stateRoot"`
	Timestamp        string         `json:"timestamp"`
	TotalDifficulty  string         `json:"totalDifficulty"`
	Transactions     []*Transaction `json:"transactions"`
	TransactionsRoot string         `json:"transactionsRoot"`
	Uncles           []string       `json:"uncles"`
}

type Transaction struct {
	BlockHash            string `json:"blockHash"`
	BlockNumber          string `json:"blockNumber"`
	From                 string `json:"from"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
	Hash                 string `json:"hash"`
	Input                string `json:"input"`
	Nonce                string `json:"nonce"`
	To                   string `json:"to"`
	TransactionIndexTxid string `json:"transactionIndex"`
	Value                string `json:"value"`
	Typ                  string `json:"type"`
	AccessList           any    `json:"accessList"`
	ChainId              string `json:"chainId"`
	V                    string `json:"v"`
	R                    string `json:"r"`
	S                    string `json:"s"`
}

type TransactionLog struct {
	Address          string   `json:"address"`          // 生成日志的合约地址。
	Topics           []string `json:"topics"`           //事件主题，用于标识事件类型和参数。
	Data             string   `json:"data"`             //附加的事件数据。
	BlockNumber      string   `json:"blockNumber"`      // 此交易所在区块的编号（十六进制格式）。
	TransactionHash  string   `json:"transactionHash"`  //生成此日志的交易哈希。
	TransactionIndex string   `json:"transactionIndex"` //交易在区块中的索引。
	BlockHash        string   `json:"blockHash"`        // 表示此交易所在区块的哈希值。
	LogIndex         string   `json:"logIndex"`         //日志在区块中的索引。
	Removed          bool     `json:"removed"`
}

type TransactionReceipt struct {
	BlockHash         string            `json:"blockHash"`         // 表示此交易所在区块的哈希值。
	BlockNumber       string            `json:"blockNumber"`       // 此交易所在区块的编号（十六进制格式）。
	ContractAddress   string            `json:"contractAddress"`   // 如果这是一个合约创建交易，则此字段将包含合约地址；否则为 null。
	CumulativeGasUsed string            `json:"cumulativeGasUsed"` // 在区块中所有交易使用的累计 Gas。
	EffectiveGasPrice string            `json:"effectiveGasPrice"` // 实际支付的 Gas 价格。
	From              string            `json:"from"`              //发送交易的地址。
	GasUsed           string            `json:"gasUsed"`           // 此交易实际使用的 Gas。
	Logs              []*TransactionLog `json:"logs"`              // 交易生成的事件日志数组。
	LogsBloom         string            `json:"logsBloom"`         // Bloom 过滤器，帮助快速查找日志信息。
	Status            string            `json:"status"`            // 交易的状态，0x1 表示成功，0x0 表示失败。
	To                string            `json:"to"`                //交易接收方的地址。
	TransactionHash   string            `json:"transactionHash"`   //此交易的哈希值。
	TransactionIndex  string            `json:"transactionIndex"`  //交易在区块中的索引。
	Typ               string            `json:"type"`              //交易类型，0x2 表示 EIP-1559 交易。
}

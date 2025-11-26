package getblockio

import (
	"encoding/hex"
	"xbase/errors"
	"xrobot/internal/code"
	"xrobot/internal/cryptocurrencies/btcd/getblock-io/internal"
	optionCurrencyNetworkCfg "xrobot/internal/option/option-currency-network"
	"xrobot/internal/xtypes"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/shopspring/decimal"
)

const (
	jsonrpc            = "2.0"
	id                 = "getblock.io"
	netWorkChannelType = xtypes.NetWorkChannelType_BTC
)

type CommonReq struct {
	ID      string `json:"id"`      // "id": "getblock.io"
	JsonRpc string `json:"jsonrpc"` //"jsonrpc": "2.0",
	Method  string `json:"method"`  //"method": "getblockcount",
	Params  []any  `json:"params"`  //"params": [],

}

type getBlockIO struct {
	ID      string // "id": "getblock.io"
	JsonRpc string //"jsonrpc": "2.0",
	Method  string
	Client  *internal.Client
}

func NewBlockIO() (*getBlockIO, error) {

	apiCfg := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(netWorkChannelType, xtypes.APIGetblockIO)
	if apiCfg == nil {
		return nil, errors.NewError(code.OptionNotFound)
	}
	return &getBlockIO{
		ID:      id,
		JsonRpc: jsonrpc,
		Method:  "",
		Client:  internal.NewClient(apiCfg.Url+"/"+apiCfg.AppID, true),
	}, nil

}

func NewBlockIOByAppID(appID, url string) (*getBlockIO, error) {

	return &getBlockIO{
		ID:      id,
		JsonRpc: jsonrpc,
		Method:  "",
		Client:  internal.NewClient(url+"/"+appID, true),
	}, nil

}

type BlockInfo struct {
	Hash              string             `json:"hash"`              //"hash": "000000000000000000046b9302e08c16ea186950f42a5498320ddd1bd7ab3428",
	Confirmations     int64              `json:"confirmations"`     // "confirmations": 1197,
	Height            int64              `json:"height"`            //"height": 677119,
	Version           int64              `json:"version"`           //"version": 1073733632,
	VersionHex        string             `json:"versionHex"`        //"versionHex": "3fffe000",
	Merkleroot        string             `json:"merkleroot"`        //"merkleroot": "d14c9f467c4bdd5135837696150ab5f52f3f5043de324ca4e5766b195b9f8f37",
	Time              int64              `json:"time"`              // "time": 1617180599,
	Mediantime        int64              `json:"mediantime"`        //"mediantime": 1617176373,
	Nonce             int64              `json:"nonce"`             //"nonce": 3669423616,
	Bits              string             `json:"bits"`              //"bits": "170cdf6f",
	Difficulty        decimal.Decimal    `json:"difficulty"`        //"difficulty": 21865558044610.55,
	Fee               decimal.Decimal    `json:"fee"`               //"fee": 0.000087,
	ChainWork         string             `json:"chainwork"`         // "chainwork": "00000000000000000000000000000000000000001b633a711a2334c78a29bb40",
	NTx               int64              `json:"nTx"`               //"nTx": 2815,
	PreviousBlockHash string             `json:"previousblockhash"` // "previousblockhash": "0000000000000000000aec32aa6edda6c888e8d6a0183d9c976064f98430c2da",
	NextBlockHash     string             `json:"nextblockhash"`     //"nextblockhash": "000000000000000000006d8e1eb870bd281b30ed621acf6b8d6af2a3c7ab61f1",
	StrippedSize      int64              `json:"strippedsize"`      //"strippedsize": 882816,
	Size              int64              `json:"size"`              //"size": 1350854,
	Weight            int64              `json:"weight"`            //"weight": 3999302
	Tx                []*TransactionInfo `json:"tx"`                //

}

type ScriptPubKey struct {
	Asm     string `json:"asm"`     //"asm": "OP_RETURN 52534b424c4f434b3a043421be599ad2492b27d36a2ae44bae98541da533b31f655208180f006c3999",
	Desc    string `json:"desc"`    //"desc": "raw(6a2952534b424c4f434b3a043421be599ad2492b27d36a2ae44bae98541da533b31f655208180f006c3999)#597d8g88",
	Hex     string `json:"hex"`     //"hex": "6a2952534b424c4f434b3a043421be599ad2492b27d36a2ae44bae98541da533b31f655208180f006c3999",
	Typ     string `json:"type"`    // "type": "nulldata"
	Address string `json:"address"` //"address": "bc1ph8pmcyte76g60p9rgr26qg3q9jdmexta8r6zeswncehea2jdda6ste99x3",
}
type Out struct {
	Value        decimal.Decimal `json:"value"`        //   "value": 0.0,
	N            int64           `json:"n"`            //"n": 4,
	ScriptPubKey *ScriptPubKey   `json:"scriptPubKey"` //"n": 4,

}
type ScriptSig struct {
	Asm string `json:"asm"` //"asm": "",
	Hex string `json:"hex"` //"hex": ""
}
type Vin struct {
	Txid        string     `json:"txid"` //"txid": "c864f5111c4b55a44c8ba783d9615357e2528d96cd1065946ca46139d05adac9",
	Vout        int64      `json:"vout"` //"vout": 1,
	ScriptSig   *ScriptSig `json:"scriptSig"`
	Txinwitness []string   `json:"txinwitness"`
	Sequence    int64      `json:"sequence"`
}
type TransactionInfo struct {
	Txid     string          `json:"txid"`     // "txid": "22b4cba266517d14297a47eb19f4febe28b3761ae351a08d169b2959e645d4e0",
	Hash     string          `json:"hash"`     //"hash": "234b4046fe0c04d4657ccbe2471ca77bfb02b8f95402ecc96785f9bcd7fe059d",
	Version  int64           `json:"version"`  //"version": 1,
	Size     int64           `json:"size"`     //"size": 421,
	Vsize    int64           `json:"vsize"`    //"vsize": 394,
	Weight   int64           `json:"weight"`   //"weight": 1576,
	Locktime int64           `json:"locktime"` //"locktime": 0,
	Vin      []*Vin          `json:"vin"`
	Out      []*Out          `json:"vout"`
	Fee      decimal.Decimal `json:"fee"` //"fee": 0.000087,
	Hex      string          `json:"hex"`
}

type TransactionRawInfo struct {
	InActiveChain bool            `json:"in_active_chain"`
	Txid          string          `json:"txid"`     // "txid": "22b4cba266517d14297a47eb19f4febe28b3761ae351a08d169b2959e645d4e0",
	Hash          string          `json:"hash"`     //"hash": "234b4046fe0c04d4657ccbe2471ca77bfb02b8f95402ecc96785f9bcd7fe059d",
	Version       int64           `json:"version"`  //"version": 1,
	Size          int64           `json:"size"`     //"size": 421,
	Vsize         int64           `json:"vsize"`    //"vsize": 394,
	Weight        int64           `json:"weight"`   //"weight": 1576,
	Locktime      int64           `json:"locktime"` //"locktime": 0,
	Vin           []*Vin          `json:"vin"`
	Out           []*Out          `json:"vout"`
	Fee           decimal.Decimal `json:"fee"` //"fee": 0.000087,
	Hex           string          `json:"hex"`
	BlockHash     string          `json:"blockhash"`
	Confirmations int64           `json:"confirmations"`
	Time          int64           `json:"time"`
	BlockTime     int64           `json:"blocktime"`
}

func PubKeyToAddress(pubKeyHex string) (string, error) {

	hex, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return "", err
	}
	// 解析公钥
	pubKey, err := btcec.ParsePubKey(hex)
	if err != nil {
		return "", err
	}

	// 生成 P2PKH 地址
	address, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return "", err
	}
	return address.EncodeAddress(), nil

}

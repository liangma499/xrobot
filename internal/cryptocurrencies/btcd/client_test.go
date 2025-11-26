package btcd_test

import (
	"fmt"
	"testing"
	"xrobot/internal/cryptocurrencies/btcd/blockcypher"
	getblockio "xrobot/internal/cryptocurrencies/btcd/getblock-io"

	"xbase/log"
	"xbase/utils/xconv"
)

var (
	keys1   blockcypher.AddrKeychain
	keys2   blockcypher.AddrKeychain
	txhash1 string
	txhash2 string
	bcy     blockcypher.API
	err     error
)

func TestClient_Main(t *testing.T) {
	bcy = blockcypher.API{
		Token: "7d73fb49f40f48dfb5f44e066bb4254c",
		Coin:  "bcy",
		Chain: "test",
	}
	//Set Coin/Chain to BlockCypher testnet

	//Create/fund the test addresses
	keys1, err = bcy.GenAddrKeychain()
	if err != nil {
		log.Fatal("Error generating test addresses: ", err)
	}
	keys2, err = bcy.GenAddrKeychain()
	if err != nil {
		log.Fatal("Error generating test addresses: ", err)
	}
	txhash1, err = bcy.Faucet(keys1, 1e5)
	if err != nil {
		log.Fatal("Error funding test addresses: ", err)
	}
	txhash2, err = bcy.Faucet(keys2, 2e5)
	if err != nil {
		log.Fatal("Error funding test addresses: ", err)
	}

	conf, err := bcy.GetTXConf(txhash2)
	if err != nil {
		t.Error("Error encountered: ", err)
	}
	log.Warnf("keys1:%s", xconv.Json(keys1))
	log.Warnf("keys2:%s", xconv.Json(keys2))
	log.Warnf("txhash1:%s", txhash1)
	log.Warnf("txhash2:%s", txhash2)
	log.Warnf("conf:%s", xconv.Json(conf))
}

func TestClient_GetBlock(t *testing.T) {
	btc := blockcypher.API{
		Token: "7d73fb49f40f48dfb5f44e066bb4254c",
		Coin:  "btc",
		Chain: "main",
	}
	params := make(map[string]string)
	params["txstart"] = "1"
	params["limit"] = "500"
	block, err := btc.GetBlock(877608, "", params)
	if err != nil {
		fmt.Println(err)
	}
	log.Warnf("block:%s", xconv.Json(block))
	params = make(map[string]string)
	for _, txHash := range block.TXids {
		tx, err := btc.GetTX(txHash, params)
		if err != nil {
			log.Errorf("%v", err)
			continue
		}
		log.Warnf("tx:%s", xconv.Json(tx))
	}
}

func TestClient_GetUnTX(t *testing.T) {
	btc := blockcypher.API{
		Token: "2798e3b58280438592f56ecac1008df8",
		Coin:  "btc",
		Chain: "main",
	}

	unTX, err := btc.GetUnTX()
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	for _, txHash := range unTX {

		log.Warnf("tx:%s", xconv.Json(txHash))
	}
}

func TestClient_GetBlockCount(t *testing.T) {

	blockio, err := getblockio.NewBlockIOByAppID("9a4af4b7b8634ad9bfc0971f4499272d", "https://go.getblock.io")
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	blockHash, err := blockio.GetBlockHash(877762)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	log.Warnf("blockHash:%s", blockHash)

	blockTx, err := blockio.Getblock(blockHash)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	if blockTx.Result == nil {
		log.Errorf("%v", err)
		return
	}
	for _, item := range blockTx.Result.Tx {
		log.Warnf("blockTx:%s", xconv.Json(item))
	}

	/*for _, item := range blockTx.Tx {
		txInfo, err := blockio.GetRawTransaction(blockHash, item)
		if err != nil {
			continue
		}
		if txInfo.Err != nil {
			continue
		}
		log.Warnf("Result:%s", txInfo.Result)
	}*/
}
func TestClient_PubKeyToAddress(t *testing.T) {
	address, err := getblockio.PubKeyToAddress("1e100e8506a4c44ae7a56177df3483382b50bce2eb579e3c8ad1f832188e7cf8")
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	log.Warnf("address:%s", address)
}
func TestClient_GetRawTransaction(t *testing.T) {
	blockio, err := getblockio.NewBlockIOByAppID("9a4af4b7b8634ad9bfc0971f4499272d", "https://go.getblock.io")
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	txInfo, err := blockio.GetRawTransaction("00000000000000000002690252ad91ea10f6f84f7ae88bb2d0687b6f3b693bb0", "9bce3d2cf0202f9868642328401b6a34ac6ffa9a25010b1c344825aea9fee140")
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	log.Warnf("%s", xconv.Json(txInfo))
}

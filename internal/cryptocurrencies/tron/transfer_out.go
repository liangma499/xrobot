package tron

import (
	"fmt"
	"sync"
	"time"
	"tron_robot/internal/cryptocurrencies/tron/transfer"
	tronscanapi "tron_robot/internal/cryptocurrencies/tron/tronscan-api"
	optionCurrencyNetworkCfg "tron_robot/internal/option/option-currency-network"
	"tron_robot/internal/xtypes"
	"xbase/log"

	"github.com/shopspring/decimal"
)

const (
	rechargeChannelType = xtypes.NetWorkChannelType_TRON
)

type TransferOutRes struct {
	Amount     float64
	TxID       string
	OutAddress string
}
type TransferOut struct {
	mux sync.Mutex
}

var (
	once     sync.Once
	instance *TransferOut
)

func Instance() *TransferOut {
	once.Do(func() {
		instance = &TransferOut{}
	})
	return instance
}

// -1错误，0表示不需要能量 1表示对方账号已经激活 2表对方账号已经激活

func (to *TransferOut) Transfer_Out(toAddress string, currency xtypes.Currency, amount decimal.Decimal) (*TransferOutRes, *xtypes.EnergyExtra, error) {
	to.mux.Lock()
	defer to.mux.Unlock()

	privateKeyCfg, err := optionCurrencyNetworkCfg.GetPrivateKeyCfg(rechargeChannelType, currency)
	if err != nil {
		return nil, nil, fmt.Errorf("privateKeyCfg is not found:%v", err)
	}
	if privateKeyCfg == nil {
		return nil, nil, fmt.Errorf("privateKeyCfg is not found")
	}
	if privateKeyCfg.FromAddress == "" || privateKeyCfg.PrivateKey == "" {
		return nil, nil, fmt.Errorf("fromAddress or privateKey is not found")
	}
	apiCfg := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(rechargeChannelType, xtypes.APITrongrid)
	if apiCfg == nil {
		return nil, nil, fmt.Errorf("apiCfg is not found")
	}
	apiCfgScan := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(rechargeChannelType, xtypes.APITronscan)
	if apiCfgScan == nil {
		return nil, nil, fmt.Errorf("apiCfgScan is not found")
	}
	transferInfo := transfer.NewTransfer(apiCfg.Url, apiCfg.AppID)

	if transferInfo == nil {
		return nil, nil, fmt.Errorf("transferInfo new failed")
	}
	contract := currency.Trc20Contract()
	//转出TRX
	outAmout := xtypes.CoefficientInt64(amount, xtypes.Trc20_Places)

	rstFormInfo, err := tronscanapi.GetAccountDetailV2(apiCfgScan.Url, apiCfgScan.AppID, &tronscanapi.AccountDetailV2Req{
		Address: privateKeyCfg.FromAddress,
	})
	if err != nil {
		log.Warnf("%v", err)
		return nil, nil, fmt.Errorf("get info err:%v", err)
	}
	walletMustBlance := xtypes.CoefficientInt64(decimal.NewFromFloat(xtypes.Tron_Wallet_Must_Blance), xtypes.Trc20_Places)
	maxGasFeeLimmit := xtypes.CoefficientInt64(decimal.NewFromFloat(xtypes.Tron_MaxGasFeeLimmit), xtypes.Trc20_Places)
	tronWalletNetFee := xtypes.CoefficientInt64(decimal.NewFromFloat(xtypes.Tron_Wallet_Net_Fee), xtypes.Trc20_Places)
	// = 1
	//MaxGasFeeLimmit    = 3

	if contract == "-" || currency == xtypes.TRX {
		if rstFormInfo.Balance.LessThanOrEqual(outAmout.Add(walletMustBlance).Add(tronWalletNetFee)) {
			return nil, nil, fmt.Errorf("trx insufficien")
		}
		txHash, err := transferInfo.TRXSend(outAmout.IntPart(), privateKeyCfg.FromAddress, privateKeyCfg.PrivateKey, toAddress)
		if err != nil {
			return nil, nil, fmt.Errorf("trx send err:%v", err)
		}
		return &TransferOutRes{
			Amount:     amount.InexactFloat64(),
			TxID:       txHash,
			OutAddress: privateKeyCfg.FromAddress,
		}, nil, nil
	} else if len(contract) > 0 {
		fromTcy := rstFormInfo.GetWithPriceTokensByTokenId(contract)
		if fromTcy == nil {
			return nil, nil, fmt.Errorf("from not have contract")
		}
		if fromTcy.Balance.LessThanOrEqual(outAmout.Add(walletMustBlance)) {
			return nil, nil, fmt.Errorf("%s insufficien", currency.String())
		}

		if privateKeyCfg.ExtraCfg != nil && privateKeyCfg.ExtraCfg.IsUserEnergy {
			var energyExtra *xtypes.EnergyExtra
			if !transferInfo.ValidateAddress(privateKeyCfg.ExtraCfg.ExchangeAddress) {
				return nil, nil, fmt.Errorf("is not validateAddress")
			}

			//对方账号信息
			rstToInfo, err := tronscanapi.GetAccountDetailV2(apiCfgScan.Url, apiCfgScan.AppID, &tronscanapi.AccountDetailV2Req{
				Address: toAddress,
			})
			if err != nil {
				return nil, nil, fmt.Errorf("toAddress info err:%v", err)
			}

			toCurryInfo := rstToInfo.GetWithPriceTokensByTokenId(contract)
			needEnergy := privateKeyCfg.ExtraCfg.EnergyU
			bHaveU := true
			if toCurryInfo == nil {
				needEnergy = privateKeyCfg.ExtraCfg.EnergyNoU
				bHaveU = false
			} else {
				if toCurryInfo.Balance.LessThanOrEqual(decimal.Zero) {
					needEnergy = privateKeyCfg.ExtraCfg.EnergyNoU
					bHaveU = false
				}
			}
			//兑换能量
			needTrx := privateKeyCfg.ExtraCfg.PriceU
			if !bHaveU {
				needTrx = privateKeyCfg.ExtraCfg.PriceNoU
			}
			outNeedTrx := xtypes.CoefficientInt64(needTrx, xtypes.Trc20_Places)

			//TRX 不够
			//needMin := outNeedTrxDe.Add(decimal.NewFromFloat(Wallet_Must_Blance + maxGasFeeLimmit))
			needMin := outNeedTrx.Add(walletMustBlance).Add(tronWalletNetFee)
			log.Warnf("TRC20:outNeedTrxDe:%s needMin:%s balance:%s", outNeedTrx.String(), needMin.String(), rstFormInfo.Balance.String())
			if rstFormInfo.Balance.LessThan(needMin) {
				return nil, nil, fmt.Errorf("energyTrxInsufficient info err:%v", err)
			}
			energyRemaining := decimal.Zero

			if rstFormInfo.Bandwidth != nil {
				if rstFormInfo.Bandwidth.EnergyRemaining.GreaterThan(energyRemaining) {
					energyRemaining = rstFormInfo.Bandwidth.EnergyRemaining.Copy()
				}
			}

			if energyRemaining.LessThan(needEnergy) {
				txEtHash, err := transferInfo.TRXSend(outNeedTrx.IntPart(), privateKeyCfg.FromAddress, privateKeyCfg.PrivateKey, privateKeyCfg.ExtraCfg.ExchangeAddress)
				if err != nil {
					return nil, nil, fmt.Errorf("buy energy send err:%v", err)
				}
				//等待能量到账
				for i := 0; i < 30; i++ {
					tmr := time.After(4 * time.Second)
					<-tmr
					rstFormInfo2, err := tronscanapi.GetAccountDetailV2(apiCfgScan.Url, apiCfgScan.AppID, &tronscanapi.AccountDetailV2Req{
						Address: privateKeyCfg.FromAddress,
					})
					if err != nil {
						log.Warnf("TRC20 %v", err)
						continue
					}

					if rstFormInfo2.Bandwidth != nil {
						if rstFormInfo2.Bandwidth.EnergyRemaining.GreaterThan(energyRemaining) {
							energyRemaining = rstFormInfo2.Bandwidth.EnergyRemaining.Copy()
							break
						}
					}
				}
				energyExtra = &xtypes.EnergyExtra{
					ExchangeAddress: txEtHash,                               //兑换能量地址
					ExchangeTXHash:  privateKeyCfg.ExtraCfg.ExchangeAddress, //交易Hash
					ToBalance:       toCurryInfo.Balance.Copy(),             //是否需要能量
					ExchangePrice:   outNeedTrx.Copy(),                      //有U的价格

				}
			}
			log.Warnf("TRC20 needEnergy:%s energyRemaining:%s", needEnergy.String(), energyRemaining.String())
			if energyRemaining.GreaterThanOrEqual(needEnergy) {
				txID, err := transferInfo.TRC20Send(outAmout.IntPart(), privateKeyCfg.FromAddress, privateKeyCfg.PrivateKey, toAddress, contract, maxGasFeeLimmit.IntPart())
				if err != nil {
					return nil, energyExtra, err
				}
				return &TransferOutRes{
					Amount:     amount.InexactFloat64(),
					TxID:       txID,
					OutAddress: privateKeyCfg.FromAddress,
				}, energyExtra, nil
			}

			return nil, energyExtra, fmt.Errorf("energy insufficient needEnergy:%s energyRemaining:%s", needEnergy.String(), energyRemaining.String())
		} else {
			needTRx := maxGasFeeLimmit.Add(tronWalletNetFee)
			if rstFormInfo.Balance.LessThanOrEqual(maxGasFeeLimmit.Add(tronWalletNetFee)) {
				return nil, nil, fmt.Errorf("energy insufficient needTRx:%s energyRemaining:%s", needTRx.String(), rstFormInfo.Balance.String())
			}
			fromTcy := rstFormInfo.GetWithPriceTokensByTokenId(contract)
			if fromTcy == nil {
				return nil, nil, fmt.Errorf("from not have contract")
			}
			needContract := outAmout.Add(walletMustBlance)
			if fromTcy.Balance.LessThan(needContract) {
				return nil, nil, fmt.Errorf("energy insufficient needContract:%s energyRemaining:%s", needContract.String(), fromTcy.Balance.String())
			}
			txID, err := transferInfo.TRC20Send(outAmout.IntPart(), privateKeyCfg.FromAddress, privateKeyCfg.PrivateKey, toAddress, contract, maxGasFeeLimmit.IntPart())
			if err != nil {
				return nil, nil, err
			}
			return &TransferOutRes{
				Amount:     amount.InexactFloat64(),
				TxID:       txID,
				OutAddress: privateKeyCfg.FromAddress,
			}, nil, nil
		}

	}

	return nil, nil, fmt.Errorf("not support contract")
}

func (to *TransferOut) DelegateResourceEnegy(toAddress string, amount decimal.Decimal) (*TransferOutRes, *xtypes.EnergyExtra, error) {
	to.mux.Lock()
	defer to.mux.Unlock()

	privateKeyCfg, err := optionCurrencyNetworkCfg.GetPrivateKeyCfg(rechargeChannelType, xtypes.ENERGY)
	if err != nil {
		return nil, nil, fmt.Errorf("privateKeyCfg is not found:%v", err)
	}
	if privateKeyCfg == nil {
		return nil, nil, fmt.Errorf("privateKeyCfg is not found")
	}
	if privateKeyCfg.FromAddress == "" || privateKeyCfg.PrivateKey == "" {
		return nil, nil, fmt.Errorf("fromAddress or privateKey is not found")
	}
	apiCfg := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(rechargeChannelType, xtypes.APITrongrid)
	if apiCfg == nil {
		return nil, nil, fmt.Errorf("apiCfg is not found")
	}
	/*
		apiCfgScan := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(rechargeChannelType, xtypes.APITronscan)
		if apiCfgScan == nil {
			return nil, nil, fmt.Errorf("apiCfgScan is not found")
		}
	*/
	transferInfo := transfer.NewTransfer(apiCfg.Url, apiCfg.AppID)

	if transferInfo == nil {
		return nil, nil, fmt.Errorf("transferInfo new failed")
	}

	txHash, err := transferInfo.DelegateResourceEnegy(privateKeyCfg.FromAddress, privateKeyCfg.PrivateKey, toAddress, 2000000)
	if err != nil {
		log.Warnf("%v", err)
		return nil, nil, fmt.Errorf("transferInfo new failed err:%v", err)
	}

	log.Warnf("%s", txHash)
	return &TransferOutRes{
		Amount:     amount.InexactFloat64(),
		TxID:       txHash,
		OutAddress: privateKeyCfg.FromAddress,
	}, nil, nil
}
func (to *TransferOut) UnDelegateResource(toAddress string, amount decimal.Decimal) (*TransferOutRes, *xtypes.EnergyExtra, error) {
	to.mux.Lock()
	defer to.mux.Unlock()

	privateKeyCfg, err := optionCurrencyNetworkCfg.GetPrivateKeyCfg(rechargeChannelType, xtypes.ENERGY)
	if err != nil {
		return nil, nil, fmt.Errorf("privateKeyCfg is not found:%v", err)
	}
	if privateKeyCfg == nil {
		return nil, nil, fmt.Errorf("privateKeyCfg is not found")
	}
	if privateKeyCfg.FromAddress == "" || privateKeyCfg.PrivateKey == "" {
		return nil, nil, fmt.Errorf("fromAddress or privateKey is not found")
	}
	apiCfg := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(rechargeChannelType, xtypes.APITrongrid)
	if apiCfg == nil {
		return nil, nil, fmt.Errorf("apiCfg is not found")
	}
	/*
		apiCfgScan := optionCurrencyNetworkCfg.Instance().GetNeedApiByChannelType(rechargeChannelType, xtypes.APITronscan)
		if apiCfgScan == nil {
			return nil, nil, fmt.Errorf("apiCfgScan is not found")
		}
	*/
	transferInfo := transfer.NewTransfer(apiCfg.Url, apiCfg.AppID)

	if transferInfo == nil {
		return nil, nil, fmt.Errorf("transferInfo new failed")
	}

	txHash, err := transferInfo.UnDelegateResource(privateKeyCfg.FromAddress, privateKeyCfg.PrivateKey, toAddress, 2000000)
	if err != nil {
		log.Warnf("%v", err)
		return nil, nil, fmt.Errorf("transferInfo new failed err:%v", err)
	}

	log.Warnf("%s", txHash)
	return &TransferOutRes{
		Amount:     amount.InexactFloat64(),
		TxID:       txHash,
		OutAddress: privateKeyCfg.FromAddress,
	}, nil, nil
}

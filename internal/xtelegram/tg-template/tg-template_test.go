package tgtemplate_test

import (
	"fmt"
	"testing"
	"xbase/log"
	"xbase/utils/xconv"
	"xrobot/internal/utils/expand"
	tgtemplate "xrobot/internal/xtelegram/tg-template"

	"github.com/shopspring/decimal"
)

func TestClient_Start(t *testing.T) {
	str := expand.Expand(tgtemplate.Start(), map[string]string{
		tgtemplate.CustomerKey:      "@888888888",
		tgtemplate.EnergySavingsKey: "999999999999999%",
	})
	log.Warnf("%s", str)
}
func TestClient_EnergyFlashRental(t *testing.T) {

	expandMap := map[string]string{
		tgtemplate.EnergySavingsKey: "%80",
		tgtemplate.PriceKey:         "3",
		tgtemplate.PriceNoUKey:      "6",
		tgtemplate.PriceBiShuMaxKey: "10",
		tgtemplate.Tron20AddressKey: "bbbbbbbbbbbbbbbb",
	}
	priceHaveU := decimal.NewFromFloat(3)
	for i := 1; i <= 3; i++ {
		priceNumKey := fmt.Sprintf(tgtemplate.PriceNumIndexKey, i)
		price := priceHaveU.Mul(decimal.NewFromInt(int64(i)))
		expandMap[priceNumKey] = price.String()

		priceBiShuKey := fmt.Sprintf(tgtemplate.PriceBiShuKey, i)
		expandMap[priceBiShuKey] = xconv.String(i)

	}
	str := expand.Expand(tgtemplate.EnergyFlashRental(), expandMap)
	log.Warnf("%s", str)
}

func TestClient_EnergyFlashRentalBiShu(t *testing.T) {

	expandMap := map[string]string{
		tgtemplate.PriceNumKey:      "100000",
		tgtemplate.PriceKey:         "3",
		tgtemplate.Tron20AddressKey: "bbbbbbbbbbbbbbbb",
	}

	str := expand.Expand(tgtemplate.EnergyFlashRentalBiShu(), expandMap)
	log.Warnf("%s", str)
}

func TestClient_RechargeOtherAddressesRet(t *testing.T) {

	expandMap := map[string]string{
		tgtemplate.ComboKindEnergyFlashRentalNumKey:  "1",
		tgtemplate.ComboKindEnergyFlashRentalNameKey: "小时",
		tgtemplate.PriceNumKey:                       "10",
		tgtemplate.NotActivatedAddressCountKey:       "1",
		tgtemplate.ReceivingAddressCountKey:          "2",
		tgtemplate.EnergyFeeKey:                      "3",
		tgtemplate.ActivationfeeKey:                  "4",
		tgtemplate.PayAmountKey:                      "5",
		tgtemplate.Tron20AddressKey:                  "bbbbbbbbbbbbbbbb",
	}

	str := expand.Expand(tgtemplate.RechargeOtherAddressesRet(), expandMap)
	log.Warnf("%s", str)
}

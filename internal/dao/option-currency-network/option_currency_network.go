package optioncurrencynetwork

import (
	"context"
	"sync"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	"tron_robot/internal/dao/option-currency-network/internal"
	modelpkg "tron_robot/internal/model"
	"tron_robot/internal/xtypes"
	"xbase/utils/xtime"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type (
	Columns    = internal.Columns
	OrderBy    = internal.OrderBy
	FilterFunc = internal.FilterFunc
	UpdateFunc = internal.UpdateFunc
	ColumnFunc = internal.ColumnFunc
	OrderFunc  = internal.OrderFunc
)

type OptionCurrencyNetwork struct {
	*internal.OptionCurrencyNetwork
}

func NewOptionCurrencyNetwork(db *gorm.DB) *OptionCurrencyNetwork {
	return &OptionCurrencyNetwork{OptionCurrencyNetwork: internal.NewOptionCurrencyNetwork(db)}
}

var (
	once     sync.Once
	instance *OptionCurrencyNetwork
)

func Instance() *OptionCurrencyNetwork {
	once.Do(func() {
		instance = NewOptionCurrencyNetwork(mysqlimp.Instance())
	})
	return instance
}
func (dao *OptionCurrencyNetwork) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionCurrencyNetwork{})
		if err != nil {
			panic(err)
		}
		dao.doInit()
	}
	return nil
}
func (dao *OptionCurrencyNetwork) doInit() {
	cxt := context.Background()
	now := xtime.Now()

	solana := xtypes.PrivateKeyCfgList{}
	dao.Insert(cxt, &modelpkg.OptionCurrencyNetwork{
		Type:    xtypes.NetWorkChannelType_Solana,           //类型 1:EVM 2:Solana 3:Ton
		NetWork: xtypes.NetWorkChannelType_Solana.NetWork(), //充值渠道名
		ApiCfg: xtypes.ApiCfgList{
			xtypes.APISolana: &xtypes.ApiCfgListInfo{
				DurationKind: xtypes.DurationKind_Daily,
				Cfg: xtypes.KeyToApiCfgInfo{
					"api.mainnet-beta.solana.com": {
						Url:         "https: //api.mainnet-beta.solana.com", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                                     //周期内可以最大数量
						//周期时长 秒
						AppID:  "", //appID
						Secret: "", //SignatureKey

					},
					"solana-mainnet.g.alchemy.com": {
						Url:         "https: //solana-mainnet.g.alchemy.com/v2/cVPXctqANDK4DdkvmW1mp1cumr_E-TJm", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                                                                          //周期内可以最大数量
						//周期时长 秒
						AppID:  "", //appID
						Secret: "", //SignatureKey

					},
				},
			},
		}, //连接地址
		PrivateKey:    solana.ToJsonByte(), //加密数据
		DecimalPlaces: 9,
		Status:        xtypes.OptionStatus_Normal, //状态（1=启用;2=停用）
		CreateAt:      now,                        //创建时间
		UpdateAt:      now,                        //更新时间
	})
	ton := xtypes.PrivateKeyCfgList{}
	dao.Insert(cxt, &modelpkg.OptionCurrencyNetwork{
		Type:          xtypes.NetWorkChannelType_Ton,           //类型 1:EVM 2:Solana 3:Ton
		NetWork:       xtypes.NetWorkChannelType_Ton.NetWork(), //充值渠道名
		ApiCfg:        nil,                                     //连接地址,                                     //连接地址
		PrivateKey:    ton.ToJsonByte(),                        //加密数据
		DecimalPlaces: 9,
		Status:        xtypes.OptionStatus_Normal, //状态（1=启用;2=停用）
		CreateAt:      now,                        //创建时间
		UpdateAt:      now,                        //更新时间
	})
	tonCrypto := xtypes.PrivateKeyCfgList{}
	dao.Insert(cxt, &modelpkg.OptionCurrencyNetwork{
		Type:          xtypes.NetWorkChannelType_TonCrypto,           //类型 1:EVM 2:Solana 3:Ton
		NetWork:       xtypes.NetWorkChannelType_TonCrypto.NetWork(), //充值渠道名
		ApiCfg:        nil,                                           //连接地址
		PrivateKey:    tonCrypto.ToJsonByte(),                        //加密数据
		DecimalPlaces: 9,
		Status:        xtypes.OptionStatus_Normal, //状态（1=启用;2=停用）
		CreateAt:      now,                        //创建时间
		UpdateAt:      now,                        //更新时间
	})
	tron := xtypes.PrivateKeyCfgList{

		PrivateKeyCfg: map[xtypes.Currency]*xtypes.PrivateKeyCfg{
			xtypes.USDT: {

				FromAddress: "",       // trc20转转账地址
				PrivateKey:  "",       // 私钥
				MaxFeeLimit: 45000000, //最高手续续
				AppID:       "",
				Testing:     false,
				Status:      xtypes.OptionStatus_Normal,
				ExtraCfg: &xtypes.PlatformExtraCfg{
					ExchangeAddress: "TKwiPBgVkr44egdUyCP1VLNSUEyxPTLepu", //兑换能量地址
					PriceU:          decimal.NewFromFloat(2),              //有U的价格
					PriceNoU:        decimal.NewFromFloat(4),              //无U的价格
					EnergyU:         decimal.NewFromFloat(65000),          //有U的需要能量
					EnergyNoU:       decimal.NewFromFloat(131000),         //无U的需要能量
					IsUserEnergy:    true,                                 //是否需要能量
				},
			},
		},
	}
	dao.Insert(cxt, &modelpkg.OptionCurrencyNetwork{
		Type:    xtypes.NetWorkChannelType_TRON,           //类型 1:EVM 2:Solana 3:Ton
		NetWork: xtypes.NetWorkChannelType_TRON.NetWork(), //充值渠道名
		ApiCfg: xtypes.ApiCfgList{
			xtypes.APITrongrid: &xtypes.ApiCfgListInfo{
				DurationKind: xtypes.DurationKind_Daily,
				Cfg: xtypes.KeyToApiCfgInfo{
					"a3d05166-6cd9-41d6-8781-f88849067c37": {
						Url:         "grpc.trongrid.io:50051", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                       //周期内可以最大数量
						//周期时长 秒
						AppID:  "a3d05166-6cd9-41d6-8781-f88849067c37", //appID
						Secret: "",                                     //SignatureKey

					},
					"e97dda44-43d2-4f7c-ac77-7b8526fe0615": {
						Url:         "grpc.trongrid.io:50051", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                       //周期内可以最大数量
						//周期时长 秒
						AppID:  "e97dda44-43d2-4f7c-ac77-7b8526fe0615", //appID
						Secret: "",                                     //SignatureKey

					},
				},
			},
			xtypes.APITronscan: &xtypes.ApiCfgListInfo{
				DurationKind: xtypes.DurationKind_Daily,
				Cfg: xtypes.KeyToApiCfgInfo{
					"13bd4333-960f-4ced-9131-91fcfb8d64a5": {
						Url:         "https://apilist.tronscanapi.com", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                                //周期内可以最大数量
						//周期时长 秒
						AppID:  "13bd4333-960f-4ced-9131-91fcfb8d64a5", //appID
						Secret: "",                                     //SignatureKey

					},
					"2ac702c6-77b1-4fb3-9e83-718402c88106": {
						Url:         "https://apilist.tronscanapi.com", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                                //周期内可以最大数量
						//周期时长 秒
						AppID:  "2ac702c6-77b1-4fb3-9e83-718402c88106", //appID
						Secret: "",                                     //SignatureKey

					},
				},
			},
		}, //连接地址
		PrivateKey:    tron.ToJsonByte(), //加密数据
		DecimalPlaces: 6,
		Status:        xtypes.OptionStatus_Normal, //状态（1=启用;2=停用）
		CreateAt:      now,                        //创建时间
		UpdateAt:      now,                        //更新时间
	})

	btc := xtypes.PrivateKeyCfgList{}
	dao.Insert(cxt, &modelpkg.OptionCurrencyNetwork{
		Type:    xtypes.NetWorkChannelType_BTC,           //类型 1:EVM 2:Solana 3:Ton
		NetWork: xtypes.NetWorkChannelType_BTC.NetWork(), //充值渠道名
		ApiCfg: xtypes.ApiCfgList{
			xtypes.APIGetblockIO: &xtypes.ApiCfgListInfo{
				DurationKind: xtypes.DurationKind_Daily,
				Cfg: xtypes.KeyToApiCfgInfo{
					"9a4af4b7b8634ad9bfc0971f4499272d": {
						Url:         "https://go.getblock.io", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                       //周期内可以最大数量
						//周期时长 秒
						AppID:  "9a4af4b7b8634ad9bfc0971f4499272d", //appID
						Secret: "",                                 //SignatureKey

					},
					"9831ebbad30440b8a8be046bcd859b5f": {
						Url:         "https://go.getblock.io", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                       //周期内可以最大数量
						//周期时长 秒
						AppID:  "9831ebbad30440b8a8be046bcd859b5f", //appID
						Secret: "",                                 //SignatureKey

					},
				},
			},
		}, //连接地址
		PrivateKey:    btc.ToJsonByte(), //加密数据
		DecimalPlaces: 18,
		Status:        xtypes.OptionStatus_Normal, //状态（1=启用;2=停用）
		CreateAt:      now,                        //创建时间
		UpdateAt:      now,                        //更新时间
	})
	eth := xtypes.PrivateKeyCfgList{}
	dao.Insert(cxt, &modelpkg.OptionCurrencyNetwork{
		Type:    xtypes.NetWorkChannelType_ETH,           //类型 1:EVM 2:Solana 3:Ton
		NetWork: xtypes.NetWorkChannelType_ETH.NetWork(), //充值渠道名
		ApiCfg: xtypes.ApiCfgList{
			xtypes.APIEtherscan: &xtypes.ApiCfgListInfo{
				DurationKind: xtypes.DurationKind_Daily,
				Cfg: xtypes.KeyToApiCfgInfo{
					"MCQUWV69GXCF37ZJFXUSYE7CFE35XGJCG1": {
						Url:         "https://api.etherscan.io", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                         //周期内可以最大数量
						//周期时长 秒
						AppID:  "MCQUWV69GXCF37ZJFXUSYE7CFE35XGJCG1", //appID
						Secret: "",                                   //SignatureKey

					},
					"3HIBH5WUFF8H34E5JSFDXNXBVX9M56SU55": {
						Url:         "https://api.etherscan.io", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                         //周期内可以最大数量
						//周期时长 秒
						AppID:  "3HIBH5WUFF8H34E5JSFDXNXBVX9M56SU55", //appID
						Secret: "",                                   //SignatureKey

					},
				},
			},
		}, //连接地址
		PrivateKey:    eth.ToJsonByte(), //加密数据
		DecimalPlaces: 18,
		Status:        xtypes.OptionStatus_Normal, //状态（1=启用;2=停用）
		CreateAt:      now,                        //创建时间
		UpdateAt:      now,                        //更新时间
	})

	bnb := xtypes.PrivateKeyCfgList{}
	dao.Insert(cxt, &modelpkg.OptionCurrencyNetwork{
		Type:    xtypes.NetWorkChannelType_BNB,           //类型 1:EVM 2:Solana 3:Ton
		NetWork: xtypes.NetWorkChannelType_BNB.NetWork(), //充值渠道名
		ApiCfg: xtypes.ApiCfgList{
			xtypes.APIBscscan: &xtypes.ApiCfgListInfo{
				DurationKind: xtypes.DurationKind_Daily,
				Cfg: xtypes.KeyToApiCfgInfo{
					"8NVI4IQTGA1PMZN2R58BRV93F1D4NDKV1M": {
						Url:         "https://api.bscscan.com", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                        //周期内可以最大数量
						//周期时长 秒
						AppID:  "8NVI4IQTGA1PMZN2R58BRV93F1D4NDKV1M", //appID
						Secret: "",                                   //SignatureKey

					},
					"1KX2AA9GYCT2CMG9M7ZBH4BUK9X6UKRV2R": {
						Url:         "https://api.bscscan.com", //API接口基础地址 或者网络(主要针对区块)
						DurationMax: -1,                        //周期内可以最大数量
						//周期时长 秒
						AppID:  "1KX2AA9GYCT2CMG9M7ZBH4BUK9X6UKRV2R", //appID
						Secret: "",                                   //SignatureKey

					},
				},
			},
		}, //连接地址
		PrivateKey:    bnb.ToJsonByte(), //加密数据
		DecimalPlaces: 18,
		Status:        xtypes.OptionStatus_Normal, //状态（1=启用;2=停用）
		CreateAt:      now,                        //创建时间
		UpdateAt:      now,                        //更新时间
	})
}

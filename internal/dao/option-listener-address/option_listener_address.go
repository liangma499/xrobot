package optionlisteneraddress

import (
	"context"
	"fmt"
	"sync"
	"tron_robot/internal/code"
	mysqlimp "tron_robot/internal/component/mysql/mysql-default"
	redisCryptoCurrencies "tron_robot/internal/component/redis/redis-crypto-currencies"
	"tron_robot/internal/dao/option-listener-address/internal"
	modelpkg "tron_robot/internal/model"
	"tron_robot/internal/xtypes"
	"xbase/errors"
	"xbase/log"
	"xbase/utils/xtime"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Columns    = internal.Columns
	OrderBy    = internal.OrderBy
	FilterFunc = internal.FilterFunc
	UpdateFunc = internal.UpdateFunc
	ColumnFunc = internal.ColumnFunc
	OrderFunc  = internal.OrderFunc
)

type OptionListenerAddress struct {
	*internal.OptionListenerAddress
}

func NewOptionListenerAddress(db *gorm.DB) *OptionListenerAddress {
	return &OptionListenerAddress{OptionListenerAddress: internal.NewOptionListenerAddress(db)}
}

var (
	once     sync.Once
	instance *OptionListenerAddress
)

func Instance() *OptionListenerAddress {
	once.Do(func() {
		instance = NewOptionListenerAddress(mysqlimp.Instance())
	})
	return instance
}
func (dao *OptionListenerAddress) CreateTable() error {
	table := dao.TableName
	if !dao.Table.Migrator().HasTable(table) {
		err := dao.Table.Migrator().AutoMigrate(&modelpkg.OptionListenerAddress{})
		if err != nil {
			panic(err)
		}
		dao.doInitCfg()
	}
	dao.doInitCfgToRedis()
	return nil
}
func (dao *OptionListenerAddress) doInitCfg() {
	now := xtime.Now()
	dao.Insert(context.Background(), &modelpkg.OptionListenerAddress{
		NetWork:     xtypes.TRON, //周期时长 秒
		Address:     "AAAAAAAAAAAAAAAAAAABiTTTTTTTT",
		ChannelCode: xtypes.OfficialChannelCode,
		AddressKind: xtypes.AddressKind_Official,
		Status:      xtypes.OptionStatus_Normal,
		OperateUid:  -1,
		OperateUser: "init",
		CreateAt:    now,
		UpdateAt:    now,
	})
}
func (dao *OptionListenerAddress) doInitCfgToRedis() {
	limit := 200
	page := 1
	ctx := context.Background()
	for {
		offest := (page - 1) * 200
		addresses, err := dao.FindMany(ctx, nil, nil, nil, limit, offest)
		if err != nil {
			panic(err)
		}
		if addresses == nil {
			break
		}
		for _, item := range addresses {
			if item.Status == xtypes.OptionStatus_Normal {
				if err = dao.AddListenerAddr(item.NetWork, item.Address); err != nil {
					panic(err)
				}
			} else {
				_ = dao.RemoveListenerAddr(item.NetWork, item.Address)
			}
		}
		if len(addresses) < limit {
			break
		}
		page++
	}
}
func (dao *OptionListenerAddress) AddListenerAddr(netWork xtypes.NetWork, address string) error {
	key := fmt.Sprintf(xtypes.ListenerAddrKey, netWork.String())

	_, err := redisCryptoCurrencies.Instance().SAdd(context.Background(), key, address).Result()
	return err
}
func (dao *OptionListenerAddress) RemoveListenerAddr(netWork xtypes.NetWork, address string) error {
	key := fmt.Sprintf(xtypes.ListenerAddrKey, netWork.String())

	_, err := redisCryptoCurrencies.Instance().SRem(context.Background(), key, address).Result()
	return err
}
func (dao *OptionListenerAddress) CheckListenerAddress(netWork xtypes.NetWork, address string) *modelpkg.OptionListenerAddress {
	key := fmt.Sprintf(xtypes.ListenerAddrKey, netWork.String())

	bHave, err := redisCryptoCurrencies.Instance().SIsMember(context.Background(), key, address).Result()
	if err != nil {
		return nil
	}
	if !bHave {
		return nil
	}
	cfg, err := dao.GetOptionListenerAddress(netWork, address)
	if err != nil {
		log.Errorf("%v", err)
		return nil
	}
	if cfg.Status != xtypes.OptionStatus_Normal {
		dao.RemoveListenerAddr(netWork, address)
		return nil
	}
	return cfg
}
func (dao *OptionListenerAddress) GetOptionListenerAddress(netWork xtypes.NetWork, address string) (*modelpkg.OptionListenerAddress, error) {
	key := fmt.Sprintf(xtypes.CacheNetworkAddresToCfgKey, netWork.String(), address)
	cfg := &modelpkg.OptionListenerAddress{}
	ctx := context.Background()
	err := redisCryptoCurrencies.InstanceCache().GetSetMux(ctx, key, func() (any, error) {
		return dao.FindOne(ctx, func(cols *internal.Columns) any {
			return clause.And
		})

	}, xtypes.CacheNetworkAddressToCfgExpiration).Scan(cfg)
	if err != nil {
		if errors.Is(err, errors.ErrNil) {
			return nil, nil
		}
		log.Errorf("get optionListenerAddress failed, apiKind = %d err = %v", netWork, address, err)
		return nil, errors.NewError(err, code.InternalError)
	}

	return cfg, nil
}
func (dao *OptionListenerAddress) GetAddressByChannelCode(channelCode string, netWork xtypes.NetWork) string {
	ctx := context.Background()
	/*if true {
		rt, err := dao.FindOne(ctx, func(cols *internal.Columns) any {
			return clause.And(
				clause.Eq{
					Column: cols.NetWork,
					Value:  netWork.String(),
				},
				clause.Eq{
					Column: cols.ChannelCode,
					Value:  channelCode,
				},
				clause.Eq{
					Column: cols.AddressKind,
					Value:  xtypes.AddressKind_Official,
				},
			)
		})
		if err != nil {
			return ""
		}
		if rt == nil {
			return ""
		}
		return rt.Address
	}
	*/
	key := fmt.Sprintf(xtypes.CacheNetworkChannelCodeToCfgKey, channelCode, netWork.String(), xtypes.AddressKind_Official)
	cfg := &modelpkg.OptionListenerAddress{}

	err := redisCryptoCurrencies.InstanceCache().GetSetMux(ctx, key, func() (any, error) {
		return dao.FindOne(ctx, func(cols *internal.Columns) any {
			return clause.And(
				clause.Eq{
					Column: cols.NetWork,
					Value:  netWork.String(),
				},
				clause.Eq{
					Column: cols.ChannelCode,
					Value:  channelCode,
				},
				clause.Eq{
					Column: cols.AddressKind,
					Value:  xtypes.AddressKind_Official,
				},
			)
		})

	}, xtypes.CacheNetworkAddressToCfgExpiration).Scan(cfg)
	if err != nil {
		if errors.Is(err, errors.ErrNil) {
			return dao.GetAddressByOfficialChannelCode(netWork)
		}
		log.Errorf("get optionListenerAddress failed,channelCode = %s apiKind = %v err = %v", channelCode, netWork, err)
		return ""
	}
	if cfg.Address == "" {
		return dao.GetAddressByOfficialChannelCode(netWork)
	}
	return cfg.Address
}
func (dao *OptionListenerAddress) GetAddressByOfficialChannelCode(netWork xtypes.NetWork) string {
	channelCode := xtypes.OfficialChannelCode
	key := fmt.Sprintf(xtypes.CacheNetworkChannelCodeToCfgKey, channelCode, netWork.String(), xtypes.AddressKind_Official)
	cfg := &modelpkg.OptionListenerAddress{}
	ctx := context.Background()
	err := redisCryptoCurrencies.InstanceCache().GetSetMux(ctx, key, func() (any, error) {
		return dao.FindOne(ctx, func(cols *internal.Columns) any {
			return clause.And(
				clause.Eq{
					Column: cols.NetWork,
					Value:  netWork.String(),
				},
				clause.Eq{
					Column: cols.ChannelCode,
					Value:  channelCode,
				},
				clause.Eq{
					Column: cols.AddressKind,
					Value:  xtypes.AddressKind_Official,
				},
			)
		})

	}, xtypes.CacheNetworkAddressToCfgExpiration).Scan(cfg)
	if err != nil {
		if errors.Is(err, errors.ErrNil) {
			return ""
		}
		log.Errorf("get optionListenerAddress failed,channelCode = %s apiKind = %v err = %v", channelCode, netWork, err)
		return ""
	}

	return cfg.Address
}

// 验证交易ID
func (dao *OptionListenerAddress) CheckListenerTrxID(netWork xtypes.NetWork, trxID string) bool {
	key := fmt.Sprintf(xtypes.ListenerTrxIDKey, netWork.String())

	bHave, err := redisCryptoCurrencies.Instance().SIsMember(context.Background(), key, trxID).Result()
	if err != nil {
		return false
	}
	//是真临交易
	return bHave

}
func (dao *OptionListenerAddress) AddListenerTrxID(netWork xtypes.NetWork, trxID string) error {
	key := fmt.Sprintf(xtypes.ListenerTrxIDKey, netWork.String())

	_, err := redisCryptoCurrencies.Instance().SAdd(context.Background(), key, trxID).Result()
	return err
}
func (dao *OptionListenerAddress) RemoveListenerTrxID(netWork xtypes.NetWork, trxID string) error {
	key := fmt.Sprintf(xtypes.ListenerTrxIDKey, netWork.String())

	_, err := redisCryptoCurrencies.Instance().SRem(context.Background(), key, trxID).Result()
	return err
}

// 验证用户转账收回能量代理权限
func (dao *OptionListenerAddress) CheckListenerTransactionAddress(netWork xtypes.NetWork, address string) bool {
	key := fmt.Sprintf(xtypes.ListenerTrxIDKey, netWork.String())

	bHave, err := redisCryptoCurrencies.Instance().SIsMember(context.Background(), key, address).Result()
	if err != nil {
		return false
	}
	//是真临交易
	return bHave

}
func (dao *OptionListenerAddress) AddListenerTransactionAddress(netWork xtypes.NetWork, address string) error {
	key := fmt.Sprintf(xtypes.ListenerTrxIDKey, netWork.String())

	_, err := redisCryptoCurrencies.Instance().SAdd(context.Background(), key, address).Result()
	return err
}
func (dao *OptionListenerAddress) RemoveListenerTransactionAddress(netWork xtypes.NetWork, address string) error {
	key := fmt.Sprintf(xtypes.ListenerTrxIDKey, netWork.String())

	_, err := redisCryptoCurrencies.Instance().SRem(context.Background(), key, address).Result()
	return err
}

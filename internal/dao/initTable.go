package dao

import (
	optionBaseConfigDao "xrobot/internal/dao/option-base-config"
	optionChannelDao "xrobot/internal/dao/option-channel"
	optionCurrencyDao "xrobot/internal/dao/option-currency"
	optionCurrencyChannelDao "xrobot/internal/dao/option-currency-channel"
	optionCurrencyNetworkDao "xrobot/internal/dao/option-currency-network"
	optionListenerAddressDao "xrobot/internal/dao/option-listener-address"
	optionTelegramCmdDao "xrobot/internal/dao/option-telegram-cmd"
	optionTelegramInlineKeyboardButtonDao "xrobot/internal/dao/option-telegram-inline-keyboard-button"
	optionTelegramKeyboardMarkupButtonDao "xrobot/internal/dao/option-telegram-keyboard-markup-button"
	optionWithdrawCurrencyDao "xrobot/internal/dao/option-withdraw-currency"
	paymentAmountUserDao "xrobot/internal/dao/payment-amount-user"
	paymentAmountUserRecordDao "xrobot/internal/dao/payment-amount-user-record"
	paymentCryptoTransactionDao "xrobot/internal/dao/payment-crypto-transaction"
	paymentNetworkDao "xrobot/internal/dao/payment-network"
	userDao "xrobot/internal/dao/user"
	userBalanceDepositTransactionDao "xrobot/internal/dao/user-balance-deposit-transaction"
	userBaseDao "xrobot/internal/dao/user-base"
	userCommissionDao "xrobot/internal/dao/user-commission"
	userCommissionRecordDao "xrobot/internal/dao/user-commission-record"
	userLoginDayStatDao "xrobot/internal/dao/user-login-day-stat"
	userLoginLogDao "xrobot/internal/dao/user-login-log"
	userParentDao "xrobot/internal/dao/user-parent"
	userTradeDao "xrobot/internal/dao/user-trade"
	userWalletDao "xrobot/internal/dao/user-wallet"
	userWithdrawRecordDao "xrobot/internal/dao/user-withdraw-record"
)

func InitTableUser() {
	//用户表
	userDao.Instance().CreateTable()
	//用户充值记录
	userBalanceDepositTransactionDao.Instance().CreateTable()
	//用户基础表
	userBaseDao.Instance().CreateTable()
	//用户登录记录表
	userLoginLogDao.Instance().CreateTable()
	//用户返佣表
	userCommissionDao.Instance().CreateTable()
	//用户返佣记录表
	userCommissionRecordDao.Instance().CreateTable()
	//用户每日登录统计
	userLoginDayStatDao.Instance().CreateTable()
	//用户上下级关系
	userParentDao.Instance().CreateTable()
	//用户提现记录表
	userWithdrawRecordDao.Instance().CreateTable()
	//用户唯一金币表
	paymentAmountUserDao.Instance().CreateTable()
	//用户唯一金币表处理记录表
	paymentAmountUserRecordDao.Instance().CreateTable()
	//用户钱包表
	userWalletDao.Instance().CreateTable()
	//用户交易表
	userTradeDao.Instance().CreateTable()

}

// 初始化配置相关的表
func InitTableOption() {
	//基础配置表
	optionBaseConfigDao.Instance().CreateTable()
	//渠道配置表
	optionChannelDao.Instance().CreateTable()
	//GAS配置
	optionCurrencyChannelDao.Instance().CreateTable()
	//GAS配置
	optionCurrencyDao.Instance().CreateTable()
	//telegram 命令配置
	optionTelegramCmdDao.Instance().CreateTable()
	//提现配置
	optionWithdrawCurrencyDao.Instance().CreateTable()
	//内联按钮配置
	optionTelegramInlineKeyboardButtonDao.Instance().CreateTable()
	//菜单
	optionTelegramKeyboardMarkupButtonDao.Instance().CreateTable()
	//用户地址监听表
	optionListenerAddressDao.Instance().CreateTable()
	//用户渠道表
	optionCurrencyNetworkDao.Instance().CreateTable()
}

// 支付需要的表到center中去启动
func InitPaymentWork() {
	//数据拉取记录
	paymentNetworkDao.Instance().CreateTable()
	//交易记录
	paymentCryptoTransactionDao.Instance().CreateTable()

}

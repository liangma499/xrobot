package dao

import (
	optionBaseConfigDao "tron_robot/internal/dao/option-base-config"
	optionChannelDao "tron_robot/internal/dao/option-channel"
	optionCurrencyDao "tron_robot/internal/dao/option-currency"
	optionCurrencyChannelDao "tron_robot/internal/dao/option-currency-channel"
	optionCurrencyNetworkDao "tron_robot/internal/dao/option-currency-network"
	optionListenerAddressDao "tron_robot/internal/dao/option-listener-address"
	optionTelegramCmdDao "tron_robot/internal/dao/option-telegram-cmd"
	optionTelegramInlineKeyboardButtonDao "tron_robot/internal/dao/option-telegram-inline-keyboard-button"
	optionTelegramKeyboardMarkupButtonDao "tron_robot/internal/dao/option-telegram-keyboard-markup-button"
	optionWithdrawCurrencyDao "tron_robot/internal/dao/option-withdraw-currency"
	paymentAmountUserDao "tron_robot/internal/dao/payment-amount-user"
	paymentAmountUserRecordDao "tron_robot/internal/dao/payment-amount-user-record"
	paymentCryptoTransactionDao "tron_robot/internal/dao/payment-crypto-transaction"
	paymentNetworkDao "tron_robot/internal/dao/payment-network"
	userDao "tron_robot/internal/dao/user"
	userBalanceDepositTransactionDao "tron_robot/internal/dao/user-balance-deposit-transaction"
	userBaseDao "tron_robot/internal/dao/user-base"
	userCommissionDao "tron_robot/internal/dao/user-commission"
	userCommissionRecordDao "tron_robot/internal/dao/user-commission-record"
	userLoginDayStatDao "tron_robot/internal/dao/user-login-day-stat"
	userLoginLogDao "tron_robot/internal/dao/user-login-log"
	userParentDao "tron_robot/internal/dao/user-parent"
	userTradeDao "tron_robot/internal/dao/user-trade"
	userWalletDao "tron_robot/internal/dao/user-wallet"
	userWithdrawRecordDao "tron_robot/internal/dao/user-withdraw-record"
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

package userwallet

// 增加账户余额脚本
const incrBalanceScript = `
	if #KEYS < 2 then
		return {"-1"}
	end
	local changeKey = "cash"
	local usedKey = "used"

	local resetZero = false


	local rate = tonumber(KEYS[2])
	if rate <= 0 then
		return {"-3"}
	end
	local redisKey = KEYS[1]
	local exist = redis.call('EXISTS', redisKey)
	if exist == 0 then
		return {"1"}
	end


	local cashNum = ARGV[1] and tonumber(ARGV[1]) or 0
	local changeCash = math.floor(cashNum * rate)

	local cash = redis.call('HINCRBY', redisKey, changeKey, changeCash)
	if tonumber(cash) < 0 then
		usedCash = redis.call('HINCRBY', redisKey, changeKey, math.abs(changeCash))
		return {"2"}
	end
	local usedCash  = 0
	if cashNum < 0 then
		usedCash = redis.call('HINCRBY', redisKey, usedKey, math.abs(changeCash))
	end

	local def = redis.call('HGET', redisKey, 'def')
	usedCash = redis.call('HGET', redisKey, usedKey)
	usedCash = tonumber(usedCash)/rate
	local after = tonumber(cash)/rate
	local before = after - cashNum
	if before < 0.00000001 then
		before = 0
	end
	-- 设置key 两个星期过期
	redis.call('Expire', redisKey, 1209600)

	return {"3",tostring(before), tostring(after), tostring(cashNum), tostring(def),tostring(usedCash),changeKey}
`

//UPDATE user_wallet SET cash = cash*0.000000001;

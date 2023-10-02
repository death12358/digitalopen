-- probability.jackpot.pull.lua
-- database: 2
-- jackpot data model(HASH) example:
--	KEY: 	{WORKSPACE.<workspace>:POOLNAME.<pool_code>|JACKPOT} 
--	VALUE: 	[
--	key{<jackpotname>}: value{<jackpot_point>},
--]
-- INPUT: 
--  EVALSHA <script_sha1> 4 <workspace> <brand> <currency> <pool_code> <point> <default_point> 
-- OUTPUT:
--  <point>

-- Choose database
redis.call('SELECT', 2)

-- Get Keys
if #KEYS < 4 then
    return redis.error_reply('ILLEGAL_KEYS')
end
local WORKSPACE = tostring(KEYS[1])
local BRAND = tostring(KEYS[2])
local CURRENCY = tostring(KEYS[3])
local POOLNAME = tostring(KEYS[4])

-- Get Args
if #ARGV < 2 then
    return redis.error_reply('ILLEGAL_ARGV')
end
local POINT = tonumber(ARGV[1])
local DEFAULT_POINT = tonumber(ARGV[2])

-- Generate the jackpot key
local jackpot_key = string.format('WORKSPACE.%s:BRAND.%s:CURRENCY.%s|JACKPOT', WORKSPACE, BRAND, CURRENCY)
-- Generate the jackpot vkey
local jackpot_vkey = string.format('POOLNAME.%s', POOLNAME)

-- HGET 
local jackpot_point = redis.call('HGET', jackpot_key, jackpot_vkey)
jackpot_point = jackpot_point - POINT

-- Generate the result
local result = {
	status = 'OK',
	point = POINT,
	pool = jackpot_point,
}

-- Is the balance sufficient
if jackpot_point < 0 then
    result.status = 'INSUFFICIENT_BALANCE'
    return cjson.encode(result)
end

-- Is the balance sufficient
if jackpot_point < DEFAULT_POINT then
    jackpot_point = DEFAULT_POINT
end

-- HSET
redis.call('HSET', jackpot_key, jackpot_vkey, jackpot_point)

return cjson.encode(result)

-- probability.jackpot.peek.lua
-- database: 2
-- jackpot data model(HASH) example:
--	KEY: 	{WORKSPACE.<workspace>:BRAND.<brand>:CURRENCY:<currency>|JACKPOT} 
--	VALUE: 	[
--	key{<pool_code>}: value{<jackpot_point>},
--]
-- INPUT: 
--  EVALSHA <script_sha1> 4 <workspace> <brand> <currency> <pool_code>
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
local POOLCODE = tostring(KEYS[4])

-- Generate the jackpot key
local jackpot_key = string.format('WORKSPACE.%s:BRAND.%s:CURRENCY.%s|JACKPOT', WORKSPACE, BRAND, CURRENCY)
-- Generate the jackpot vkey
local jackpot_vkey = string.format('POOLCODE.%s', POOLCODE)

-- HGET
local jackpot_point = redis.call('HGET', jackpot_key, jackpot_vkey)

local result = {
    status = 'OK',
    jackpot_pool = jackpot_point,
}

return cjson.encode(result)

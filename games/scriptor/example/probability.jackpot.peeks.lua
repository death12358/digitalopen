-- probability.jackpot.peeks.lua
-- database: 2
-- jackpot data model(HASH) example:
--	KEY: 	{WORKSPACE.<workspace>:POOLCODE.<pool_code>|JACKPOT} 
--	VALUE: 	[
--]
-- INPUT: 
--  EVALSHA <script_sha1> 3 <workspace> <brand> <currency> <pool_code>...
-- OUTPUT:
--  <point>

-- gets all fields from a hash as a dictionary
local hgetall = function (key)
  local bulk = redis.call('HGETALL', key)
	local result = {}
	local nextkey
	for i, v in ipairs(bulk) do
		if i % 2 == 1 then
			nextkey = v
		else
			result[nextkey] = v
		end
	end
	return result
end

-- Choose database
redis.call('SELECT', 2)

-- Get Keys
if #KEYS < 3 then
    return redis.error_reply('ILLEGAL_KEYS')
end
local WORKSPACE = tostring(KEYS[1])
local BRAND = tostring(KEYS[2])
local CURRENCY = tostring(KEYS[3])

-- Generate the jackpot key
local jackpot_key = string.format('WORKSPACE.%s:BRAND.%s:CURRENCY.%s|JACKPOT', WORKSPACE, BRAND, CURRENCY)

-- HGET
local pools = hgetall(jackpot_key)

local result = {
    status = 'OK',
    pools = pools,
}

return cjson.encode(result)

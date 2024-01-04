-- probability.jackpot.currency.peeks.lua
-- database: 2
-- jackpot data model(HASH) example:
--	KEY: 	{WORKSPACE.<workspace>:POOLCODE.<pool_code>|JACKPOT} 
--	VALUE: 	[
--]
-- INPUT: 
--  EVALSHA <script_sha1> 2 <workspace> <brand>
-- OUTPUT:

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
if #KEYS < 2 then
    return redis.error_reply('ILLEGAL_KEYS')
end
local WORKSPACE = tostring(KEYS[1])
local BRAND = tostring(KEYS[2])

-- Generate the jackpot key
local jackpot_key = string.format('WORKSPACE.%s:BRAND.%s:CURRENCY.*|JACKPOT', WORKSPACE, BRAND)

local jp_keys = redis.call('KEYS', jackpot_key)

local pools = {}
for i, v in ipairs(jp_keys) do
	local pool_code = string.match(v, 'CURRENCY.(.*)|JACKPOT')
    pools[pool_code] = hgetall(v)
end

local result = {
    status = 'OK',
    pools = pools,
}

return cjson.encode(result)

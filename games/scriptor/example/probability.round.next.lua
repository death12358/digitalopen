-- probability.round.next.lua
-- database: 3
-- round data model(HASH) example:
--	KEY: 	{WORKSPACE.<workspace>:GAME.<game_code>|ROUND} 
--	VALUE: 	[
--	key{<round_id>:<user_id>}: value{<round>},
--]
-- 
-- INPUT: 	
--  EVALSHA  <script_sha1> 2 <workspace> <game_code> <round_id> <user_id>
-- OUTPUT:
-- 	json{
--		"status": "ok",
--		"round": <round>,
--  }
-- Status: 	
-- 	"OK"			- success
--  "ERROR"			- error
-- 	"EXIST"			- exist
-- 	"NOT_EXIST"		- not exist
-- 	"NOT_ACTIVE"	- not active
-- 	"NOT_STARTED"	- not started
--  "ENDROUND"		- end round
if #KEYS < 2 then
    return redis.error_reply('ILLEGAL_KEYS')
end

-- Get Keys
local WORKSPACE = tostring(KEYS[1])
local GAMECODE = tostring(KEYS[2])

-- Get Args
local ROUNDID = tostring(ARGV[1])            
local USERID = tostring(ARGV[2])

-- Choose database
redis.call('SELECT', 3)

-- Initialize the result table
local result = {
    status = 'NOT_ACTIVE',
    round = nil,
}

-- Generate the round key
local round_key = string.format('WORKSPACE.%s:GAME.%s|ROUND', WORKSPACE, GAMECODE)
local round_vkey = string.format('ROUNDID.%s:USERID.%s', ROUNDID, USERID)

-- Add code here to handle the case where the round does not exist
-- Check if the round exists in the hash table
local round_exist = redis.call('HEXISTS', round_key, round_vkey)
if round_exist == 0 then
    result.status = 'NOT_EXIST'
    return cjson.encode(result)
end

 -- Get the round from the hash table
local ROUND_VALUE = cjson.decode(redis.call('HGET', round_key, round_vkey))

-- Get the round result
local ROUND_RESULT = ROUND_VALUE['result']

-- Check if the round is rest
-- if ROUND_VALUE['position'] & 16 == 16 then
if bit.band(ROUND_VALUE['position'], 0x0010) == 0x0010 then
    ROUND_VALUE['position'] = bit.band(ROUND_VALUE['position'], bit.bnot(0x8000))
    result.status = 'OK'
    result.round = ROUND_VALUE
    return cjson.encode(result)
end

-- Decode the round
ROUND_VALUE['stages'] = ROUND_VALUE['stages'] + 1 
local CURRENT_STAGE = tostring(ROUND_VALUE['stages'])

-- Check if the round result exists
if ROUND_RESULT[CURRENT_STAGE] == nil then
    result.status = 'NOT_EXIST'

    --  HDEL round
    redis.call('HDEL', round_key, round_vkey)
    return cjson.encode(result)
end

-- Get current stage result
local ROUND_RESULT_CASE = ROUND_RESULT[CURRENT_STAGE]

-- Update the round result case
ROUND_RESULT_CASE['case'] = bit.band(ROUND_RESULT_CASE['case'], bit.bnot(0x8000))
if tonumber(ROUND_RESULT_CASE['point']) > 0 then
    ROUND_RESULT_CASE['case'] = bit.bor(ROUND_RESULT_CASE['case'], 0x0001)
end

-- Check if the round is over
result.status = 'OK'
-- Update the round result
ROUND_RESULT[CURRENT_STAGE] = ROUND_RESULT_CASE
ROUND_VALUE['result'] = ROUND_RESULT

-- next stage
local NEXT_STAGE = tostring(ROUND_VALUE['stages'] + 1)

-- player is rest
if bit.band(ROUND_RESULT_CASE['case'], 0x0010) == 0x0010 then
    -- ROUND_VALUE['position'] | 16 | 32768
    ROUND_RESULT_CASE['case'] = bit.band(ROUND_RESULT_CASE['case'], bit.bnot(0x0010))
    ROUND_VALUE['position'] = bit.bor(ROUND_VALUE['position'], 0x0010)
    ROUND_VALUE['position'] = bit.bor(ROUND_VALUE['position'], 0x8000)
elseif ROUND_RESULT[NEXT_STAGE] == nil then
       -- The game is over
    ROUND_VALUE['position'] = 0x0
    --  HDEL round
    redis.call('HDEL', round_key, round_vkey)
    -- Update the round
    result.round = ROUND_VALUE
    return cjson.encode(result)
end
-- Update the round result
ROUND_RESULT[CURRENT_STAGE] = ROUND_RESULT_CASE
ROUND_VALUE['result'] = ROUND_RESULT

-- Update the round status & set round.total_point is 0 & return round.
redis.call('HSET', round_key, round_vkey, cjson.encode(ROUND_VALUE))

-- Update the round
result.round = ROUND_VALUE

-- The game is not over yet, so the round.total_point is set to 0.
ROUND_VALUE['total_point'] = 0
return cjson.encode(result)

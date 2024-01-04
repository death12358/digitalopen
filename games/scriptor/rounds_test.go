package scriptor_test

import (
	"strconv"
	"testing"

	"games"
	"games/scriptor"
)

var (
	endcast = `
	{"id":"1234567890","game_code":"SG001","brand":"brand_test","username":"user_test","status":0,"position":0,"stages":0,"result":{"0":{"id":"1234567890","brand":"brand_test","username":"user_test","case":0,"stages":0,"pickem":["8"],"symbols":["3","2","4","0","1","2","10","1","4","7","2","11","5","4","10","6","8","5","10","2","13","13","5","6","8"],"multiplier":"0","bet":"8","point":"0"}},"currency":"TestCoin","start":1670772098,"finish":1670772098,"total_bet":"8","total_point":"0"}
	`

	fgcast = `
	{"id":"1234567890","game_code":"SG001","brand":"brand_test","username":"user_test","status":289,"position":289,"stages":14,"result":{"0":{"id":"0","brand":"brand_test","username":"user_test","case":1,"stages":0,"pickem":["8","0"],"symbols":["3","2","5","9","3","2","0","6","0","9","2","1","4","10","2","6","7","0","8","4","13","6","11","5","9"],"multiplier":"7","bet":"8","point":"56"},"1":{"id":"1","brand":"brand_test","username":"user_test","case":32768,"stages":1,"pickem":[""],"symbols":["3","3","3","3","2","7","11","11","11","8","2","11","11","11","7","1","11","11","11","4","4","6","6","6","4"],"multiplier":"0","bet":"0","point":"0"},"10":{"id":"10","brand":"brand_test","username":"user_test","case":32768,"stages":10,"pickem":[""],"symbols":["2","3","5","4","1","2","3","0","4","3","2","2","2","3","5","4","2","4","5","5","4","2","5","4","2"],"multiplier":"5","bet":"0","point":"40"},"11":{"id":"11","brand":"brand_test","username":"user_test","case":32768,"stages":11,"pickem":[""],"symbols":["3","5","3","4","3","1","4","3","4","3","1","2","3","4","1","1","3","5","3","1","1","1","1","3","1"],"multiplier":"0","bet":"0","point":"0"},"12":{"id":"12","brand":"brand_test","username":"user_test","case":32768,"stages":12,"pickem":[""],"symbols":["2","2","4","3","3","4","4","5","4","5","4","1","3","4","5","1","1","3","3","2","1","1","3","5","2"],"multiplier":"0","bet":"0","point":"0"},"13":{"id":"13","brand":"brand_test","username":"user_test","case":32768,"stages":13,"pickem":[""],"symbols":["2","4","4","4","5","2","4","3","4","3","4","5","5","0","4","4","3","5","3","1","1","1","3","5","1"],"multiplier":"0","bet":"0","point":"0"},"14":{"id":"14","brand":"brand_test","username":"user_test","case":32768,"stages":14,"pickem":[""],"symbols":["2","3","2","4","1","2","4","4","2","1","2","1","5","2","5","1","5","3","5","2","1","4","3","5","2"],"multiplier":"0","bet":"0","point":"0"},"2":{"id":"2","brand":"brand_test","username":"user_test","case":32768,"stages":2,"pickem":[""],"symbols":["2","4","4","4","9","3","9","9","9","6","11","9","9","9","1","5","9","9","9","10","5","4","4","4","8"],"multiplier":"0","bet":"0","point":"0"},"3":{"id":"3","brand":"brand_test","username":"user_test","case":32768,"stages":3,"pickem":[""],"symbols":["1","5","5","5","9","1","2","2","2","5","2","2","2","2","1","3","2","2","2","9","3","4","4","4","11"],"multiplier":"33.75","bet":"0","point":"270"},"4":{"id":"4","brand":"brand_test","username":"user_test","case":32768,"stages":4,"pickem":[""],"symbols":["2","4","4","4","8","9","6","6","6","3","3","6","6","6","11","5","6","6","6","6","4","3","3","3","2"],"multiplier":"0","bet":"0","point":"0"},"5":{"id":"5","brand":"brand_test","username":"user_test","case":32768,"stages":5,"pickem":[""],"symbols":["3","5","5","5","7","3","4","4","4","4","3","4","4","4","4","7","4","4","4","4","1","3","3","3","6"],"multiplier":"0","bet":"0","point":"0"},"6":{"id":"6","brand":"brand_test","username":"user_test","case":32768,"stages":6,"pickem":[""],"symbols":["1","5","5","5","4","3","4","4","4","7","4","4","4","4","9","1","4","4","4","3","3","3","3","3","11"],"multiplier":"33.75","bet":"0","point":"270"},"7":{"id":"7","brand":"brand_test","username":"user_test","case":32768,"stages":7,"pickem":[""],"symbols":["3","3","3","3","1","5","4","5","1","3","4","1","5","5","5","3","5","3","5","5","1","4","3","3","1"],"multiplier":"10","bet":"0","point":"80"},"8":{"id":"8","brand":"brand_test","username":"user_test","case":32768,"stages":8,"pickem":[""],"symbols":["2","2","4","3","4","2","4","4","1","1","2","1","4","5","1","1","1","4","5","5","1","1","3","3","2"],"multiplier":"0","bet":"0","point":"0"},"9":{"id":"9","brand":"brand_test","username":"user_test","case":32768,"stages":9,"pickem":[""],"symbols":["2","4","4","3","2","2","2","4","5","3","2","3","5","4","1","1","3","3","4","1","1","2","3","4","1"],"multiplier":"0","bet":"0","point":"0"}},"currency":"TestCoin","start":1671066108,"finish":1671066108,"total_bet":"8","total_point":"716"}
	`

	bgchoosecast = `
	{"id":"1234567890","game_code":"SG001","brand":"brand_test","username":"user_test","status":68,"position":16,"stages":0,"result":{"0":{"id":"0","brand":"brand_test","username":"user_test","case":4,"stages":0,"pickem":["8"],"symbols":["2","2","13","3","5","2","2","13","7","1","6","5","1","13","1","13","3","1","13","1","13","5","5","11","5"],"multiplier":"0","bet":"8","point":"0"}},"currency":"TestCoin","start":1671070658,"finish":1671070658,"total_bet":"8","total_point":"0"}
	`
)

func TestRoundNotExist(t *testing.T) {
	// {WORKSPACE.<workspace>:GAME.<game_code>|ROUND}	= WORKSPACE.test:GAME.test|ROUND
	// key{<round_id>:<user_id>}: value{<round>}		= test:test
	// reset()
	// mockHset(3, "WORKSPACE.test:GAME.test|ROUND", "ROUNDID.test:USERID.test", fgcast)
	a_round := &games.Rounds{}

	err := a_round.Unmarshal([]byte(bgchoosecast))
	if err != nil {
		t.Errorf("%+v", err.Error()+"...")
	}

	res, err := testScrtiptor.RoundUpdate("test", "test", "test", "test", *a_round)
	if err != nil {
		t.Errorf("%+v", err.Error()+"...")
	}

	if res.Status == "not_exist" {
		t.Error("RoundNotExist is not not_exist.")
	}

	t.Logf("TestRoundNotExist is %+v", res)
}

func TestScript(t *testing.T) {
	// {WORKSPACE.<workspace>:GAME.<game_code>|ROUND}	= WORKSPACE.test:GAME.test|ROUND
	// key{<round_id>:<user_id>}: value{<round>}		= test:test
	// reset()
	// mockHset(3, "WORKSPACE.test:GAME.test|ROUND", "ROUNDID.test:USERID.test", fgcast)
	// reset()
	// mockHset(3, "WORKSPACE.test:GAME.test|ROUND", "ROUNDID.test:USERID.test", bgchoosecast)

	res, err := testScrtiptor.Cache.Exec(ts_r_udp, []string{"test", "test"}, "test", "test", fgcast)
	if err != nil {
		t.Errorf("%+v", err.Error()+"...")
	}

	reader, ok := res.(string)
	if !ok {
		t.Errorf("invalid response")
	}

	b_round := []byte(reader)

	t.Logf("string(b_round) is %+v", string(b_round))

	res, err = testScrtiptor.Cache.Exec(ts_r, []string{"test", "test"}, "test", "test")
	if err != nil {
		t.Errorf("%+v", err.Error()+"...")
	}

	reader, ok = res.(string)
	if !ok {
		t.Errorf("invalid response")
	}

	b_round = []byte(reader)

	round := &scriptor.RoundRecords{}
	err = round.Unmarshal(b_round)
	if err != nil {
		t.Errorf("%+v", err.Error()+"...")
	}

	if round.Status == "NOT_EXIST" {
		t.Error("RoundNotExist is not not_exist.")
	}

	t.Logf("TestRoundNotExist is %+v", round)
	t.Logf("TestRoundNotExist is %+v", round.Round.Result[strconv.Itoa(int(round.Round.Stages))])
}

var ts_r = `
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
`

var ts_r_udp = `
if #KEYS < 2 then
    return redis.error_reply('ILLEGAL_KEYS')
end

-- Get Keys
local WORKSPACE = tostring(KEYS[1])
local GAMECODE = tostring(KEYS[2])

-- Get Args
local ROUNDID = tostring(ARGV[1])
local USERID = tostring(ARGV[2])
local ROUND = tostring(ARGV[3])

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
if round_exist == 1 then
     -- Get the round from the hash table
    local ROUND_VALUE = cjson.decode(redis.call('HGET', round_key, round_vkey))
    if bit.band(ROUND_VALUE['position'], 16) == 16 then
        result.status = 'OK'
        result.round = cjson.decode(ROUND)
        result.round['stages'] = 0
        redis.call('HSET', round_key, round_vkey, cjson.encode(result.round))
        return cjson.encode(result)
    end

    result.status = 'EXIST'
    return cjson.encode(result)
end

redis.call('HSET', round_key, round_vkey, ROUND)

-- Update the round
result.status = 'OK'
result.round = ROUND
return cjson.encode(result)
`

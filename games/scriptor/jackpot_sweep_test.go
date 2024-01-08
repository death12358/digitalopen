package scriptor_test

import (
	"testing"

	"github.com/death12358/digitalopen/games/scriptor"
)

// TestJackpotSweepScript -
//
//	EVALSHA <script_sha1> 4 <workspace> <brand> <currency> <pool_code> <default_point>
func TestSJackpotSweepScript(t *testing.T) {
	res, err := testScrtiptor.Cache.Exec(ts_j_sweep, []string{"test", "test", "test", "test"}, "1")
	if err != nil {
		t.Errorf("%+v", err.Error()+"...")
	}

	reader, ok := res.(string)
	if !ok {
		t.Errorf("invalid response")
	}
	t.Logf("reader is %+v", reader)

	jppool := &scriptor.JackpotPool{}
	err = jppool.Unmarshal([]byte(reader))
	if err != nil {
		t.Errorf("invalid response")
	}

	t.Logf("string(jppool) is %+v", jppool)
}

var (
	ts_j_sweep = `
--- Choose database
redis.call('SELECT', 2)

-- Get Keys
if #KEYS < 4 then
    return redis.error_reply('ILLEGAL_KEYS')
end
local WORKSPACE = tostring(KEYS[1])
local BRAND = tostring(KEYS[2])
local CURRENCY = tostring(KEYS[3])
local POOLCODE = tostring(KEYS[4])

-- Get Args
if #ARGV < 1 then
    return redis.error_reply('ILLEGAL_ARGV')
end
local DEFAULT_POINT = tonumber(ARGV[1])

-- Generate the jackpot key
local jackpot_key = string.format('WORKSPACE.%s:BRAND.%s:CURRENCY.%s|JACKPOT', WORKSPACE, BRAND, CURRENCY)
if redis.call('EXISTS', jackpot_key) == 0 then
    result = {
        status = 'OK',
        jackpot_pool = 0,
    }
    return result
end

-- Generate the jackpot vkey
local jackpot_vkey = string.format('POOLCODE.%s', POOLCODE)

-- HGET
local jackpot_point = redis.call('HGET', jackpot_key, jackpot_vkey)
local result = {
    status = 'OK',
    jackpot_pool = jackpot_point,
}

-- HSET
redis.call('HSET', jackpot_key, jackpot_vkey, DEFAULT_POINT)

return cjson.encode(result)
	`
)

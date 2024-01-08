package scriptor_test

import (
	"testing"

	"github.com/death12358/digitalopen/games/scriptor"
)

// TestJackpotPushScript -
//
//	EVALSHA  <script_sha1> 4 <workspace> <brand> <currency> <pool_code> <point>
func TestPJackpotPushScript(t *testing.T) {
	res, err := testScrtiptor.Cache.Exec(ts_j_push, []string{"test", "test", "c", "test"}, "20.00000000000001")
	if err != nil {
		t.Errorf("%+v", err.Error()+"...")
	}

	reader, ok := res.(string)
	if !ok {
		t.Errorf("invalid response")
	}

	jppool := &scriptor.JackpotPool{}
	err = jppool.Unmarshal([]byte(reader))
	if err != nil {
		t.Errorf("invalid response")
	}

	t.Logf("string(jppool) is %+v", jppool)
}

var (
	ts_j_push = `
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

-- Get Args
if #ARGV < 1 then
	return redis.error_reply('ILLEGAL_ARGV')
end
local POINT = tostring(ARGV[1])

-- Generate the jackpot key
local jackpot_key = string.format('WORKSPACE.%s:BRAND.%s:CURRENCY.%s|JACKPOT', WORKSPACE, BRAND, CURRENCY)
-- Generate the jackpot vkey
local jackpot_vkey = string.format('POOLCODE.%s', POOLCODE)

-- HINCRBYFLOAT
local jackpot_point = redis.call('HINCRBYFLOAT', jackpot_key, jackpot_vkey, POINT)

local result = {
	status = 'OK',
	jackpot_pool = tostring(jackpot_point),
}

return cjson.encode(result)
	`
)

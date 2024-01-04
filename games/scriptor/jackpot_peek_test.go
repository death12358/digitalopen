package scriptor_test

import (
	"testing"

	"github.com/death12358/digitalopn/games/scriptor"
)

// TestJackpotPeekScript -
//
//	EVALSHA  <script_sha1> 4 <workspace> <brand> <currency> <pool_code>
func TestPJackpotPeekScript(t *testing.T) {
	res, err := testScrtiptor.Cache.Exec(ts_j_peek, []string{"test", "test", "test", "test"})
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
	ts_j_peek = `
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
	`
)

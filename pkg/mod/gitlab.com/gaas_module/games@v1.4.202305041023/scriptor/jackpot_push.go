package scriptor

import (
	"errors"

	"github.com/shopspring/decimal"
)

// JackpotPush -
//   EVALSHA <script_sha1> 4 <workspace> <brand> <currency> <pool_code> <point>
func (s *Scriptor) JackpotPush(workspace, brand, currency, pool_code string, point decimal.Decimal) (*JackpotPool, error) {
	res, err := s.Cache.ExecSha(_jackpot_push, []string{workspace, brand, currency, pool_code}, point.StringFixed(17))
	if err != nil {
		return nil, err
	}

	reader, ok := res.(string)
	if !ok {
		return nil, errors.New("invalid response")
	}

	b_pool := []byte(reader)

	pool := &JackpotPool{}
	err = pool.Unmarshal(b_pool)

	return pool, err
}

var (
	_jackpot_push          = "jackpot.push"
	_jackpot_push_Template = `
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

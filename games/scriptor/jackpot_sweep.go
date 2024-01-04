package scriptor

import (
	"errors"

	"github.com/shopspring/decimal"
)

func (s *Scriptor) JackpotSweep(workspace, brand, currency, pool_code string, default_point decimal.Decimal) (*JackpotPool, error) {
	res, err := s.Cache.ExecSha(_jackpot_push, []string{workspace, brand, currency, pool_code}, default_point.StringFixed(17))
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
	_jackpot_sweep          = "jackpot.sweep"
	_jackpot_sweep_Template = `
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
		local result = {
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

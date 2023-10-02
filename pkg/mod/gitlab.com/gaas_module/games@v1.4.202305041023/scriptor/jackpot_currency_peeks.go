package scriptor

import "errors"

// JackpotPush -
//  EVALSHA <script_sha1> 2 <workspace> <brand>
func (s *Scriptor) JackpotCurrencyPeeks(workspace, brand string) (*JackpotCurrencyPools, error) {
	res, err := s.Cache.ExecSha(_jackpot_currency_peeks, []string{workspace, brand})
	if err != nil {
		return nil, err
	}

	reader, ok := res.(string)
	if !ok {
		return nil, errors.New("invalid response")
	}

	b_pool := []byte(reader)

	pool := &JackpotCurrencyPools{}
	err = pool.Unmarshal(b_pool)

	return pool, err
}

var (
	_jackpot_currency_peeks          = "jackpot.currency.peeks"
	_jackpot_currency_peeks_Template = `
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
	`
)

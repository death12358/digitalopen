package scriptor

import "errors"

// JackpotPush -
//   EVALSHA  <script_sha1> 3 <workspace> <brand> <currency>
func (s *Scriptor) JackpotPeeks(workspace, brand, currency string) (*JackpotPools, error) {
	res, err := s.Cache.ExecSha(_jackpot_peeks, []string{workspace, brand, currency})
	if err != nil {
		return nil, err
	}

	reader, ok := res.(string)
	if !ok {
		return nil, errors.New("invalid response")
	}

	b_pool := []byte(reader)

	pool := &JackpotPools{}
	err = pool.Unmarshal(b_pool)

	return pool, err
}

var (
	_jackpot_peeks          = "jackpot.peeks"
	_jackpot_peeks_Template = `
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
	`
)

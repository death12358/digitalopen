package scriptor

import (
	"errors"

	"digitalopen/games"
)

// RoundUpdate - 更新回合單
func (s *Scriptor) RoundUpdate(workspace, gameCode, roundID, userID string, record games.Rounds) (*RoundRecords, error) {
	b_round, err := record.Marshal()
	if err != nil {
		return nil, err
	}

	res, err := s.Cache.ExecSha(_round_update, []string{workspace, gameCode}, roundID, userID, b_round)
	if err != nil {
		return nil, err
	}

	reader, ok := res.(string)
	if !ok {
		return nil, errors.New("invalid response")
	}

	b_round = []byte(reader)

	result := &RoundRecords{}
	err = result.Unmarshal(b_round)

	return result, err
}

var (
	_round_update          = "round.update"
	_round_update_Template = `
if #KEYS < 2 then
    return redis.error_reply('ILLEGAL_KEYS')
end

-- Get Keys
local WORKSPACE = tostring(KEYS[1])
local GAMECODE = tostring(KEYS[2])

-- Get Args
local ROUNDID = tostring(ARGV[1])
local USERID = tostring(ARGV[2])
local ROUND = cjson.decode(ARGV[3])

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
     -- Get the round from the hash table
    if bit.band(ROUND_VALUE['position'], 16) == 16 then
        result.status = 'OK'
        result.round = ROUND

        redis.call('HSET', round_key, round_vkey, cjson.encode(ROUND))
        return cjson.encode(result)
    end

    result.status = 'EXIST'
    return cjson.encode(result)
end

redis.call('HSET', round_key, round_vkey, cjson.encode(ROUND))

-- Update the round
result.status = 'OK'
result.round = ROUND
return cjson.encode(result)
`
)

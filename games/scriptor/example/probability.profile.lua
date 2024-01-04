-- probability.profile.<OPERATION>.lua
-- database: 0
-- profile data model(HASH) example:
--	KEY: 	{WORKSPACE.<workspace>|PROFILE} 
--	VALUE: 	[
--	key{<...>}: value{<rtp_vaule>},
--]
-- INPUT: 	
--  EVALSHA  <script_sha1> 1 <workspace> <brand> <ip> <game_code> <user_id>
-- OUTPUT:
--  <rtp_value>(string)
if #KEYS < 1 then
    return redis.error_reply('ILLEGAL_KEYS')
end
-- Get Keys
local WORKSPACE = tostring(KEYS[1])

if #ARGV < 4 then
    return redis.error_reply('ILLEGAL_ARGV')
end
-- Get Args
local BRAND = tostring(ARGV[1])
local IP = tostring(ARGV[2])
local GAMECODE = tostring(ARGV[3])
local USERID = tostring(ARGV[4])

-- default rtp value
local RTP_VALUE = '98'

-- Choose database
redis.call('SELECT', 0)

-- Generate the profile key
local profile_key = string.format('WORKSPACE.%s|PROFILE', WORKSPACE)
local profile_key_exist = redis.call('EXISTS', profile_key)
if profile_key_exist == 0 then
	return RTP_VALUE
end

-- Generate the default profile vkey
local profile_vkey = string.format('DEFAULT|RTP')
local profile_vkey_exist = redis.call('HEXISTS', profile_key, profile_vkey)
if profile_vkey_exist == 1 then
	RTP_VALUE = redis.call('HGET', profile_key, profile_vkey)
	return RTP_VALUE
end

-- Generate the brand profile vkey
profile_vkey = string.format('BRAND.%s:DEFAULT|RTP', BRAND)
profile_vkey_exist = redis.call('HEXISTS', profile_key, profile_vkey)
if profile_vkey_exist == 1 then
	RTP_VALUE = redis.call('HGET', profile_key, profile_vkey)
	return RTP_VALUE
end

-- Generate the ip profile vkey
profile_vkey = string.format('IP.%s|RTP', IP)
profile_vkey_exist = redis.call('HEXISTS', profile_key, profile_vkey)
if profile_vkey_exist == 1 then
	RTP_VALUE = redis.call('HGET', profile_key, profile_vkey)
	return RTP_VALUE
end

-- Generate the user profile vkey
profile_vkey = string.format('BRAND.%s:USERID.%s|RTP', BRAND, USERID)
profile_vkey_exist = redis.call('HEXISTS', profile_key, profile_vkey)
if profile_vkey_exist == 1 then
	RTP_VALUE = redis.call('HGET', profile_key, profile_vkey)
	return RTP_VALUE
end

-- Generate the game profile vkey
profile_vkey = string.format('BRAND.%s:GAMECODE.%s:DEFAULT|RTP', BRAND, GAMECODE)
profile_vkey_exist = redis.call('HEXISTS', profile_key, profile_vkey)
if profile_vkey_exist == 1 then
	RTP_VALUE = redis.call('HGET', profile_key, profile_vkey)
	return RTP_VALUE
end

return RTP_VALUE



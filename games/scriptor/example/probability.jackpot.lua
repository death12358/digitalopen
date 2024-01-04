-- probability.jackpot.<OPERATION>.lua
-- database: 2
-- round data model(HASH) example:
--	KEY: 	{WORKSPACE.<workspace>:POOLNAME.<pool_code>|JACKPOT} 
--	VALUE: 	[
--	key{<jackpotname>}: value{<jackpot_point>},
--]
-- OPERATION: 
--  PUSH:           push jackpot. (帶入jackpotname, default_point, point)
--  PULL:           pull jackpot. (帶入jackpotname, default_point, pull_point)
--  SWEEP:          sweep jackpot. (帶入jackpotname, default_point, pull_point)
--  PEEK:           peek jackpot.
-- INPUT: 	
--  EVALSHA  <script_sha1> 2 <workspace> <pool_code>
-- OUTPUT:
--  <rtp_value>(string)

-- 引入Redis Lua庫
local redis = require "resty.redis"

-- 建立Redis連接
local red = redis:new()
local ok, err = red:connect("127.0.0.1", 6379)
if not ok then
    ngx.say("無法連接到Redis: ", err)
    return
end

-- 從Redis中獲取資料
local data, err = red:get("your_key")
if not data then
    ngx.say("無法獲取資料: ", err)
    return
end

-- 關閉Redis連接
local ok, err = red:close()
if not ok then
    ngx.say("無法關閉Redis連接: ", err)
    return
end

-- 輸出獲取到的資料
ngx.say("從Redis獲取到的資料: ", data)
local val = redis.call('GET', KEYS[1])

if val == false then
    if redis.call('SET', KEYS[1], ARGV[1], 'EX', ARGV[2])['ok'] == "OK" then
        return 1
    else
        return 0
    end
elseif val == ARGV[1] then
    return redis.call('EXPIRE', KEYS[1], ARGV[2])
else
    return 0
end
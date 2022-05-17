package storage

import "secretservice/storage/rediskeeper"

var Keep = rediskeeper.GetRedisKeeper()

//var Keep = keeper.NewKeeper()

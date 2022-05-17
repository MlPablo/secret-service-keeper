package storage

import "secretservice/storage/keeper"

//var Keep = rediskeeper.GetRedisKeeper()

var Keep = keeper.NewKeeper()

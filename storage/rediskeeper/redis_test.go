package rediskeeper

//func TestKeeperSet(t *testing.T) {
//	keeper := GetRedisKeeper()
//	key := "foo"
//	text := "bar"
//	keeper.Set(key, text)
//	v, _ := keeper.Get(key)
//	fmt.Println(v)
//	if v != text {
//		t.Error("Should be founded")
//	}
//}

//	err := keeper.Set(key, text)
//	if err == nil {
//		t.Error("Should be errored")
//	}
//}

//
//func TestKeeperGet(t *testing.T) {
//	keeper := DummyKepper{make(map[string]string)}
//	key := "foo"
//	text := "bar"
//	keeper.Set(key, text)
//	v, _ := keeper.Get(key)
//	if v != text {
//		t.Error("Should be returned")
//	}
//	if len(keeper.mem) != 0 {
//		t.Error("Should be deleted")
//	}
//	_, err := keeper.Get(key)
//	if err == nil {
//		t.Error("Should be errored")
//	}
//
//}
//
//func TestKeeperDelete(t *testing.T) {
//	keeper := DummyKepper{make(map[string]string)}
//	key := "foo"
//	text := "bar"
//	keeper.Set(key, text)
//	keeper.Delete(key)
//	if len(keeper.mem) != 0 {
//		t.Error("Should be deleted")
//	}
//}

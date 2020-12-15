package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Set(key string, value interface{}) error {
	dataS, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var dataJson = string(dataS)
	err = GetRdS().GetClient().Set(ctx, key, dataJson, 0).Err()
	return err
}

func ExampleClient() {
	rdb := new(RDSClient).GetClient()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

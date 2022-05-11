package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	fmt.Println("Testing Golang Redis")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()

	fmt.Println(pong, err)

	err = rdb.Set(ctx, "key", "value", 0).Err()

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

}

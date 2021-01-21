package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var (
	fromAddr    = flag.String("from-addr", "localhost:6379", "address of source redis")
	fromDBIndex = flag.Int("from-db", 0, "db index of source redis")
	fromKey     = flag.String("from-key", "from-key", "key in source redis")

	toAddr    = flag.String("to-addr", "localhost:6379", "address of target redis")
	toDBIndex = flag.Int("to-db", 0, "db index of target redis")
	toKey     = flag.String("to-key", "to-key", "key in target redis")
)

var ctx = context.Background()

func main() {
	flag.Parse()
	flag.PrintDefaults()
	// source rdb
	// target rdb
	src := getRedis(*fromAddr, *fromDBIndex)
	target := getRedis(*toAddr, *toDBIndex)

	// count zset from source rdb
	count, err := src.ZCount(ctx, *fromKey, "-inf", "+inf").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key count", count)

	items, err := src.ZRangeWithScores(ctx, *fromKey, 0, count).Result()
	if err != nil {
		panic(err)
	}

	for _, z := range items {
		target.ZAdd(ctx, *toKey, &z)
		fmt.Println("copy", z)
	}

	// fetch all from source rdb
	// add all to target rdb
}

func getRedis(addr string, db int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       db, // use default DB
	})
	fmt.Println("connect: ", addr, db)
	return rdb
}

package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// redis
// cache
// simple queue
// ranking

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err = redisdb.Ping().Result()
	return
}

// set/get

func set() (_, err error) {
	err = redisdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return nil, err
	}

	val, err := redisdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score fialed, err:%v\n", err)
		return nil, err
	}
	fmt.Println("score:", val)
	return
}

func rank() (_, err error) {
	err = initRedis()
	if err != nil {
		fmt.Printf("connect to redis failed, err:%v\n", err)
		return nil, err
	}
	fmt.Println("connect to redis succeed")

	key := "rank"
	items := []redis.Z{
		{Score: 90, Member: "PHP"},
		{Score: 96, Member: "Golang"},
		{Score: 97, Member: "Python"},
		{Score: 90, Member: "Java"},
		{Score: 90, Member: "C++"},
	}

	num, err := redisdb.ZAdd(key, items...).Result()
	if err != nil {
		fmt.Printf("zdd failed, err:%v\n", err)
		return nil, err
	}
	fmt.Printf("zadd %d succeed. \n", num)

	// give 'golang' 10 points
	newScore, err := redisdb.ZIncrBy(key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("ZIncrBy failed, err:%v\n", err)
		return nil, err
	}
	fmt.Println("newscore:", newScore)

	// get top 3

	return
}

func main() {
	initRedis()
	rank()
}

package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// 初始化模拟 Redis 客户端
func initRedis() *redis.Client {
	var client *redis.Client
	redisopt := &redis.Options{
		Addr: `127.0.0.1:6379`,
		//Password:    Passwd, // no password set
		//DB:          DB,     // use default DB
		PoolSize:    300, //reids连接池
		DialTimeout: 10 * time.Second,
		ReadTimeout: 10 * time.Second,
	}

	client = redis.NewClient(redisopt)
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return nil
	}
	return client
}

func pipeline(client *redis.Client) (err error) {
	createTime := time.Now().UnixNano()
	cmders, err := client.Pipelined(func(pipeliner redis.Pipeliner) error {
		pipeliner.HSet("pipeline_hash", "1", createTime)
		pipeliner.HSet("pipeline_hash", "2", createTime+1)
		pipeliner.HSet("pipeline_hash", "3", createTime+2)
		pipeliner.Set("pipeline_str", "1111", 20*time.Second)

		return nil
	})
	if err != nil {
		return err
	}
	//fmt.Printf("cmders len:%d\n", len(cmders))
	for _, cmd := range cmders {
		if cmd.Err() != nil {
			return cmd.Err()
		}
	}
	return nil
}

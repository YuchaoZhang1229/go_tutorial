package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

func connectRedis() *redis.Client {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Replace with your Redis server address
		Password: "",               // No password for local development
		DB:       0,                // Default DB
		Protocol: 2,                // Connection protocol
	})

	// Ping the Redis server to check the connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)

	return rdb
}

// Expire Keys:
func expireKeys(rdb *redis.Client) {
	// Set a key with expiration time
	err := rdb.Set(ctx, "temporary", "I will expire soon!", 10*time.Second).Err()
	if err != nil {
		log.Fatal(err)
	}
}

// Subscribe to Pub/Sub Channels
func subs(rdb *redis.Client) {
	// subscribe to a channel
	pubsub := rdb.Subscribe(ctx, "mychannel")
	defer pubsub.Close()
	channel := pubsub.Channel()

	// Listen for messages
	for msg := range channel {
		fmt.Printf("subs 收到消息: 频道=%s, 内容=%s\n", msg.Channel, msg.Payload)
	}
}

func pubs(rdb *redis.Client) {
	// 给订阅者一点启动时间
	time.Sleep(100 * time.Millisecond)
	// publish a message
	for i := 1; i <= 3; i++ {
		message := fmt.Sprintf("你好，这是第%d条消息", i)
		err := rdb.Publish(ctx, "mychannel", message).Err()
		if err != nil {
			panic(err)
		}
		fmt.Printf("pubs 已发布: %s\n", message)
	}
}

func main() {
	rdb := connectRedis()
	expireKeys(rdb)
	go subs(rdb) // 直接调用的话, 程序流程卡在 subs 函数，永远无法到达 pubs 函数
	pubs(rdb)

}

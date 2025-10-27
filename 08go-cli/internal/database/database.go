package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-tutorial/08go-cli/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// New 创建并返回一个 Redis 客户端
func New(cfg config.DatabaseConfig) (*redis.Client, error) {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,     // Redis 服务器地址
		Password: cfg.Password, // 密码
		DB:       cfg.DB,       // 数据库索引
		Protocol: cfg.Protocol, // 连接协议
	})

	// 测试连接
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Printf("Connected to Redis at %s", cfg.Addr)
	fmt.Println("Connected to Redis:", cfg.Addr)

	return client, nil
}

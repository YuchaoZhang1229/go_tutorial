package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ===== context.WithTimeout 示例 =====
// 作用：创建一个会在 2 秒后自动取消的上下文
// 执行流程：
// 1. 启动一个 goroutine 执行工作（模拟 1 秒完成的任务）
// 2. 由于工作只需 1 秒，会在超时前完成，输出 "Work completed within timeout"
// 3. 如果工作超过 2 秒，会触发 ctx.Done()，输出取消信息
func demoContextWithTimeout() {
	fmt.Println("\n--- context.WithTimeout 示例 ---")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Work completed within timeout")
		case <-ctx.Done():
			fmt.Printf("Work cancelled: %v\n", ctx.Err())
		}
	}()

	time.Sleep(3 * time.Second)
}

// ===== context.WithCancel 示例 =====
// 作用：创建可手动取消的上下文
// 执行流程：
// 1. goroutine 不断循环工作，每 500ms 输出 "Working..."
// 2. 2 秒后主函数调用 cancel() 发送取消信号
// 3. goroutine 收到信号后退出循环
func demoContextWithCancel() {
	fmt.Println("\n--- context.WithCancel 示例 ---")

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Received cancel signal: %v\n", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// 让goroutine工作一会儿
	time.Sleep(2 * time.Second)

	// 取消操作
	fmt.Println("Cancelling the context...")
	cancel()

	wg.Wait()
}

// ===== context.WithValue 示例 =====
// 作用：在上下文中存储和传递请求范围的数据
// 特点：
// 1. 使用自定义类型作为键（避免字符串冲突）
// 2. goroutine 可以从上下文中读取传递的值
func demoContextWithValue() {
	fmt.Println("\n--- context.WithValue 示例 ---")

	type keyType string

	var (
		userKey      keyType = "user"
		requestIDKey keyType = "requestID"
	)

	ctx := context.WithValue(context.Background(), userKey, "john_doe")
	ctx = context.WithValue(ctx, requestIDKey, "12345")

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()

		if user, ok := ctx.Value(userKey).(string); ok {
			fmt.Printf("User: %s\n", user)
		}

		if requestID, ok := ctx.Value(requestIDKey).(string); ok {
			fmt.Printf("Request ID: %s\n", requestID)
		}
	}(ctx)

	wg.Wait()
}

func main() {
	fmt.Println("=== Go 并发编程 Demo ===")

	// 2. context 包示例
	demoContextWithTimeout()
	demoContextWithCancel()
	demoContextWithValue()
}

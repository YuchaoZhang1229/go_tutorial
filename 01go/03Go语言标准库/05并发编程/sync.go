package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 共享资源
var (
	counter     int64
	sharedValue int
	mutex       sync.Mutex
	rwMutex     sync.RWMutex
)

// ===== sync.Mutex 示例 =====
// 作用：保证同一时间只有一个 goroutine 能访问共享资源
// 特点：
// 1. 写操作和读操作都需要获取锁
// 2. 保护 sharedValue 的并发安全
// 3. 5个 goroutine 会依次执行，不会同时修改数据
//
//	func demoMutex() {
//		 mutex.Lock()
//		 defer mutex.Unlock()
//		 sharedValue++ // 临界区操作
//	}
func demoMutex() {
	fmt.Println("\n--- sync.Mutex 示例 ---")

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			mutex.Lock()
			defer mutex.Unlock()

			// 临界区操作
			sharedValue++
			fmt.Printf("Goroutine %d: sharedValue = %d\n", id, sharedValue)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}
	wg.Wait()
}

// ===== sync.RWMutex 示例 =====
// 作用：允许多个读操作同时进行，但写操作是排他的
// 特点：
// 读锁：多个 goroutine 可以同时获取读锁
// 写锁：写锁是排他的，有写锁时不能有读锁或其他写锁
// 适合读多写少的场景
//
//	func demoRWMutex() {
//		rwMutex.Lock()    // 写锁
//		rwMutex.RLock()   // 读锁
//	}
func demoRWMutex() {
	fmt.Println("\n--- sync.RWMutex 示例 ---")

	var wg sync.WaitGroup

	// 写操作
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			rwMutex.Lock()
			defer rwMutex.Unlock()

			sharedValue = id
			fmt.Printf("Writer %d: set sharedValue = %d\n", id, sharedValue)
			time.Sleep(200 * time.Millisecond)
		}(i)
	}

	// 读操作
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			rwMutex.RLock()
			defer rwMutex.RUnlock()

			fmt.Printf("Reader %d: read sharedValue = %d\n", id, sharedValue)
			time.Sleep(100 * time.Millisecond)
		}(i)
	}

	wg.Wait()
}

// ===== sync.WaitGroup 示例 =====
// 作用：等待一组 goroutine 完成执行
// 使用模式：
// 1. 主 goroutine 调用 Add 设置要等待的数量
// 2. 每个工作 goroutine 结束时调用 Done
// 3. 主 goroutine 调用 Wait 阻塞直到所有工作完成
//
//	func demoWaitGroup() {
//		var wg sync.WaitGroup
//		wg.Add(1)     // 增加计数
//		wg.Done()     // 减少计数
//		wg.Wait()     // 等待计数归零
//	}
func demoWaitGroup() {
	fmt.Println("\n--- sync.WaitGroup 示例 ---")

	var wg sync.WaitGroup
	tasks := 5

	for i := 0; i < tasks; i++ {
		wg.Add(1)
		go func(taskID int) {
			defer wg.Done()

			fmt.Printf("Task %d started\n", taskID)
			time.Sleep(time.Duration(taskID) * 100 * time.Millisecond)
			fmt.Printf("Task %d completed\n", taskID)
		}(i)
	}

	fmt.Println("Waiting for all tasks to complete...")
	wg.Wait()
	fmt.Println("All tasks completed!")
}

// ===== sync.Once 示例 =====
// 作用：确保某个函数只被执行一次，即使在并发环境下
// 应用场景：
// 1. 单例模式初始化
// 2. 配置加载
// 3. 全局资源初始化
//
//	func demoOnce() {
//		var once sync.Once
//		once.Do(initialize) // 只会执行一次
//	}
func demoOnce() {
	fmt.Println("\n--- sync.Once 示例 ---")

	var once sync.Once
	var wg sync.WaitGroup

	initialize := func() {
		fmt.Println("Initialization performed! (This should only appear once)")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			fmt.Printf("Goroutine %d calling Once\n", id)
			once.Do(initialize)
		}(i)
	}

	wg.Wait()
}

// ===== sync.Pool 示例 =====
// 作用：缓存和复用临时对象，减少内存分配
// 特点：
// 1. 适合创建成本高的对象
// 2. 对象会被垃圾回收，不保证长期存在
// 3. 自动处理对象的创建和复用
func demoPool() {
	fmt.Println("\n--- sync.Pool 示例 ---")

	type expensiveObject struct {
		ID   int
		Data string
	}

	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new expensive object")
			return &expensiveObject{ID: -1, Data: "default"}
		},
	}
	obj := pool.Get().(*expensiveObject) // 获取对象
	fmt.Printf("Using object: %+v\n", obj)
	pool.Put(obj) // 放回对象
}

func main() {
	fmt.Println("=== Go 并发编程 Demo ===")

	// 1. sync 包示例
	demoMutex()
	demoRWMutex()
	demoWaitGroup()
	demoOnce()
	demoPool()

	// 2. atomic 包示例
	//demoAtomic()
	//demoCAS()

}

// ===== atomic 原子操作示例 =====
func demoAtomic() {
	fmt.Println("\n--- atomic 原子操作示例 ---")

	var wg sync.WaitGroup

	// 多个goroutine同时增加计数器
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))
}

// ===== CAS (Compare-And-Swap) 操作示例 =====
func demoCAS() {
	fmt.Println("\n--- CAS (Compare-And-Swap) 操作示例 ---")

	var value int32 = 0
	var wg sync.WaitGroup

	// 尝试更新值的goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int, expected int32, new int32) {
			defer wg.Done()

			if atomic.CompareAndSwapInt32(&value, expected, new) {
				fmt.Printf("Goroutine %d: CAS成功, 值从 %d 更新为 %d\n", id, expected, new)
			} else {
				current := atomic.LoadInt32(&value)
				fmt.Printf("Goroutine %d: CAS失败, 期望值 %d, 实际值 %d\n", id, expected, current)
			}
		}(i, int32(i-1), int32(i))
	}

	wg.Wait()
	fmt.Printf("最终值: %d\n", value)
}

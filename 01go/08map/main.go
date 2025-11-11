package main

import (
	"fmt"
	"sync"
	"time"
)

type safeMap struct {
	sync.RWMutex                   // 读写锁
	mp           map[string]string // map
}

// 读
func (sm *safeMap) Get(key string) string {
	sm.RLock()
	defer sm.RUnlock()
	return sm.mp[key]
}

// 写
func (sm *safeMap) Set(key, value string) string {
	sm.Lock()
	defer sm.Unlock()
	sm.mp[key] = value
	return value
}

func NewMap() *safeMap {
	return &safeMap{
		RWMutex: sync.RWMutex{},
		mp:      make(map[string]string),
	}
}

func main() {
	sm := NewMap()
	var wg sync.WaitGroup

	// 启动3个goroutine同时写入不同的键值对
	wg.Add(1)
	go func() {
		defer wg.Done()
		sm.Set("name", "张三")
		fmt.Println("设置: name -> 张三")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sm.Set("age", "25")
		fmt.Println("设置: age -> 25")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		sm.Set("city", "北京")
		fmt.Println("设置: city -> 北京")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("暂停1s")

	// 启动2个goroutine同时读取
	wg.Add(1)
	go func() {
		defer wg.Done()
		name := sm.Get("name")
		fmt.Printf("读取: name -> %s\n", name)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		age := sm.Get("age")
		fmt.Printf("读取: age -> %s\n", age)
	}()

	wg.Wait()

	// 最终结果
	fmt.Println("\n最终存储的数据:")
	fmt.Printf("name: %s\n", sm.Get("name"))
	fmt.Printf("age: %s\n", sm.Get("age"))
	fmt.Printf("city: %s\n", sm.Get("city"))
}

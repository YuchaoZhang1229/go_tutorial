### 第五阶段：并发编程
**sync 包功能：**
1. **Mutex** - 基本的互斥锁
2. **RWMutex** - 读写锁，支持并发读
3. **WaitGroup** - 等待一组 goroutine 完成
4. **Once** - 确保操作只执行一次 
5. **Pool** - 对象池，复用昂贵对象

sync.Mutex 是互斥锁，用于保护临界区，确保同一时间只有一个 goroutine 可以访问共享资源。
sync.RWMutex 是读写锁，允许多个读操作同时进行，但写操作是互斥的（同时只能有一个写操作，且写时不能读）。

**context 包功能：**
1. **WithTimeout** - 超时控制
2. **WithCancel** - 取消传播
3. **WithValue** - 上下文值传递

**atomic 包功能：**
1. **原子读写** - 无锁的原子操作
2. **CAS操作** - 比较并交换
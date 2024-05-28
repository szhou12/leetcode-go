# [Goroutine](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Concurrency vs Parallelism](#concurrency-vs-parallelism)
* [Synchronization Primitives](#synchronization-primitives)
    * [Mutex](#mutex-read-write-mutex-读写锁)

## Concurrency vs Parallelism
|| Concurrency <br>并发 | Parallelism <br>并行 |
| :-: | - | - |
|Concept|One core/program handles multiple tasks simultaneously.|Multiple cores/programs handle multiple tasks simultaneously. Each core handles a task respectively.|
|Characteristic|ONLY 逻辑上同时发生 (i.e., 多个任务"看上去"同时被处理)|BOTH 逻辑上同时发生，物理上也同时发生(i.e., 多个任务"确实"同时被处理)|
|Analogy|一个咖啡店员同时做三件事：下单，做咖啡，擦桌子。他在三个事之间来回切换 (context switching)。|三个咖啡店员同时各自负责其中一件事：A下单，B做咖啡，C擦桌子。|
|Python|Threading (多线程) in Python is often used to achieve concurrency.<br>Use Threading for I/O-bound tasks.|Multiprocessing (多进程) in Python is often used to achieve true parallelism.<br>Use Multiprocessing for CPU-bound tasks.|
* Note: Go typically doesn't use terms "threading" and "multiprocessing" like Python does. This is because goroutines handle both concurrency and parallelism. In other words, Go abstracts these complexities (**"concurrency" vs. "parallellism"** OR **"threading" vs. "multiprocessing"**) away through goroutines and its powerful runtime scheduler. 

## Synchronization Primitives
### Mutex (Read-Write Mutex, 读写锁)
1. Rule 1: Grab a lock every time when a goroutine is accessing shared data.
2. Rule 2: Locks protect invariants. 
    - Invariants should be deemed as atomic operations.
    - Invariants may NOT be a single statement or single update to shared data. Invariants can be a combo of multiple variables (region of code).
        - 在下面的例子中，(alice + bob) 的总和是一个invariant (值应该时时刻刻保持一致)，应该视作atomic。而alice和bob各自的值不需要时时刻刻保持一致，不应该视作atomic。
3. Rule 3: Each goroutine is actively seeking to acquire the lock.
    - Actively Seeking: 任一个goroutine如果有`mu.Lock()`，可以视作它时时刻刻等着锁available，一旦锁available，它就会立刻主动积极地争抢锁的持有权，谁"手速"快就归谁。这个性质造成一定的无序性，在没有准确定义invariant的情况下，会导致shared data更新结果与预期不一致。

:x: Wrong Invariant Example:
```go
func main() {
    alice := 10000
    bob := 10000
    var mu sync.Mutex

    total := alice + bob // (alice + bob) is the invariant in this example

    go func() {
        for i := 0; i < 1000; i++ {
            /* Wrong version */
            mu.Lock()
            alice -= 1
            mu.Unlock()
            mu.Lock()
            bob += 1
            mu.Unlock()
        }
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            /* Wrong version */
            mu.Lock()
            bob -= 1
            mu.Unlock()
            mu.Lock()
            alice += 1
            mu.Unlock()
        }
    }

    // Audit: check if sum of alice+bob remains invariant = total
    start := time.Now()
    for time.Since(start) < 1*time.Second {
        mu.Lock()
        if alice+bob != total {
            fmt.Printf("Observed violation, alice =%v, bob =%v, sum = %v\n", alice, bob, alice+bob)
        }
        mu.Unlock()
    }
}
```
:white_check_mark: Correct Invariant Example:
```go
func main() {
    alice := 10000
    bob := 10000
    var mu sync.Mutex

    total := alice + bob

    go func() {
        for i := 0; i < 1000; i++ {
            /* Correct version */
            mu.Lock()
            alice -= 1
            bob += 1
            mu.Unlock()
        }
    }()

    go func() {
        for i := 0; i < 1000; i++ {
            /* Correct version */
            mu.Lock()
            bob -= 1
            alice += 1
            mu.Unlock()
        }
    }

    // Audit: check if sum of alice+bob remains invariant = total
    start := time.Now()
    for time.Since(start) < 1*time.Second {
        mu.Lock()
        if alice+bob != total {
            fmt.Printf("Observed violation, alice =%v, bob =%v, sum = %v\n", alice, bob, alice+bob)
        }
        mu.Unlock()
    }
}
```

### Condition Variable

# [Goroutine](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Concurrency vs Parallelism](#concurrency-vs-parallelism)
* [Synchronization Primitives](#synchronization-primitives)
    * [Mutex](#mutex-read-write-mutex-读写锁)
    * [Condition Variable](#condition-variable)
    * [Channels](#channels)

## Concurrency vs Parallelism
|| Concurrency <br>并发 | Parallelism <br>并行 |
| :-: | - | - |
|Concept|One core/program handles multiple tasks simultaneously.|Multiple cores/programs handle multiple tasks simultaneously. Each core handles a task respectively.|
|Characteristic|ONLY 逻辑上同时发生 (i.e., 多个任务"看上去"同时被处理)|BOTH 逻辑上同时发生，物理上也同时发生(i.e., 多个任务"确实"同时被处理)|
|Analogy|一个咖啡店员同时做三件事：下单，做咖啡，擦桌子。他在三个事之间来回切换 (context switching)。|三个咖啡店员同时各自负责其中一件事：A下单，B做咖啡，C擦桌子。|
|Python|Threading (多线程) in Python is often used to achieve concurrency.<br>Use Threading for I/O-bound tasks.|Multiprocessing (多进程) in Python is often used to achieve true parallelism.<br>Use Multiprocessing for CPU-bound tasks.|
* Note: Go typically doesn't use terms "threading" and "multiprocessing" like Python does. This is because goroutines handle both concurrency and parallelism. In other words, Go abstracts these complexities (**"concurrency" vs. "parallellism"** OR **"threading" vs. "multiprocessing"**) away through goroutines and its powerful runtime scheduler. 

## Synchronization Primitives
Resource: [Lecture 5: Go, Threads, and Raft](https://www.youtube.com/watch?v=UzzcUS2OHqo&ab_channel=MIT6.824%3ADistributedSystems)
### Mutex (Read-Write Mutex, 读写锁)
1. Rule 1: Grab a lock every time when a goroutine is accessing shared data.
2. Rule 2: Locks protect invariants. 
    - Invariants should be deemed as atomic operations.
    - Invariants may NOT be a single statement or single update to shared data. Invariants can be a combo of multiple variables (region of code).
        - 在下面的例子中，(alice + bob) 的总和是一个invariant (值应该时时刻刻保持一致)，应该视作atomic。而alice和bob各自的值不需要时时刻刻保持一致，不应该视作atomic。
3. Rule 3: Each goroutine is actively seeking to acquire the lock.
    - Actively Seeking: 任一个goroutine如果有`mu.Lock()`，可以视作它时时刻刻等着锁available，一旦锁available，它就会立刻主动积极地争抢锁的持有权，谁"手速"快就归谁。这个性质造成一定的无序性，在没有准确定义invariant的情况下，会导致shared data更新结果与预期不一致。

:x: Wrong Example:
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
:white_check_mark: Correct Example:
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
High-Level pattern to avoid race condition between wait and broadcast (lost woken-up problem):
```
mu.Lock()
// do something that might affect the condition
cond.Broadcast()
mu.Unlock()

-------------------

mu.Lock()
while condition == false {
    cond.Wait()
}
// now the condition is true, and we have the lock. We can do anything we want to do here.
mu.Unlock()
```

:x: Wrong Example:
- This is an implementation that outpus correct result. However, the issue here is that it's super inefficient! Recall Rule 3 - A thread (goroutine) that declares `mu.Lock()` is actively seeking to acquire the lock. In this example, the main thread is busy waiting for the lock. This behavior will burn up CPU! One simple solution is to add sleep in the infinite loop so that the main thread only checks the condition once a while. But this solution is not ideal. You should always be aware of the use of magic constants in your code!!! Why use 50ms rather than other numbers?
```go
func main() {
    rand.Seed(time.Now().UnixNano())

    count := 0
    finished := 0
    var mu sync.Mutex

    for i := 0; i < 10; i++ { // in parallel spqwn 10 goroutines
        go func() {
            vote := requestVote()
            mu.Lock()
            defer mu.Unlock()
            if vote {
                count++
            }
            finished++
        }()
    }

    // main thread is busy waiting
    for {
        mu.Lock()
        if count >= 5 || finished == 10 {
            break
        }
        mu.Unlock()
        // time.Sleep(50* time.Millisecond)
    }

    if count >= 5 {
        fmt.Println("Received 5+ votes!")
    } else {
        fmt.Println("Lost")
    }
    mu.Unlock()
}
```
:white_check_mark: Correct Example:
```go
func main() {
    rand.Seed(time.Now().UnixNano())

    count := 0
    finished := 0
    var mu sync.Mutex
    cond := sync.NewCond(&mu) // Decorate mutex lock with condition variable

    for i := 0; i < 10; i++ { // in parallel spqwn 10 goroutines
        go func() {
            vote := requestVote()
            mu.Lock()
            defer mu.Unlock()
            if vote {
                count++
            }
            finished++
            cond.Broadcast() // notify thread(s) marked with cond.Wait() that this goroutine has finished
        }()
    }

    // main thread is now lazy waiting: only wakes up when other threads notify
    mu.Lock()
    for !(count >= 5 || finished == 10) { // change "if" to "for"
        cond.Wait() // sleep until notified. When notified, re-acquire the lock
    }
    
    if count >= 5 {
        fmt.Println("Received 5+ votes!")
    } else {
        fmt.Println("Lost")
    }
    mu.Unlock()
}
```
`Broadcast()` vs `Signal()`: `Broadcast()` wakes up every waiter with `cond.Wait()` while `Sigal()` wakes up only one waiter. For the high-level pattern above, always use `Broadcast()`.

### Channels
- Channels can be thought as queue-like synchronization primitive but they don't have queueing capacity (no internal storage).
- Channels are synchronous. If you have 2 goroutines that are going to send and receive on a channel, if someone tries to send on the channel while nobody's receiving, then that thread will block until someone is ready to receive. At that point synchronously, it will exchange that data over to the receiver. Same is true for the other direction: if someone tries to receive from the channel while nobody's sending, then the receiver will block until someone is ready to send.
- Rule: One channel MUST have two different threads do the sending and receiving. One single thread cannot do both sending and receiving on the same channel.
```go
func main() {
    c := make(chan bool)
    go func() {
        time.Sleep(1 * time.Second)
        <- c // receive from channel c
    }
    start := time.Now()
    c <- true // block until the other goroutine receives
    fmt.Printf("Send took %v\n", time.Since(start))
}
```

:x: Wrong Example: deadlock because the main goroutine blocks until somebody's ready to receive. But there's no receiver.
```go
func main() {
    c := make(chan bool)
    c <- true // blocks thus it never detects receiver, which is the next line.
    <- c
}
```
:x: Wrong Example: deadlock because the main goroutine blocks until somebody's ready to send. But there's no sender.
```go
func main() {
    c := make(chan bool)
    <- c // blocks thus it never detects sender, which is the next line.
    c <- true
}
```
:x: Wrong Example: hang (悬停)。其中一个goroutine一直在运行，而main thread又block了，导致整个program 悬停。
```go
func main() {
    go func() { // this thread is forever running
        for {}
    } ()
    c := make(chan bool)
    c <- true // main thread blocks here
    <- c
}
```
:white_check_mark: Correct Example:
- Channels are useful for producer-consumer queue situation.
```go
func main() {
    c := make(chan int)

    for i := 0; i < 4; i++ {
        go doWork(c)
    }

    for {
        v := <-c
        fmt.Println(v)
    }
}

func doWork(c chan int) {
    for {
        time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
        c <- rand.Int()
    }
}
```
:white_check_mark: Correct Example:
- Channels are useful for achieving the same functionality as WaitGroup - Wait until all thread are done doing their work.
```go
func main() {
    done := make(chan bool)
    for i := 0; i < 5; i++ {
        go func(x int) {
            sendRPC(x)
            done <- true // send on this channel when this thread finishes its work
        }
    } (i)

    for i := 0; i < 5; i++ {
        <- done // main thread receives 5 times to have the same effect as WaitGroup
    }
}

func sendRPC(i int) {
    fmt.Println(i)
}
```
WaitGroup
```go
func main() {
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        // NOTE! wg.Add() must be called outside and before spawning a groroutine.
        wg.Add(1) // +1 waiter
        go func(x int) {
            sendRPC(x)
            wg.Done() // -1 waiter
        }
    } (i)

    wg.Wait() // block/wait until wg = 0 waiter
}

func sendRPC(i int) {
    fmt.Println(i)
}

```
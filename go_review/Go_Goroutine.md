# [Goroutine](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Concurrency vs Parallelism](#concurrency-vs-parallelism)

## Concurrency vs Parallelism
|| Concurrency <br>并发 | Parallelism <br>并行 |
| :-: | - | - |
|Concept|One core/program handles multiple tasks simultaneously.|Multiple cores/programs handle multiple tasks simultaneously. Each core handles a task respectively.|
|Characteristic|ONLY 逻辑上同时发生 (i.e., 多个任务"看上去"同时被处理)|BOTH 逻辑上同时发生，物理上也同时发生(i.e., 多个任务"确实"同时被处理)|
|Analogy|一个咖啡店员同时做三件事：下单，做咖啡，擦桌子。他在三个事之间来回切换 (context switching)。|三个咖啡店员同时各自负责其中一件事：A下单，B做咖啡，C擦桌子。|
|Python|Threading (多线程) in Python is often used to achieve concurrency.<br>Use Threading for I/O-bound tasks.|Multiprocessing (多进程) in Python is often used to achieve true parallelism.<br>Use Multiprocessing for CPU-bound tasks.|
* Note: Go typically doesn't use terms "threading" and "multiprocessing" like Python does. This is because goroutines handle both concurrency and parallelism. In other words, Go abstracts these complexities (**"concurrency" vs. "parallellism"** OR **"threading" vs. "multiprocessing"**) away through goroutines and its powerful runtime scheduler. 

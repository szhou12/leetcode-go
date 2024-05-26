# [Goroutine]()

## Contents
* [Concurrency vs Parallelism](#concurrency-vs-parallelism)

## Concurrency vs Parallelism
|| Concurrency <br>并发 | Parallelism <br>并行 |
| :-: | - | - |
|Concept|One core/program handles multiple tasks simultaneously.|Multiple cores/programs handle multiple tasks simultaneously. Each core handles a task respectively.|
|Analogy|一个咖啡店员同时做三件事：下单，做咖啡，擦桌子。他在三个事之间来回切换。|三个咖啡店员同时各自负责其中一件事：A下单，B做咖啡，C擦桌子。|
｜Python｜Threading (多线程) in Python is often used to achieve concurrency｜Multiprocessing in Python is often used to achieve true parallelism｜
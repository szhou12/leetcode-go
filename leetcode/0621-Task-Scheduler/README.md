# [621. Task Scheduler](https://leetcode.com/problems/task-scheduler/description/)

## Solution idea
### Solution 1: 全模拟 Using Max Heap
1. 理解题意: 当一个任务当下被安排以后，它不能在接下来的n个时间单位内再次安排。换句话说，在一个连续的`n+1`时间间隔内，安排`n+1`个不同的任务。当一个连续的`n+1`时间窗内，没法找全`n+1`个不同的任务，则用idle补足空缺。问: 怎么安排任务可以使总耗时最少？
2. 突破点: 怎么可以总耗时最少？很显然，用idle补空缺的机会尽量少。之所以会用idle补空缺，是因为任务的多样性不足。如何在尽可能安排不同任务的同时，尽可能多地保留剩余任务的多样性？**频次高的任务优先安排**，尽量延后使用idle的时机。如果反过来，频次低的优先安排，那么任务的多样性就会很快耗完，用idle补空缺的时机就会很快到来。

Time complexity = $O(n\log n)$

### Solution 2: Greedy
仍然先看到频次最高的task(s). 

构造一个矩阵: rows = max_freq, cols = n+1

e.g. A最高. 由A打头阵, 第一col全是A. 最后, 按row拼接起来. 这样, A不会在n+1个间隔内重复出现
```
A X X X X X
A X X X X X
A X X X X X
A X X X X X
A X X X X X
```
如果有两个以上频次最高, 做同样操作 e.g. A,B最高
```
A B X X X X
A B X X X X
A B X X X X
A B X X X X
A B X X X X
```
去掉最后一row, follow A, B 的空位:
i.e., 我们只在前 max_freq-1 rows的空位里面填充剩余的tasks
```
A B X X X X
A B X X X X
A B X X X X
A B X X X X
A B
```
填充顺序为：由上至下，再由左至右
```
A, B: 5
C: 4
D: 3
E: 2
F: 2
G: 1
H: 1
-----------
A B C D E H
A B C D F X
A B C D F X
A B C E G X
A B
```
剩余的空位填上idle

这样填充可以保证连续n+1个时间内不会重复出现用一个任务

proof:

C的频次最高不超过max_freq - 1, 而空位本来就有max_freq - 1行。

因为填充顺序是由上至下，所以不可能同一行会填充两次C.

同理，可推论D, E, F, G, H...

又因为每一行的长度为n+1, 而从左到右填充的顺序保证了下一次安排同一个任务的位置所在列数不小于这一次安排的列数,

这样, 就保证了n+1内不会重复出现同一个任务。

填充后的矩阵会有两种情况：
1. 情况一: 空位填不满，需要idle补缺
min-time = (max_freq - 1) * (n+1) + 零头(=最高频次tasks的种类个数)
2. 情况二: 空位不够填，有更多的tasks等着被填
前max_freq-1行依次 (从上到下再从左到右) 扩充列数安排余下的tasks.

min-time = total # of tasks
```
A, B: 5
C: 4
D: 3
E: 2
F: 2
G: 1
H: 1
I: 1
J: 1
K: 1
l,m,n,o,p,q: 1
----------------
A B C D E H l p
A B C D F I m q
A B C D F J n
A B C E G K o
A B
```

最后，两种情况的min-time取max (注意: 是取最大值哦!)，即为答案。

Time complexity = $O(n)$

## Resource
[【每日一题】621. Task Scheduler, 8/8/2021](https://www.youtube.com/watch?v=3DZE7cfgYyg&t=161s&ab_channel=HuifengGuan)
- Solution 1: 0:00 - 20:22
- Solution 2: 20:00 - end
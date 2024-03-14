# [3049. Earliest Second to Mark Indices II](https://leetcode.com/problems/earliest-second-to-mark-indices-ii/description/)

## Solution idea
### Binary Search 猜答案 + Regret Greedy (MinHeap)
1. 理解题意: 与[3048](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3048-Earliest-Second-to-Mark-Indices-I)不同的是，本题中mark操作也需要占用一单位时间。每一个时刻t可以做三件事情:
    1. decrement: 消耗一单位时间，挑选任意一个`nums[]`中一个非0元素减1
    2. set 0: 直接把`nums[changeIndices[t]]`一次性清0
    3. mark: 消耗一单位时间，挑选任意一个`nums[]`中一个0元素做马克
2. 首要直觉：
    1. **单调性** - 时间越长，越容易清0，越容易mark尽可能多的位置。4秒钟能完成任务，那么8秒钟也一定能完成任务。**单调性** + **最优问题** => **Binary Search**
    2. 比起-1操作 (decrement), 优先使用清0操作 (set 0), 能早做尽早做。因为要留时间给mark操作。
3. 如果对于一个 `nums[] idx`，它在时间序列`changeIndices[]`里出现了多次，我们会在哪个时候去清0呢？
    * 相对而言我们尽早清0是最优的选择，因为如果晚些时候去做清0操作，可能存在一个风险：后续没有足够的机会取做mark操作了。由此，我们可以在时间序列`changeIndices[]`里面，预处理得到每个`nums[] idx`第一次出现的位置，好进行清0操作。
4. 在[1, t]时间段内，先确定好清0操作的时刻以后，剩下的时刻优先-1操作。同时，我们要在最后留有足够的时间来进行mark操作 (足够是指`nums[]`每个位置都要留有一次mark操作)。
    * 注意！并不是无脑的选最后n秒都做“mark”，因为有些“清0”可能在很靠后的时刻才会发生。怎么制定策略呢？这就需要从后往前去安排。
    * 每往前推进一次时刻，我们要始终确保mark操作的数量 $\geq$ 清0操作的数量。
    * 如果往前推到某一时刻发现，mark操作的数量 < 清0操作的数量，那么我们需要吐出一个清0的位置给mark操作，恢复平衡。
    * 吐哪一个？优先选对应`nums[] val`最小的那一个！为啥？我们要保证清0的收益最大化，把值最小的留给-1操作去完成，保留值大的给清0操作。
    * 实现：维护一个**MinHeap**, 存`nums[] val`, 每次吐出最小值。
```
O - first appearance of nums[] idx - for set 0 operation
X - for -1 or mark operation
1  <-  i  <-   t  ... m
X X X O O X O [X] // 最后一个位置不做清0，留给mark操作，这样安排okay
X X X O O X [O X] // 最后两个位置一个清0，一个mark，这样okay
X X X O O [X O X] // 最后三个位置一个清0，两个mark，这样okay
X X X O [O X O X] // 最后四个位置两个清0，两个mark，这样okay
X X X [O O X O X] // 最后五个位置三个清0，两个mark，这样不okay！因为有三个清0的位子需要mark，但是只有两次mark机会。这时候，我们需要选择一个清0的位置吐出来，让给mark操作。
```
5. 保持这样从后往前走到头:
    * 清0操作数量 = MinHeap中剩余的元素个数
    * mark操作数量 = `nums[]`的长度
    * -1操作数量 = t - 清0操作数量 - mark操作数量
6. 最后判断能否完成题目要求: 清0操作清掉的val总和 + -1操作数量 $\geq$ `nums[]`所有元素的val总和
7. 注意！每次检查一个时刻t, 要保证 t >= `nums[]`的长度，因为要至少保证有足够的时间来mark。

Time complexity = $O(\log n \cdot n\log n)$

## Resource
[【每日一题】LeetCode 3049. Earliest Second to Mark Indices II](https://www.youtube.com/watch?v=VA6TLsOVMa4&ab_channel=HuifengGuan)
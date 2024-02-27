# [1353. Maximum Number of Events That Can Be Attended](https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended/description/)

## Solution idea
### Greedy: interval scheduling + Soln: sorting + minHeap
1. 理解题意：本题和常规的scheduling问题不同的一点是，一段时间内可以做多个jobs，每个job完成只需要花费一天。e.g., job A: [1, 2], job B: [1, 2]. 两个job都可以在day 1-day 2这个区间内完成：job A在day 1完成，job B在day 2完成。
2. 突破点：理解题意后，需要有一个朴素的想法：既然一天只能完成一个job, 那么在所有可以在这一天进行的job中 (注意！这里的视角是某一天!)，选择哪个最合适？答：手头有很多作业要完成，先做哪个？显然是优先选择做due date早的作业！先完成due date早的作业，保证这个作业能做完，然后再看能不能补救due date晚的作业。如果先做due date晚的作业，那么due date早的作业铁定就错过了。落实到这道题: **Given a day, 优先选择end date早的event。**
3. 还需要解决一个问题：**Given a day, 哪些event是可以在这一天进行的？**答：看start day. 只要event的start date <= 当前day，那么这个event就可以在这一天进行。
4. 代码实现:
    * STEP 1: sort events by start date. Why? 因为这样做，在下一步loop每一天的时候，我们就可以得到所有可以在这一天进行的event。<- 对应解决第3条
    * STEP 2: loop days so that given a day:
        * 2.1: 挑出start date <= 当前day的event。因为STEP 1排序过了，所以可以不回头地一直往前看 (用`i`标识)。把这些event的end date放入minHeap中。<- 对应第2条，优先due date早的
        * 2.2: 排除掉minHeap中，end date < 当前day的event。因为这些event已经无法在当天进行。
        * 2.3: 如果minHeap不为空，那么就可以从minHeap中蹦出“最紧迫的”event在这一天进行。<- 对应第2条，优先due date早的

Time complexity = $O(max\{10^5, n\logn\})$ = $O(n\log n)$
注意！loop days和内层挑event的循环是相互独立的。也就是说，loop $10^5$ days, minHeap中只会往里放入n个event并且每个event只会放一次。所以是取max的关系。

## Resource
[【每日一题】1353. Maximum Number of Events That Can Be Attended, 6/13/2021](https://www.youtube.com/watch?v=9bJvSySPcZM&ab_channel=HuifengGuan)
# [787. Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/description/)

## Solution idea

### Method 1 - Dijkstra

#### 要点总结
1. 这道题的难点在于有中转站次数限制，找最省钱的路线我们知道可以用Dijkstra，但怎么把中转站次数限制这个信息融入到Dijkstra找最短路径的过程中呢?
    * **突破点: 每个node记录双状态 - `[node label, # stops used]`**
        * 到达一个node时, 不仅要关注这是哪一个node (node的位置, 用node label表示), 还要关注到达这个node时已经用掉了多少次中转 (# of stops used so far).
        * i.e., node i = (i, stops)
        * 这道题告诉我们: 一个node不是只能储存单条信息(node位置/node是谁)，也可以储存多条信息. 就想DP不是只能一维，也可以有二维DP甚至更高维.

2. 跟以往一维node的Dijkstra题目不太一样的是: 这道题Dijkstra return在loop里, 表示如果此时走到了终点, 那么第一次pop出来的一定是最短距离(并且满足中转站次数限制); 反之, 如果跳出了loop, 说明终点是无法到达的. return在loop里还有一个原因: 现在的code写法是如果到达一个node时刚好用掉k+1次, 会被跳过, 所以, 如果到达终点刚好要k+1次中转, 目前这样写的code不会记录到达终点所需的最小花费，反而保留-1表示无法到达. return在loop里还有另一个原因是: 这样写code简洁, 不需要收尾工作 (指输出时一个一个去查表格看dst 用掉几次中转花费最少)
    * 这样写还有一个值得学习的地方: Dijkstra算法的物理意义就是第一次pop出来的一定是最短距离, 所以第一次pop出来终点就说明已经找到答案, 直接return就好了, 不用等到loop结束跳出

3. 实现细节: 题目给的 `k` 表示起点和终点之间的中转站次数, 它没有包含跳到终点的最后一步. 所以, 实际上的限制应该是`k+1`, 表示从起点最多跳`k+1`步到达终点.

Time complexity = $O(E\log E)$

### Method 2 - DP

## Resource
[【每日一题】LeetCode 787. Cheapest Flights Within K Stops](https://www.youtube.com/watch?v=Q8oMHlThySQ&ab_channel=HuifengGuan)
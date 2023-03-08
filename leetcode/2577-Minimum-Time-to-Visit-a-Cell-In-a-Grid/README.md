# [2577. Minimum Time to Visit a Cell In a Grid](https://leetcode.com/problems/minimum-time-to-visit-a-cell-in-a-grid/description/)

## Solution idea

### Dijkstra

#### 要点总结
1. 题目是在一个矩阵上下左右走暗示了是图问题 (DP只适合解向下和向右走的问题, 上下左右走无法做到DP要求的无后效性, 因为可能回到之前的状态做更新) $\Rightarrow $ 题目"Minimum Time" 暗示了求最短路径 $\Rightarrow $ **解题框架: 用 Min Heap 实现 Dijkstra 算法** + **入PQ的初始值是谁** + **走下一步时Push到达时间需要分类讨论**
    * 无后效性: 
        * 解释1: 某阶段的状态一旦确定，则此后过程的决策不再受此前各种状态及决策的影响
        * 解释2: 对于一个确定的状态，我们不必关心这个状态是怎么出现的，也不必考虑这个状态的前一个状态是什么
        * 解释3: 未来的状态和决策不会影响以前的状态、决策和目标。以前发生过的事就是板上钉钉的事，不会更改。
    * 有后效性: 就是某个状态之后要做的决策会受之前的状态及决策的影响
2. 题目中的关键句 - **you must move to any adjacent cell in the four directions** $\Rightarrow $ **每一时刻t必须移动, 不能原地停留**
3. 每个cell有一个准入时间(门槛) $\Rightarrow $ **如果尝试进入的时刻没到准入时间, 就需要我们在门外“反复横跳”耗去一部分时间再进入**
    * 反复横跳: 在两个相邻格子间跳来跳去
4. **一定要能迈出第一步 (从起点cell走出来)** $\Rightarrow $ **才有"反复横跳"的资本** $\Rightarrow $ **才可以走到终点cell**
    * 这说明起点cell的邻居们不能有太高的准入时间 (i.e. $\leq 1$); 否则, 我们就困死了, 根本走不出来
5. 走下一步时怎么Push到达时间需要分类讨论:
    * Case 1: 到达下一个cell用到的时间 $t + 1 \geq $ 下一个cell的准入时间: Push $t+1$
    * Case 2+3: 否则, 需要在当前cell和到达过的当前cell的邻居之间"反复横跳"耗时间; *注意*: 反复横跳一来一回需要2个单位的时间, 所以反复横跳总时间需要`2*横跳次数`(偶数倍); *再注意*: 从当前cell走到下一个cell还要花费1单位时间
        * Case 2: 下一个cell准入时间 - 当前cell到达时间$t$ = 奇数: Push 下一个cell准入时间 (刚好可以掐表进入)
        * Case 3: 下一个cell准入时间 - 当前cell到达时间$t$ = 偶数: Push 下一个cell准入时间+1 (准入时刻的下一秒进入)

#### 代码实现总结
Step 1: 先把 PQ实现Dijstra 的大框架实现

Step 2: Push部分需要分类讨论

Step 3: 添加code: 1. check 是否可以迈出第一步; 2. Push的初始值 从起点cell替换成它右侧和下边的邻居


Time complexity = $O(mn\log mn)$ (个人感觉这只描述了下限, 没有计算入反复横跳多花出去的时间)

## Resource
[【每日一题】LeetCode 2577. Minimum Time to Visit a Cell In a Grid](https://www.youtube.com/watch?v=bQ-ZMe2Udtw&ab_channel=HuifengGuan)
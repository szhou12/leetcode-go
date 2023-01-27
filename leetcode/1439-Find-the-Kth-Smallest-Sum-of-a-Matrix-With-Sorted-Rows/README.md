# [1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows](https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/description/)

## Solution idea

### Brute force - Greedy Algorithm
* `sum` array: 前i行个array的所有可能sum，从小到大排列，前k个最小的sum ($0 \leq i \leq m-1$ where $m$是总行数)
* 逐行把当前行的每个元素加入`sum` array中；重新排序并且取 top k smallest sums

Time complexity = $O(m * (kn + kn* \log kn)) = O(m*k*n* \log kn)$

### Binary Search 猜答案

* 收敛条件: `mat` 每一行选一个element 组成的 array (sum <= mid) 的个数至少有 k 个, 满足这个条件, mid有可能是答案

* 如何获得所有可能的array组合？
    * Ans: DFS - All Combination
    * 有几层 = `mat` 的行数
    * 每层的物理意义 = `mat` 的一行
    * 每层几个分支 = `mat` 下一行的所有element (0...n-1)
    * Base case: 满足条件 (sum <= mid) 的array 个数 (i.e. count) >= k, 返回; 或者已经到底 (`mat`最后一行), 返回.

Time complexity = $O(k * \log D)$ where $D = $ difference between the largest sum and the smallest sum of `mat`, $k = $ DFS 最多看前k小个 array sum



## Resource

* [【每日一题】1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows, 6/17/2021](https://www.youtube.com/watch?v=y9EzW-S63Vo&ab_channel=HuifengGuan)
    * Brute force: 0:00 ~ 10:33
    * k-way merge: 10:43 ~ 32:00
    * Binary search: 32:01 ~ end
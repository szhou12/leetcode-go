# [1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows](https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/description/)

## Solution idea

### Brute force - Greedy Algorithm
* `sum` array: 前i行个array的所有可能sum，从小到大排列，前k个最小的sum ($0 \leq i \leq m-1$ where $m$是总行数)
* 逐行把当前行的每个元素加入`sum` array中；重新排序并且取 top k smallest sums

Time complexity = $O(m * (kn + kn* \log kn)) = O(m*k*n* \log kn)$


## Resource

* [【每日一题】1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows, 6/17/2021](https://www.youtube.com/watch?v=y9EzW-S63Vo&ab_channel=HuifengGuan)
    * Brute force: 0:00 ~ 10:33
    * k-way merge: 10:43 ~ 32:00
    * Binary search: 32:01 ~ end
# [354. Russian Doll Envelopes](https://leetcode.com/problems/russian-doll-envelopes/)

## Solution idea

* 每个元素 `=(w, h)`， 都是二维的，需要先按照一个维度递增排序，再在另一个维度找 Longest Increasing Subsequence (LIS)

* `w` 相同的两个元素，应该再按照 `h` 递增排列，还是递减排列？
    * 答：递减排列，因为这样我们在找 LIS 时，就可以完全剥离(忽略) `w` 维度, 而只考虑`h` 维度
    * 因为题目要求两个维度严格递增，也就是套娃只能发生在两个维度都变大的情况下
    * Example: (4, 2), (4, 6)
        * 如果 `h` 递增排列得到: (4, 2), (4, 6)
            * 在构建以(4, 6)结尾的LIS时，除了要`h`满足要求，仍然要检查前一个元素`w` < 当前元素的 `w`, 所以算法中还是要保留`w`的信息
        * 如果 `h` 递增排列得到: (4, 6), (4, 2)
            * 在构建以(4, 2)结尾的LIS时，(4, 6)会因为在`h`阶段就不满足要求而被踢出，这样就不再需要检查`w`了，所以算法中不再需要保留`w`的信息

### DP - LIS 会超时！！！

* Step 1: Sort by `w` in increasing order

* Step 2: DP

* Definition

```
DP[i] = max # of envelopes we can Russian Doll ending at envelopes[i]
```

* Base Cases
```
DP[i] = 1 for all i because the min # of envelopes can be itself
```

* Recurrence
```
dp[i] = max(dp[j] + 1) for 1 <= j < i and i's (w, h) > j's (w, h)
```

Time complexity = $O(n\log n + n^2) = O(n^2)$


### Greedy + Binary Search
* Step 1: Sort by `w` in increasing order; then if `w` identical, sort by `h` by decreasing order
    * Only like this cane we totally ignore `w`-dimension

* Step 2: 维护一个数组A
    * Case 1: 如果当前元素大于所有目前数组A内的元素, 直接append进A
    * Case 2: 如果当前元素小于某些目前数组A内的元素, 找到数组A内第一个比它大的元素并用它替换掉
        * Binary search to find the first element larger than the target


Time complexity = $O(n\log n + n\log n) = O(n\log n)$


## Resource
* Greedy at 15:00 [【每日一题】354. Russian Doll Envelopes, 12/17/2020](https://www.youtube.com/watch?v=B1d2wV6mfzA)
# [31. Next Permutation](https://leetcode.com/problems/next-permutation/)

## Solution idea

step 1: 从后往前，找到第一个降序的元素 A[i]

step 1b: 如果整个数组从后往前是升序排列，即，数组降序排列，则直接返回升序排列的数组 eg. [3,2,1] -> [1,2,3]

step 2: 从后往前，找到第一个比 A[i] 大的元素 A[j] (find first j s.t. A[j] > A[i])

step 3: swap(A[i], A[j])

step 4: 对 i 之后的所有元素(一定还是降序的)进行升序排列

```
[1, 2, 8, 4, 7, 6, 5, 3]
          i <----------
                   j <-
[1, 2, 8, 5, 7, 6, 4, 3]
             |---------|
[1, 2, 8, 5, 3, 4, 6, 7]
```

## Resource
[贾考博 LeetCode 31. Next Permutation](https://www.youtube.com/watch?v=K-QCteGM-Bk)
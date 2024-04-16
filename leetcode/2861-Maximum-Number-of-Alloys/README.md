# [2861. Maximum Number of Alloys](https://leetcode.com/problems/maximum-number-of-alloys/description/)

## Solution idea
### Binary Search 猜答案
1. 题目要求只能使用一台机器制造alloys，容易想到遍历每种机器，分别找出可以造出的alloys上限，再取最大值。每台机器可以造出的上限可以通过Binary Search猜答案来快速查找，收敛条件是猜的alloys数不超过budget。
2. 骨架：
```
loop k machines:
    binary search max # of alloys machine i can produce without exceeding buedget
```
3. 注意！Binary Search的最高上界不能定得太高 (e.g. `right = math.MaxInt` or `right = math.MaxInt/2`)，都会在计算cost时发生溢出。网上的参考答案可以给到 `right = int(1e9)` (`int(1e8)`不行)，猜测可能与constraint `0 <= stock[i] <= 10^8`有关。

Time complexity = $O(k\cdot \log(10^9) \cdot n)$
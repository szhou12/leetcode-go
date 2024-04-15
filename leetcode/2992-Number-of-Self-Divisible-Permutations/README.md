# [2992. Number of Self-Divisible Permutations](https://leetcode.ca/2024-01-04-2992-Number-of-Self-Divisible-Permutations/)

## Solution idea
### DP + Bit Manipulation
1. 突破点：从constraint入手 `1 <= n <= 12`. 如果是暴力全排列，时间复杂度是$12! \appox 4*10^8$，太高了。进而考虑把时间复杂度降到 $2^12 = 4096$，这样好像可以接受了。那么，这道题的时间复杂度的可接受范围大概就在 `n * 2^n`或者 `n^2 * 2^n`这个范围。
2. 如果时间复杂度中有`2^n`能想到什么？容易想到n个位子，每个位子的元素选或者不选。这个想法应用到本题上就是表示“状态”。
3. `2^n` 表示什么？state: `2^n` 表示的是每个数字(1~n)的状态，是否在使用中。
    - e.g. 1001000111 表示 1, 4, 7, 8, 9 位子上的数字使用过了
    - 所以，我们可以使用一个长度为`2^n`的bit array来表示state
    - 注意！题目中1~n是1-indexed。bit array是0-indexed，所以在状态转移是需要把1-indexed转换为0-indexed。同时！bit array的顺序是从右至左：最右边是0号位，最左边是n-1号位。
```
digit  n   n-1    ...    ...   3   2   1
index  n-1 n-2    ...    ...   2   1   0
state [0,  1,  0,  0,  1,  0,  0,  1,  1]
```
4. 二维DP
    1. Definition: `DP[i][state] := ` total number of **self-divisible** permutations from the first `i` positions by using the digits marked as 1 in `state`.
    2. Base Case: `DP[0][0] = 1` 表示前0个元素，状态是没用使用任何元素的情况。一个'空'的排列，是一种合法的排列
    3. Recurrence: `DP[i][state] += DP[i-1][state - (1<<(digit-1))]` for all `1 <= digit <= n` where `gcd(i, digit)==1` (self-divisible) and `state & (1<<(digit-1)) == 1` (表示state的array中digit-1位子上是1表示digit处于使用状态)

Time complexity = $O(n*2^n*n)$

## Resource
[【每日一题】LeetCode 2992. Number of Self-Divisible Permutations](https://www.youtube.com/watch?v=LHOAR6uSoFA&ab_channel=HuifengGuan)
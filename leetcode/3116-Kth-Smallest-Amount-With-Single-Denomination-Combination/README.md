# [3116. Kth Smallest Amount With Single Denomination Combination](https://leetcode.com/problems/kth-smallest-amount-with-single-denomination-combination/description/)

## Solution idea
### Binary Search (Guess k) + Grosper's Hack (Combinatorics) + Exclusion-Inclusion Principle + Divisors
1. 理解题意: 将所有组合数 $a_0coins[0] + a_1coins[1] + \cdots + a_{n-1}coins[n-1]$ where $a_0, a_1, \cdots, a_{n-1} = 0\cdots \infty$ and not all $a_i = 0$ at the same time 从小到大排列，找到第 $k$ 个数。换言之，找到第 $k$ 个小的数，使得它能被`coins[]`里任意一个元素或者多个元素的组合整除。
2. 容斥原理: 根据题意，所有的组合数都能被`coins[]`里的元素或者元素组合整除，那么根据：**`[1, ..., N]` 自然数区间内能被 `d` 整除的个数 = `N / d`。**组合数的数量就是能某一元素/元素组合被整除的数量。
    - e.g. For a given number `N`, `coins = [a, b, c]`
    ```
    总组合数 = + count(N%a==0) + count(N%b==0) + count(N%c==0) 
              - count(N%a==0 && N%b==0) - count(N%a==0 && N%c==0) - count(N%b==0 && N%c==0)
              + count(N%a==0 && N%b==0 && N%c==0)
    ```
    - 注意: 奇数个元素时为+，偶数个元素时为-
3. constraint: `1 <= k <= 2 * 10^9`. 说明$k$很大，无法一一罗列组合数找第 $k$ 个数。所以需要二分搜值：从`[1, mid]`里面数有多少个能被`coins[]`里的元素或者元素组合整除。
4. 如何高效地枚举任意X个元素的组合？e.g. `[a, b, c, d]`中由三个元素的组合有多少个: `0111`, `1011`, `1101`, `1110`. 使用[Grosper's Hack](https://github.com/szhou12/leetcode-go/blob/main/mathematics/GospersHack.go).

Time complexity = $O(\log n \cdot {n \choose k})$


## Resrouce
[【每日一题】LeetCode 3116. Kth Smallest Amount With Single Denomination Combination](https://www.youtube.com/watch?v=R63BH6UDQXQ&ab_channel=HuifengGuan)
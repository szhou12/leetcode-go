# [3449. Maximize the Minimum Game Score](https://leetcode.com/problems/maximize-the-minimum-game-score/description/)

## Solution idea
### Binary Search Guess Value + Greedy
1. 理解题意: **maximum possible minimum** means to maximize possible min. i.e., 在m次moves之后，`gameScore[]`每个元素的值有大有小，我们要把其中最小值尽可能最大化。
2. 如何破题:
    1. observation 1: To maximize, we will definitely use up ALL m moves.
    2. observation 2: Look at contraint, `m <= 10^9`, it means that DP solution or enumeration is not possible. Likely will use binary search.
    3. How to set up binary search: Guess a target score X
        - Definition: X = can we make all positions have points >= X after m moves?
        - if yes, current guessing is a possible answer.
        - if no, current guessing is too large, we need to squeeze down.
3. Okay, now we need to figure out how to fill up position `i` so that it achieves X.
    - Greedy Strategy: 当在0号位子时，如果pre-filled的分数 < X，说明我们需要补齐“差价”。怎么补齐？通过向右跳到1号再跳回0号位子 (`0 -> 1 -> 0`)的反复横跳。注意，显然不会跳到更远的2号位子再跳回来 (`0 -> 1 -> 2 -> 1 -> 0`) 因为这样做0号位子补差价的次数没有变(之前要叠加n次，现在还是要叠加n次)，然而你却额外消耗了moves。假设0-1左右横跳d次(d次来回)，0号可以达到X，这时1号位子pre-fill了`d+1`次(d次来回 + 1次跳到1号)`points[1]`。
    - 综上：在`i`号位子，先检查pre-filled的分数是否 >= X，如果已经满足，就可以直接跳到`i+1`；如果不满足，先计算出补齐`i`号位子差价需要的反复横跳次数`d`，补齐后，再跳1次到`i+1`，同时准备`i+1`的pre-fill分数为`points[i+1] * (d+1)`。
    - 两个需要特殊处理的边界条件:
        1. 如果走到了`n-1`，如果此时pre-filled的分数还不够，补齐的方式是向左反复横跳 (`(n-1) -> (n-2) -> (n-1)`)
        2. 如果走到了`n-2`，如果我们在这里进行了反复横跳，也就是给`n-1`pre-fill了一定分数，我们需要检查这些pre-filled的分数是否 >= X，如果是肯定的，那就没必要再跳到`n-1`了，可以提前结束。
    - 注意，反复横跳次数`d`需要向上取整。如果差价/单价是小数 (e.g. 1.4次)，我们要跳2次来回。
        - **向上取整技巧:** `ceil(a/b) = (a-1)/b + 1` 


Time complexity = $O(n \log X)$ where $X = 10^{15}$, $n$ is length of `points[]`

## Resource
[【每日一题】LeetCode 3449. Maximize the Minimum Game Score](https://www.youtube.com/watch?v=N6aScon-ehY&t=30s&ab_channel=HuifengGuan)
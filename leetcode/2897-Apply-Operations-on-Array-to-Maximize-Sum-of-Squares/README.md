# [2897. Apply Operations on Array to Maximize Sum of Squares](https://leetcode.com/problems/apply-operations-on-array-to-maximize-sum-of-squares/description/)

## Solution idea
### Greedy + Bit Operation
1. 搞明白题干中 bitwise AND, OR 的意义:
    * 可以发现，OR 操作像 “磁铁” 一样，会把 1 “吸” 过来，而 AND 操作则会把 1 “排斥” 出去。经过 AND 和 OR 操作后，x, y 两个元素的 “贫富差距” 进一步拉大。 
    ```
    x   y  => (x AND y)  (x OR y)
    1   1        1          1
    1   0        0          1
    0   1        0          1
    0   0        0          0
    ```
2. 两个较为平均的数的平方和 < 两个较为极端的数的平方和
    * e.g.
$2^2 + 2^2 < 1^2 + 3^2$
    * proof: if x <= y and x, y >= 0, d > 0, then
   $$
   x^2 + y^2 < (x-d)^2 + (y+d)^2 = x^2 + y^2 + 2d(y-x) + 2d^2
   $$
3. 落实到解题:
    1. Step 1: 把 `nums` 里每个元素 "打散" 成 32-bit表示。用一个 `counter[]` 计算每一个bit位出现的 1 的频次.
    2. Step 2: 选 k 个元素 相当于 loop k 遍 从大到小选元素。如何从大到小选？通过 OR 操作把每一个 bit位的 1 尽可能往“上”吸。每个数的大小想象water level，我们想把 1 尽可能地“浮”在水面上。因为这样越靠近“水面”的数每个bit位的1越多，它的值就越大。

Time complexity = $O(32n) + O(32k) = O(n)$



## Resource
[【每日一题】LeetCode 2897. Apply Operations on Array to Maximize Sum of Squares](https://www.youtube.com/watch?v=96xk1INe134&t=637s&ab_channel=HuifengGuan)

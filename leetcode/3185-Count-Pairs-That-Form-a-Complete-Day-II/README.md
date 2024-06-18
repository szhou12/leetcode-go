# [3185. Count Pairs That Form a Complete Day II](https://leetcode.com/problems/count-pairs-that-form-a-complete-day-ii/description/)

## Solution idea
### Two Sum + Remainder + Frequency
1. 题目求两个元素的和得是24的倍数。题干设置容易联想到Two Sum：用一个元素，去HashMap里找它的Complement。
2. 因为两个元素之和是24的倍数，所以单纯地只关注和的话，范围可以无限大。为了把限定一个范围，考虑取余: $(a+b)\% 24 = a\% 24 + b\%24 \in [0, 23]$。根据分配律，我们分别对每个元素取余。
3. 用一个frequency map记录每一种余数的频次: `{一种余数 : 频次}`
4. Two Sum型遍历每一种余数，并分类讨论:
    1. 余数 = 0：说明原数是24的倍数，这样的原数只能和同样是余数是0的配对，所以总配对数 = `freq * (freq - 1)/2`
    2. 余数 = 12：说明原数是12的倍数 (12, 36, ...) 并且不是24的倍数，这样的原数只能和同样的数配对，所以总配对数 = `freq * (freq - 1)/2`
    3. 余数不属于上述两种情况: 看它的补是否存在，存在的话，总配对数 = `freq_a * freq_a_complement`。注意！这里我们只检查小于它的补的情况。即，添加判定条件`a < 24-a`。这样做可以避免重复计算。
        * 为什么不使用 `freq_a * freq_a_complement/2` 的形式避免重复计算？这是因为除以2的操作会向下取整，导致漏算的情况。考虑`(13, 35)`这两个数 (它们对应的余数是`(13, 11)`)，假如它们只出现过一次，如果先看到13，符合的配对数为 `1*1/2 = 0`；下一次看到35(11)，符合的配对数为 `1*1/2 = 0`。这样就漏算了一个配对数。

Time complexity = $O(n)$
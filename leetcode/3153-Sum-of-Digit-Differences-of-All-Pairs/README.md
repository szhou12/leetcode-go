# [3153. Sum of Digit Differences of All Pairs](https://leetcode.com/problems/sum-of-digit-differences-of-all-pairs/description/)

## Solution idea
### Frequency + Combination
1. 理解题意: 给一串数`nums[]`，每个数都有相同的位数。任意挑两个数，比较个，十，百，千...每个数位上不同值count 1，相同值count 0。求所有pairs都比较数位后的count总和。
2. 突破点: `nums[]`的长度constraint在 `10^5`，无法iterate所有pairs (Grosper's Hack也不行，当`n`太高时，for loop都进不去)。有什么方法可以$O(n)$或者至多$O(nlogn)$的时间检查所有pairs?
    - “频率法”：注意到每个数位上的值有且只有0-9，所以可以记录每个数位上的值的频率。
    - `count[i][0...9] :=` 第i位上值为0-9时分别的频次。
    - 那么，第i位上值不同的pair数 = `count[i][x] * (n - count[i][x])`，其中`x`为0-9。
    - 总pair数就是所有位数相加，最后结果除以2 (因为每个pair都被计算了两次 `(a,b) = (b,a)`)。
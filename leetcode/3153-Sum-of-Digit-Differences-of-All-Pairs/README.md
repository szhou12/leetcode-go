# [3153. Sum of Digit Differences of All Pairs](https://leetcode.com/problems/sum-of-digit-differences-of-all-pairs/description/)

## Solution idea
### Frequency + Combination
1. 理解题意: 给一串数`nums[]`，每个数都有相同的位数。任意挑两个数，比较个，十，百，千...每个数位上不同值count 1，相同值count 0。求所有pairs都比较数位后的count总和。
2. 突破点: `nums[]`的长度constraint在 `10^5`，无法iterate所有pairs (Grosper's Hack也不行，当`n`太高时，for loop都进不去)。
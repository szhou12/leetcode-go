# [2851. String Transformation](https://leetcode.com/problems/string-transformation/description/)

## Solution idea
### KMP + Recursion + State Transformation Matrix
1. 首先，必须要注意到：题目中 shift `s` 的机制 相当于 扑克切牌 => cyclic string rotation
Let +k = move suffix of length k to the front
那么，每一轮是等可能地选择其中一种 shift 的形式:
round 1: +1, +2, +3, ..., +(n-1)
round 2: +1, +2, +3, ..., +(n-1)
round 3: +1, +2, +3, ..., +(n-1)
...
round k: +1, +2, +3, ..., +(n-1)
Time = $O((n-1)^k)$ (计算每一条路径会超时)
2. 因为机制相当于扑克切牌，所以每个字母的相对位置“不变”。切 k次 相当于 在合适的位置只切一次。
Assume `s'` is `s` after k shifts, `s'` can only be one of the following forms: 
`s' = s(0), s(1), s(2), ..., s(n-1)` where `s(x)` means shifting the suffix of `s` starting at index x
then how do we know if t is among them? => **double the string s to find cyclic shift**
Trick: Since `s(0)` == `s(n)` == no shift at all (will cause double counting), cutoff the last char when double the string `s`. e.g. `s = abababa` => `2s-1 = ababab|ababa`
Find the number of appearances of t (pattern string) in 2s-1 (target string) => **KMP**
3. 

## Resource
[【每日一题】LeetCode 2851. String Transformation](https://www.youtube.com/watch?v=l2hNd3BlHkc&t=2s&ab_channel=HuifengGuan)
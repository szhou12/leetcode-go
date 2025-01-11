# [3399. Smallest Substring With Identical Characters II](https://leetcode.com/problems/smallest-substring-with-identical-characters-ii/description/)

## Solution idea
### Binary Search 猜答案 + Two Pointers + Greedy

intuition: binary search

why? 隐含单调性 - more operations given, more likely we can cut the str into pieces,
and thus, shorter substrings we can have

what are we gussing? guess the upper limit that the longest substr can exist.
-> assume a substr length k, identify # of ops we need

统计s中identical substr长度 (two pointers) -> 对于每一个长度，看需要切多少刀。

贪心: 假设现在允许identical的长度上限为k, 则对于一个总长为x的identical substr, 一个循环节就是 (k+1)。

rightend最多只能保留k长identical。总刀数 = ceil((x-k)/(k+1)). 无法整除需要取上界, 多一刀. 

e.g. 11/4=2+1, 12/4=3, 13/4=3+1

特殊情况: k=1, 0110就不再适用。跳出贪心，单独计算k=1的情况：either 010101... or 101010...

## Resource
[【每日一题】LeetCode 3399. Smallest Substring With Identical Characters II](https://www.youtube.com/watch?v=AdT2F0x-uKo&ab_channel=HuifengGuan)
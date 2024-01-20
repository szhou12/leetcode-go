# [3008. Find Beautiful Indices in the Given Array II](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-ii/description/)

## Solution idea
### KMP + Binary Search
1. 理解题意：在字符串`s`中找到子串`a`出现的一个起始位置`i`以及子串`b`出现的一个起始位置`j`，要求满足`|j-i|<=k`。找出所有满足条件的`i`。
2. 字符串中寻找一个子串是否出现以及出现的位置 -> KMP
    * KMP模版：寻找子串所有出现的位置
    * `aIndices`: 所有`a`出现的起始index位置
    * `bIndices`: 所有`b`出现的起始index位置
3. 题目要求：$|j-i| \leq k \Rightarrow -k \leq j-i \leq k \Rightarrow i-k \leq j \leq i+k$
    * 翻译：对于每一个`i`，符合条件的`j`要在`[i-k, i+k]`的区间范围
    * 实现：
        1. 经过 KMP 算法，`bIndices` 一定是递增的
        2. $i-k \leq j$: 使用 `lowerBound(bIndices, i-k)` 找到第一个(最小的)`j >= i-k` ($j_{s}$)
        3. $j \leq i+k$: 使用 `upperBound(bIndices, i+k)` 找到第一个(最小的)`j > i+k` ($j_{e}$)
        4. $j_{e} - j_{s}$ 是长度，也是区间`[i-k, i+k]`内`j`的个数。个数>=1，说明目前的`i`是符合要求的。 
    
Time complexity = $O(n\log n)$

## Resource
[【每日一题】LeetCode 3008. Find Beautiful Indices in the Given Array II](https://www.youtube.com/watch?v=_SBlKX9zc_0&ab_channel=HuifengGuan)
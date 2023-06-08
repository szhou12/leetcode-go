# [2537. Count the Number of Good Subarrays](https://leetcode.com/problems/count-the-number-of-good-subarrays/description/)

## Solution idea
### Sliding Window (Flex)
#### 思路总结
1. 题目要求 good subarray内需要 至少k对相同元素(不同位置，值相同)。顺着题目要求其实不容易想出如何利用滑窗。这里首先需要有**逆向思维**: 滑窗的物理意义表示为区间内只包含 < k对相同元素，这样的一个区间尽可能延伸。当滑窗停止延伸时，右边界到数组末尾的元素个数就是当前固定的左边界引领的good subarray的个数。
2. 确定了滑窗的思路，然后是怎么吃和怎么吐？也就是，怎么计算配对？
    * 使用一个map记录滑窗内每个元素的频次
    * 吃：如果右边界的元素在map中不存在，那么就更新map；如果存在，右边界元素可以跟滑窗内每一个相同元素配对，配对数增加量 = map中已有相同元素的频次
    * 吐：左边界元素与其他位置上相同元素的每一个配对都会消失，配对数减少量 = map中除去左边界元素后，剩余相同元素的频次
3. 代码中滑窗为左闭右开。更新结果需要分类讨论!

Time complexity = $O(n)$

## Resource
* Leetcode hint
* [【每日一题】LeetCode 2537. Count the Number of Good Subarrays](https://www.youtube.com/watch?v=ZlbgHJOevvU&ab_channel=HuifengGuan)
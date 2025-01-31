# [3113. Find the Number of Subarrays Where Boundary Elements Are Maximum](https://leetcode.com/problems/find-the-number-of-subarrays-where-boundary-elements-are-maximum/description/)

## Solution idea
### Stack - prevGreater + Binary Search + HashMap
1. 题目中 largest element in subarray 容易想到用单调栈以 $O(n)$ 时间找。但是问题是如何保证每个subarray的首位元素与这个最大元素相等？
2. 转换思路：考虑 给定一个元素 `nums[i]`，将其作为subarray的尾元素，数它前面有多少值相同的可以作为头元素。显然，为了保证尾元素的值是最大值，不能一个劲儿地往前找，何时停下？第一次遇到更大的值时就停下，那这不就是找 prevGreater ?
    - 这里要注意，本题是找 prevGreater，而不是 prevGreaterOrEqual。所以如果套用 [2104](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2104-Sum-of-Subarray-Ranges) 的模版，那么需要修改2104中 prevGreaterOrEqual 的写法: `<=` 改为 `<`，栈顶元素遇到严格更大的来者才出栈。因为2104中是为了去重，人为规定相同的值越靠左边“值越大”；而本题中是为了在遇到严格更大的值之前，数有多少个与尾元素相同的值排在它前面，显然不能把它们拒之门外。
3. 每一个尾元素都确定了prevGreater之后，如何数prevGreater到尾元素之间有多少个相同的值可以当作头元素？
    - 可以用一个hashmap，提前预处理：对于每一个 `nums[i]` 的值，用一个数组依次存入出现位置的index
    - 注意！目前为止，我们都记录的是index的信息！！
    - 发现，hashmap中每一个key对应的数组都是 递增序列！
    - 那么就可以用 binary search 的 upper_bound 从数组中找第一个出现的比 prevGreater所存index 严格大的index
    - 当前数组长度 - 这个index 就是可以当作头元素的个数

Time complexity = $O(n)$

## Resource
[每日一题】LeetCode 3113. Find the Number of Subarrays Where Boundary Elements Are Maximum](https://www.youtube.com/watch?v=FViZk7J_SHE&ab_channel=HuifengGuan)
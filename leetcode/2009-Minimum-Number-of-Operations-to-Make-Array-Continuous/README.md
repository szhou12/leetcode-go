# [2009. Minimum Number of Operations to Make Array Continuous](https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/description/)

## Solution idea

### Sliding Window - 窗口长度可变

* 预处理:
    1. `nums` 需要先去重
        * 为什么需要去重复? 因为每多一个重复元素，我们需要把它变大使得数组整体变得连续, 如果不连续的话，双指针的代码部分会把重复的元素都当成不需要变换的元素 (`min(N-(right-left))`)，这样显然是不对的
    2.  去重后，需要从小到大排列

* 左边界: 一定是某个 `arr[left]`
    * i.e. 对于最优解而言，连续区间的左端点一定是落在 `arr` 的某一个数值上

* 右边界延伸的条件: `arr[right]-arr[left]+1 <= len(nums)`
    * meaning: **数轴上**的连续区间 `[left, right]` 的跨度大于 `len(nums)`, 说明把 `nums` 所有元素都往里填充都填不满这个区间了, 就说明, `right` 跑的太远了, `left` 需要收缩了.
    * 也就是说，是否 `arr[right]-arr[left]+1 <= len(nums)`? 是的话，那么这个数值区间`[left, right]`一定能被 `nums`的元素 填满，`nums`也有可能还有多余的元素，我们就将它们继续拼接在区间的右边即可。这样，对于固定的 `left`，我们能找到最大的 `right`，使得这个区间尽量地大，那么需要从外面拿进来填充的`nums`元素自然就会尽量的少了。
    * e.g. `arr[left]=1, arr[right]=6`, 那么 **数轴上**的连续区间 `[left, right] = [1,2,3,4,5,6]`. 如果 `nums=[1,2,3,5,6]`, 显然, `right` 走的太远了 (多出了 用光 `nums` 所有元素 都填不满的空位), 因为 `nums`只需要把 6 变成 4 得到 `[1,2,3,4,5]`, 所以 **数轴上**的连续区间 `[left, right] = [1,2,3,4,5]` 就可以

* 需要变换的元素个数: `len(nums) - (right - left)`
    * input数组中, 实际被 **数轴上**的连续区间 `[left, right)` 囊括的元素有 `right-left` 个, 那么需要变换的元素 (连续区间以外的元素) 就是 `len(nums) - (right - left)` 这么多个

Time complexity = $O(n\log n + n) = O(n\log n)$

## Resource
[【每日一题】LeetCode 2009. Minimum Number of Operations to Make Array Continuous, 9/23/2021](https://www.youtube.com/watch?v=FiWjpCmgUwM&ab_channel=HuifengGuan)

[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Two_Pointers/2009.Minimum-Number-of-Operations-to-Make-Array-Continuous)
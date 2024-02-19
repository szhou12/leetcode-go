# [41. First Missing Positive](https://leetcode.com/problems/first-missing-positive/description/)

## Solution idea
1. 重要观察: for any array whose length is `n`, the smallest missing positive must be in range `[1,...,n+1]`, so we only have to care about those elements in this range and remove the rest.
    * 例如：`i=0`放1，`i=1`放2，`i=2`放3，`i=3`放4，...，`i=9`放10。
2. 实现上: 
    * STEP 1: '众神归位' - `i`位上对应的元素`nums[i]`，根据它的值，应该swap到`nums[i]-1`的index位上 (注意! 只有目标位置还没有被占时swap，如果已经有了，则跳过)
    * STEP 2: 再遍历一遍数组，找到第一个`nums[i] != i+1`的位置，返回`i+1`即可。找不到，最小的missing positive就是`n+1`。

Time complexity = $O(n)$

## Resource
[Golang concise solution with some explanation](https://leetcode.com/problems/first-missing-positive/solutions/17157/golang-concise-solution-with-some-explanation)
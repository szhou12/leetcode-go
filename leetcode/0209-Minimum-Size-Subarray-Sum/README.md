# [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)

## Solution idea
**Two Pointers**
`left`, `right` 指针都从 `index=0` 开始.

如果 subarray sum 小于 target, 就一直move `right`, 直到 sum $\geq$ target.

此时，停下`right`, 记录subarray length, 开始move `left`, 直到 sum $<$ target.

再重新开始 move `right`, 重复上述动作.

**双指针的物理意义**：`left`和`right`圈出来的subarray 永远表示 subarray ending at `right` whose sum $\geq$ target.

**实际编程中**：用for循环来控制右指针的思路比较清晰，左指针仅当固定了一个右指针的前提下进行移动；相对而言，用while来控制双指针比较容易出错.

Time complexity = $O(n)$

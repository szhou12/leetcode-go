# [2772. Apply Operations to Make All Array Elements Equal to Zero](https://leetcode.com/problems/apply-operations-to-make-all-array-elements-equal-to-zero/description/)

## Solution idea
### Diff Array (Greedy)
1. 理解题意: 逆向思维，使所有元素都减为0，反过来就是，从 `[0, 0, ..., 0]` 开始，通过若干次 "k-size subarray + 1" 操作变成 `nums[]`
2. 破题: 从左至右，每次给长度为k的subarray `nums[l:l+k-1]` (inclusive) 增加量为 `X`。相当于，差分数组的 `diff[l] + X` 和 `diff[l+k] + X`。e.g., `nums[0:0+k-1]`中每个元素增加`nums[0]`，相当于，`diff[l] + X`以及`diff[l+k] + X`
3. 规律:
    1. 对于第一个元素，我们想要的增加量 delta 使得 `0 -> nums[0]`。所以, `delta = nums[0] - 0`。整个 `nums[0:k-1]` (inclusive) 都要增加 delta，也就是 `diff[0] + delta` 和 `diff[k] + delta`。
    2. 对于第二个元素，此时第二个元素已经有一个基础量 `base = nums[0]`。如果这个基础量已经大于 `nums[1]`，那么我们无法通过只增不减的操作得到目标`nums[]`，直接返回 `false`。否则，我们还需要增加 `delta = nums[1] - base`，并且是整个 `nums[1:k]` (inclusive) 都要增加 delta，也就是 `diff[1] + delta` 和 `diff[k+1] + delta`。
    3. 对于第 i 个元素，已经有的基础量为 `base`。 如果 `base > nums[i]`，直接返回 `false`。否则，我们还需要增加 `delta = nums[i] - base`，并且是整个 `nums[i:i+k-1]` (inclusive) 都要增加 delta，也就是 `diff[i] + delta` 和 `diff[i+k] + delta`。
    4. 重复上述步骤，直到最后一个元素。如果最后一个元素刚好得到 `diff[n-1] + base == nums[n-1]`，那么返回 `true`，否则返回 `false`。

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 2772. Apply Operations to Make All Array Elements Equal to Zero](https://www.youtube.com/watch?v=7R7-4eIjWqM&ab_channel=HuifengGuan)

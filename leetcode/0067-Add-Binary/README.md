# [67. Add Binary](https://leetcode.com/problems/add-binary/)

## Solution idea

从后往前依次计算：`extra`表示后一位的贡献.

因为是binary，所以每一位的和只可能出现：0, 1, 2, 3

如果出现0, 1: 没有进位，`extra` 变成 0

如果出现2, 3: 有进位，`extra` 变成 1

Loop 完以后如果 `extra=1` 说明要在最前端增加 1.

Time complexity = $O(n)$
# [54. Spiral Matrix](https://leetcode.com/problems/spiral-matrix/)

## Solution idea

解法与[59. Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/)基本一致。

唯一区别是这道题所给的矩阵不一定是正方形。当为矩形的时候，有可能在里层填写时，先按照top方向从左至右填写了一遍，再按照bottom方向从右至左又填写了一遍一样的。(left/right 同理)

所以，为了避免重复填写，每一个方向填写时增加一个check条件：`num <= target`

如果文字说明不好理解，可以按照题目给的矩形例子走一遍，就会发现，代码填写里层时，先写了 6->7, 然后又来了 6

Time complexity = $O(m*n)$
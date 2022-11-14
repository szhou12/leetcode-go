# [836. Rectangle Overlap](https://leetcode.com/problems/rectangle-overlap/description/)

## Solution idea

* 1-D: Two intervals `(left1, right1)` and `(left2, right2)` overlapped 的充分必要条件: `left1 < right2 && left2 < right1`

* 2-D: Two rectangles overlapped 的充分必要条件 就是 横坐标，纵坐标 分别满足上述条件

Time complexity = $O(1)$


## Resource
[1-line Solution, 1D to 2D](https://leetcode.com/problems/rectangle-overlap/solutions/132340/c-java-python-1-line-solution-1d-to-2d/)
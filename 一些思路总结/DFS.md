# Depth First Search

## N-皇后

* 求所有解法 [51. N-Queens](https://leetcode.com/problems/n-queens/)

* 求解法数量 [52. N-Queens II](https://leetcode.com/problems/n-queens-ii/)

**思路**

```
# levels = # rows
# branches = # cols
```

一个可行解的表示方法：`array = {index: row index; value: column index to place Queen}`

不合法放置Queen的三种情况：

1. 同列: column index = previous column index

2. slope = 1

3. slope = -1

## 岛屿沉没类

* 四面八方 [417. Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/)
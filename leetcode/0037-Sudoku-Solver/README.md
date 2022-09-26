# [37. Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)

## Solution idea

### DFS

```
# of levels = # of empty cells
# of branches = 1~9
```

Time complexity = $O(9^n)$ where $n=$ # of empty cells

**如何找到任意一个格子 `[i, j]`对应的 3X3 box?**

对应 3X3 box 的起始(top-left)坐标为：`[(i/3)*3, (j/3)*3]`

以 `[4,6]` 为例：

* `4/3=1`: 把整个棋盘的行以 3行 为单位进行切割，4 属于第“一”行

* `(4/3)*3=3`: 代表“跳” 一个3: 0 -> 3

* `6/3=2`: 把整个棋盘的列以 3列 为单位进行切割，6 属于第“二”列

* `(6/3)*3=6`: 代表“跳” 两个3: 0 -> 3 -> 6

 * 实在不懂，可以对照附录图片再自己算一遍


## Resource

以下给出个人理解起来比较容易的Resources:

* DFS主体写法参考: [回溯算法最佳实践：解数独](https://labuladong.github.io/algo/4/31/109/)

* `isValid()`写法参考: [代码随想录：解数独](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0037.%E8%A7%A3%E6%95%B0%E7%8B%AC.md)


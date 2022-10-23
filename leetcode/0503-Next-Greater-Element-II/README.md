# [503. Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii/)

## Solution idea

### My Solution
* 既然又是 **Rotate/Circular array**, 那么就又可以用 **复制一倍**的方法
    * 即，`nums + nums`
    * e.g. `[1, 5, 7] --> [1, 5, 7, 1, 5, 7]`
* 增长一倍后, 我们只用看前 n 个元素; 且每个元素往后看 n-1 个元素
    * 即, `i-th` element 往后看 `[i+1, i+n)` 左闭右开, 因为 `(i+n)-th` element 是它自己

Time complexity = $O(n^2)$

### Better Solution
* 套用单调栈模版 - 从后往前看，栈内矮个子起开!
* 用 取模 (`%`) 的方法 **模拟** **复制一倍**的方法

Time complexity = $O(n)$

Similar question: [496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/)

## Resource
[单调栈结构解决三道算法题](https://labuladong.github.io/algo/2/23/63/)

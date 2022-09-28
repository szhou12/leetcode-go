# [98. Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/)

## Solution idea

### Recursion

**注意**传入 `min`, `max` 值，通过比较`root.Val`与 `min`, `max` 的值来判断是否合法。不能只比较`root.Val`与 `root.Left.Val` 或者 `root.Right.Val`, 这样是不够的！！！

Time complexity = O(height)

## Resource

[东哥带你刷二叉搜索树（基操篇）](https://labuladong.github.io/algo/2/21/43/)

[代码随想录](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0098.%E9%AA%8C%E8%AF%81%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91.md)
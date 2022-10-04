# [150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)

## Solution idea

### Stack

规则：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中

Time complexity = $O(n)$

## Resource
[代码随想录-150. 逆波兰表达式求值](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0150.%E9%80%86%E6%B3%A2%E5%85%B0%E8%A1%A8%E8%BE%BE%E5%BC%8F%E6%B1%82%E5%80%BC.md)
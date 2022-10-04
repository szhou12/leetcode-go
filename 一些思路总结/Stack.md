# Stack 的一些思路总结

* 由于栈结构的特殊性，非常适合做**对称匹配类**的题目

* 难点：在于如何针对题目设计入栈、弹栈的规则. ie. 啥东西入栈，啥时候入栈，啥时候出栈

## 经典题

* [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)
    * idea: 遇到左括号则入栈；遇到右括号，若栈顶是匹配的左括号则弹栈，否则，invalid. 若所有括号都check过，且此时栈内已空，则valid

* [150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)
    * idea：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中
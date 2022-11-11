# Stack 的一些思路总结

* 由于栈结构的特殊性，非常适合做**对称匹配类**的题目

* 难点：在于如何针对题目设计入栈、弹栈的规则. ie. 啥东西入栈，啥时候入栈，啥时候出栈

## 经典题

* [20. Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)
    * idea: 遇到左括号则入栈；遇到右括号，若栈顶是匹配的左括号则弹栈，否则，invalid. 若所有括号都check过，且此时栈内已空，则valid

* [150. Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/)
    * idea：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中

* 删除字符再比较: [844. Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/description/)
    * idea: 遇到字母则入栈；遇到#则取出栈顶元素

* 消消乐-简单版: [1047. Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/)

## 单调栈模板

**Idea**: 把数组的元素想象成并列站立的人，元素大小想象成人的身高。这些人面对你站成一列，如何求元素「2」的下一个更大元素呢？很简单，如果能够看到元素「2」，那么他后面可见的第一个人就是「2」的下一个更大元素，因为比「2」小的元素身高不够，都被「2」挡住了，第一个露出来的就是答案。

* [496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/)
    * 思路: 从后往前看，挤掉栈顶矮个子
    * 栈顶元素的物理意义: 维护栈顶元素, 使得它总是**大于**当前元素

* [503. Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii/)
    * 思路: 从后往前看，挤掉栈顶矮个子
    * 用 取模 (`%`) 的方法 **模拟** **复制一倍**的方法

* [739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)
    * 思路: 从后往前看，挤掉栈顶矮个子
    * 本题要算距离, 所以, **Stack 内存入元素的index, 而不是元素本身!!!**
# [1047. Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/)

## Solution idea

### Stack

* 用 Stack 解决 消消乐类型问题

* 机制：
    * loop 每个元素，如果栈空或者栈顶元素与 当前元素 不同，则当前元素入栈
    * 如果与栈顶元素相同，pop掉 栈顶元素
    * 注意：这个机制只能消掉 adjacent two elements, 如果adjacent elements多于2个就会出错

Time complexity = $O(n)$

* 拓展：如何消掉adjacent elements多于2个?
    * 栈内元素: 定义成一个struct{val: 元素本身, counter: 已push进栈的次数}
    * loop 每个元素:
        1. 如果栈空: push 当前元素{val, counter=1}
        2. 如果栈顶元素与当前元素相同: 栈顶元素counter+1
        3. 如果栈顶元素与当前元素不同: 
            1. 如果栈顶元素counter = 1, push当前元素{val, counter=1}
            2. 如果栈顶元素counter > 1, pop栈顶元素 直到: 1. 栈空; 2. 栈顶元素与当前元素相同; 3. 栈顶元素与当前元素不同 但 counter=1

## Resource
[代码随想录-1047. 删除字符串中的所有相邻重复项](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1047.%E5%88%A0%E9%99%A4%E5%AD%97%E7%AC%A6%E4%B8%B2%E4%B8%AD%E7%9A%84%E6%89%80%E6%9C%89%E7%9B%B8%E9%82%BB%E9%87%8D%E5%A4%8D%E9%A1%B9.md)

# Stack

* 由于栈结构的特殊性，非常适合做**对称匹配类**的题目

* 看到**输入带括号**的题，首先试试用 Stack 来解决

* 难点：在于如何针对题目设计入栈、弹栈的规则. ie. 啥东西入栈，啥时候入栈，啥时候出栈

## 目录
* [经典题](#经典题)
* [PrevSmaller, NextSmaller; PrevGreater, NextGreater](#prevsmaller-nextsmaller-prevgreater-nextgreater)

## 经典题

* :yellow_circle: 双stack实现queue: [232. Implement Queue using Stacks](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0232-Implement-Queue-using-Stacks)

* :yellow_circle: 双queue实现stack: [225. Implement Stack using Queues](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0225-Implement-Stack-using-Queues)

* [20. Valid Parentheses](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0020-Valid-Parentheses)
    * idea: 遇到左括号则入栈；遇到右括号，若栈顶是匹配的左括号则弹栈，否则，invalid. 若所有括号都check过，且此时栈内已空，则valid

* [150. Evaluate Reverse Polish Notation](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0150-Evaluate-Reverse-Polish-Notation)
    * idea：遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中

* 删除字符再比较: [844. Backspace String Compare](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0844-Backspace-String-Compare)
    * idea: 遇到字母则入栈；遇到#则取出栈顶元素

* 消消乐-简单版: [1047. Remove All Adjacent Duplicates In String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1047-Remove-All-Adjacent-Duplicates-In-String)

* 字符串解压缩: [394. Decode String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0394-Decode-String)


## PrevSmaller, NextSmaller; PrevGreater, NextGreater

**Idea**: 把数组的元素想象成并列站立的人，元素大小想象成人的身高。这些人面对你站成一列，如何求元素「2」的下一个更大元素呢？很简单，如果能够看到元素「2」，那么他后面可见的第一个人就是「2」的下一个更大元素，因为比「2」小的元素身高不够，都被「2」挡住了，第一个露出来的就是答案。

### 模版
* :red_circle: 所有subarray的极值的和: [2104. Sum of Subarray Ranges](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2104-Sum-of-Subarray-Ranges)
    * 单调栈: PrevSmaller, NextSmaller; prevGreater, nextGreater 都涉及到了

* :red_circle: 所有subarray的最小值的和: [907. Sum of Subarray Minimums](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0907-Sum-of-Subarray-Minimums)
    * **3-Pass + Stack (find next smaller element)**
    * 入手着眼点: 从 **分界点** 的物理意义 出发

* :yellow_circle: [1856. Maximum Subarray Min-Product]()
    * prevSmaller, nextSmaller + prefix sum

* [496. Next Greater Element I](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0496-Next-Greater-Element-I)
    * 思路: 从后往前看，挤掉栈顶矮个子
    * 栈顶元素的物理意义: 维护栈顶元素, 使得它总是**大于**当前元素

### 变体

* :red_circle: [3113. Find the Number of Subarrays Where Boundary Elements Are Maximum](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3113-Find-the-Number-of-Subarrays-Where-Boundary-Elements-Are-Maximum)
    * 横向拓展：新增花样 - 只考虑每个"中心"对象作为尾巴元素的情况，往前找值相同的头元素

* :red_circle: [2454. Next Greater Element IV](https://github.com/szhou12/leetcode-go/blob/main/leetcode/2454-Next-Greater-Element-IV/2454-Next-Greater-Element-IV.go)
    * 纵向拓展：一个stack不够用，就来两个

* [503. Next Greater Element II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0503-Next-Greater-Element-II)
    * 思路: 从后往前看，挤掉栈顶矮个子
    * 用 取模 (`%`) 的方法 **模拟** **复制一倍**的方法

* [739. Daily Temperatures](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0739-Daily-Temperatures)
    * 思路: 从后往前看，挤掉栈顶矮个子
    * 本题要算距离, 所以, **Stack 内存入元素的index, 而不是元素本身!!!**

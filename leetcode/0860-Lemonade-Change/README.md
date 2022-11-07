# [860. Lemonade Change](https://leetcode.com/problems/lemonade-change/description/)

## Solution idea

### Greedy + 模拟

* **解题关键**是分清楚情况，分别讨论
    * Case 1: 账单是5，直接收下
    * Case 2: 账单是10，消耗一个5，增加一个10
    * Case 3: 账单是20，优先消耗一个10和一个5，如果不够，再消耗三个5
* 账单是20的情况，为什么要优先消耗一个10和一个5呢？
    * 因为美元10只能给账单20找零，而美元5可以给账单10和账单20找零，美元5更万能！(本题**贪心**的地方)

Time complexity = $O(n)$

## Resource
[代码随想录-860.柠檬水找零](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0860.%E6%9F%A0%E6%AA%AC%E6%B0%B4%E6%89%BE%E9%9B%B6.md)
# [907. Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums/)

## Solution idea

### 3-Pass + Stack
1. 本题求所有可能subarray的最小值的和, 核心是怎么找到每个subarray的最小值. 最自然的想法是: 先找每个subarray再取min value. 但是复杂度太高 (确定一个起点, 确定一个终点需要$O(n^2)$, 题目array长度最长可以到$10^4$, 那么$n^2 = 10^8 > 10^6$。显然, 需要$O(n)$的算法来解).
2. **核心思路**: 反过来想, 先确定min value (每个元素轮流当作min value), 再看以该元素为中心能找到多少subarray. min value就相当于分界点, 然后以分界点出发向两边拓展边界看能拓展到哪里.
3. **Stack - Find Next Smaller Element**算法: 怎么向两边拓展边界? 从左至右扫一遍所有元素, 维护一个**Stack**, 作用是标记**每个元素右侧**第一个 $< $它的元素的index (next smaller element's index); 从右至左扫一遍所有元素, 维护一个**Stack**, 作用是标记**每个元素左侧**第一个$\leq $它的元素的index (previous smaller element's index) 
    * **Next Greater Stack的物理意义**: 如果 栈空 或者 当前candidate < 栈顶元素, 直接进栈; 否则, 栈顶元素出栈, 并把当前candidate的index标记为它的next greater, 重复这个动作直到栈顶元素不满足出栈条件

#### 要点总结

#### 物理意义

#### 代码整体结构总结


## Resource
[【每日一题】907. Sum of Subarray Minimums, 5/11/2021](https://www.youtube.com/watch?v=TZyBPy7iOAw&t=629s&ab_channel=HuifengGuan)
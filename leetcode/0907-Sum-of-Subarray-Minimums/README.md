# [907. Sum of Subarray Minimums](https://leetcode.com/problems/sum-of-subarray-minimums/)

## Solution idea

### 3-Pass + Stack

#### 要点总结
1. 本题求所有可能subarray的最小值的和, 核心是怎么找到每个subarray的最小值. 最自然的想法是: 先找每个subarray再取min value. 但是复杂度太高 (确定一个起点, 确定一个终点需要$O(n^2)$, 题目array长度最长可以到$10^4$, 那么$n^2 = 10^8 > 10^6$。显然, 需要$O(n)$的算法来解).
2. **核心思路**: 反过来想, 先确定min value (每个元素轮流当作min value), 再看以该元素为中心能找到多少subarray. min value就相当于分界点, 然后以分界点出发向两边拓展边界看能拓展到哪里.
3. **难点 1** - **Stack - Find Next Smaller Element**算法: 怎么向两边拓展边界? 从左至右扫一遍所有元素, 维护一个**Stack**, 作用是标记**每个元素右侧**第一个 < 它的元素的index (next smaller element's index); 从右至左扫一遍所有元素, 维护一个**Stack**, 作用是标记**每个元素左侧**第一个$\leq $它的元素的index (previous smaller element's index) 
    * **Next Smaller Stack的物理意义**: 如果 栈不空 或者 栈顶元素 > 当前candidate, 找到next smaller element, 把当前candidate的index标记为它的next smaller, 并重复这个动作直到栈顶元素不满足出栈条件; 否则, 直接进栈
    * **Previous Smaller Stack的物理意义**: 如果 栈不空 或者 栈顶元素 $\geq $ 当前candidate (这里为什么更严格, 需要小于等于?), 找到previous smaller element, 把当前candidate的index标记为它的previous smaller, 并重复这个动作直到栈顶元素不满足出栈条件; 否则, 直接进栈
4. **难点 2** - **去重**: 为什么从右至左找previous smaller element的时候, 出栈条件是 栈顶元素 $\geq $ 当前candidate, 这个 = 的作用是什么?
    * 这个 = 的作用是为了避免: 当array中有重复元素, 那相当于, 在轮流每个元素当min value的时候, 相同的min value会被反复取到, 从而产生重复的subarray
    ```
    // 拓展左右边界只有严格小时才停下
    2 [8 5 6 5 7 5 6] 2
         ^
    2 [8 5 6 5 7 5 6] 2
             ^
    // 结论: 当min value=5时, 不同位置的5会拓展到相同的最长subarray, 这是因为后一个位置的5把前一个位置的5也包含进去了, 导致了重复计算

    // 拓展右边界严格小时才停下, 拓展左边界小于等于时就停下
    2 [8 5 6 5 7 5 6] 2
         ^
    2 8 5 [6 5 7 5 6] 2
             ^
    // 结论: 不同位置的5不会拓展到相同的最长subarray, 这是因为后一个位置的5把前一个位置的5挡在门外了
    ```
    * *从右至左 出栈条件是 栈顶元素 $\geq $ 当前candidate* 相当于是添加了一个额外的人为规定: 认为如果有若干个相同的数，则最左边的那个才是最小值。
5. **难点 3** - **初始值定位找不到的情况**
    * 从左至右, min value找不到`next smaller element`时, 右边界拓展到哪里？答: 最后一个元素的下一个, i.e. n
    * 从右至左, min value找不到`previous smaller or equal element`时, 左边界拓展到哪里？答: 第一个元素的前一个, i.e. -1
6. **难点 4** - **如何计算最小值的和**?
    * 所有可能subarray的最小值的和 = 累加 (当前元素为min value的subarray个数 * min value) = 累加 ((左边界相对当前元素的长度 * 右边界相对当前元素的长度) * min value) = 累加 (`(i-a)*(b-i)` * min value) where a=左边界, b=右边界
    ```
    2 8 5 6 5 [7 5 5 5 5 6] 2
            a    i          b
    ```

#### 物理意义
*Note: `n=len(arr)`*

*Note: 为了容易理解，这里提及的 `[x:y]` 都是左闭右闭*

```
nextSmaller[i] := index of next element > ith element
初始值: nextSmaller[i] = n for all i

prevSmaller[i] := index of previous element >= ith element (必须用 >= 去重)
初始值: prevSmaller[i] = -1 for all i
```


#### 代码整体结构总结
```
Step 1: 构造 nextSmaller[]: 维护一个stack 从左至右 标记栈顶元素右侧第一个遇到的小于它的元素index
Step 2: 构造 prevSmaller[]: 维护一个stack 从右至左 标记栈顶元素左侧第一个遇到的小于等于它的元素index
Step 3: iterate 分界点 i, 累加 (i-prevSmaller[i])*(nexSmaller[i]-i)*arr[i]
```

Time complexity = $O(n)$

## Resource
[【每日一题】907. Sum of Subarray Minimums, 5/11/2021](https://www.youtube.com/watch?v=TZyBPy7iOAw&t=629s&ab_channel=HuifengGuan)
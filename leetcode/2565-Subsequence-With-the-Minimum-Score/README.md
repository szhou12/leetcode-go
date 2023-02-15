# [2565. Subsequence With the Minimum Score](https://leetcode.com/problems/subsequence-with-the-minimum-score/)

## Solution idea

### Binary Search 猜答案 + 3-Pass

*Note: 为了容易理解，这里提及的 `[x:y]` 都是左闭右闭*

#### 要点总结
1. 此题需要对题干进行翻译。翻译题目: 删除t的一个subarray (甭管内部怎么删的，删了几个char), 使得剩下的chars能组成s的subseq
2. 根据翻译, 可以归纳一个规律: 删除的这个subarray越长, 剩下的chars越容易组成s的subseq (eg. 最极端的例子是t删除到什么都不剩了，一个空串肯定是s的subseq) --> 单调性 --> 暗示可以**用 binary search 猜答案**
3. 圈定一个区间作为删除的subarray, 不需要在它内部额外操作, 直接整个删掉
4. 作图:
```
t:   x x x  {y y y}  z z z
         i           j
```
根据上图, t被圈出的subarray `{y y y}` 把t 分成L, M, R三段. 显然, 可以使用 `3-Pass` 来解

使用两个`DP`思想的arrays来分别记录L段和R段的信息: `left[]`, `right[]`

#### 物理意义
*Note: `m=len(s)`, `n=len(t)`*
```
left[i] := rightmost index r for the shortest prefix of s (i.e. s[0:r]) containing t[0:i]
right[j] := leftmost index l for the shortest suffix of s (ie. s[l:m-1]) containing t[j:n-1]
```

* 翻译一下
    * `left[]` 记录 s的前缀 中 右端点r 的index信息
    * `right[]` 记录 s的后缀 中 左端点l 的index信息

* 物理意义的几个**要点**:
    * `left[]`, `right[]` 记录的是(滑窗)端点的**index信息**, 不是长度或者其他什么物理量
    * What is **prefix of s** & **suffix of s**?
        * prefix of s: s的前缀 指的是 `s[0:r]`, 从开始index到r框住的区域, r不断往后延伸直到结尾index
        * suffix of s: s的后缀 指的是 `s[l:m-1]`, 从l到结尾index框住的区域, l不断向前延伸直到开始index
    * Why **shortest**?
        * `left[i]`, `right[j]` 寻找的是最小可行解 (下限), 最大可行解很容易想到就是s整个字符串; 同时, 也是**greedy/sliding window**的思想, 当计算下一个 `left[i+1]`的时候, r可以继承上一次的位置继续延伸，避免重复解或者漏解 (**shortest**相当于起到一个限制的作用, 每一轮只要刚刚好就收手, 进行下一轮)
    * 找不到的情况
        * 如果 `t[0:i]` 无法成为subseq (i.e. 包含了s没有的char), 则 `left[i] = MaxInt` (谨遵物理意义, `MaxInt`表示s前缀的右端点r不论怎么向后延伸, 即使伸到无限远处都无法包含`t[0:i]`)
        * 如果 `t[j:n-1]` 无法成为subseq (i.e. 包含了s没有的char), 则 `right[j] = MinInt` (谨遵物理意义, `MinInt`表示s后缀的左端点l不论怎么向前延伸, 即使伸到无限小处都无法包含`t[j:n-1]`)


## Resource
[【每日一题】LeetCode 2565. Subsequence With the Minimum Score](https://www.youtube.com/watch?v=vcjfoFhqzcI&ab_channel=HuifengGuan)
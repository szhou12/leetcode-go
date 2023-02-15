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
    left[] 存储s的前缀中r的index信息
    right[] 存储s的后缀中l的index信息

* What is **prefix of s** & **suffix of s**?

    prefix of s: s的前缀 指的是 `s[0:r]`, 从开始index到r框住的区域, r不断往后延伸直到结尾index
    suffix of s: s的后缀 指的是 `s[l:m-1]`, 从l到结尾index框住的区域, l不断向前延伸直到开始index



## Resource
[【每日一题】LeetCode 2565. Subsequence With the Minimum Score](https://www.youtube.com/watch?v=vcjfoFhqzcI&ab_channel=HuifengGuan)
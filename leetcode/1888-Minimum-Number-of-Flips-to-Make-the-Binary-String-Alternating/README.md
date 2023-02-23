# [1888. Minimum Number of Flips to Make the Binary String Alternating](https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/description/)

## Solution idea
### 3-Pass

*Note: For simpliciy, Type-1 operation called 'rotate'. Type-2 operation called 'flip'*

#### 要点总结
1. flip 和 rotate互相之间没有制约。 即，不会发生使用了一种操作导致另一种操作用不了了。即, x次 flips and y次 rotates 任意排列都能得到同一个结果
```
[1] [1] 1 0 [0] 0
r   r&f      f
不管对这三个元素先进行哪个操作rotate/flip, 都是一样的结果
```
2. **rotate 的作用**: 当L段和R段分别已经alternating(大前提)，但是，接触的地方非互异而**各自远心端互异**时，rotate可以使整个alternating
```
远心端            远心端
[1]010 ｜  010101[0]
     ------
     接触的地方
```

#### 物理意义

#### 代码整体结构总结


## Resource
[【每日一题】1888. Minimum Number of Flips to Make the Binary String Alternating, 6/9/2021](https://www.youtube.com/watch?v=shqRII8gvCo&ab_channel=HuifengGuan)
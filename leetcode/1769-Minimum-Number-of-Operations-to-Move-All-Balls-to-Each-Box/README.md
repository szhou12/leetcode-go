# [1769. Minimum Number of Operations to Move All Balls to Each Box](https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/)

## Solution idea

### Naive Solution

遍历分界点, 每个点都分别看一下左手边总共多少moves，右手边总共多少moves

Time complexity = $O(n^2)$


### 3-Pass

#### 要点总结
1. **优化思路**: 注意到遍历每一个分界点时，不需要重新把每个ball还原到它原来的位子再开始计算移动所用moves。直接继承移动ball到之前一个点的位置所用moves，再+1就行了 (`DP`思想)

#### 物理意义
*Note: `n=len(boxes)`*
*Note: 为了容易理解，这里提及的 `[x:y]` 都是左闭右闭*

```
leftMoves[i] := # of moves to move all balls in boxes[0:i-1] to cell i
Base case:
    leftMoves[0] = 0 // 最左边没有ball可以移动，显然也就没有moves
Recurrence:
    leftMoves[i] = leftMoves[i-1] + left[i]*1
```

```
left[i] := total # of balls in boxes[0:i-1] // i位置的左边有多少个ball, 有ball才需要移动, 没有就不用管
Base case:
    left[0] = 0
Recurrence:
    left[i] = left[i-1] + 1 if boxes[i-1] == 1
            = left[i-1]     0/w
// 显然, left[]是可以用一个标量left简化, 实现时不需要array left[]
```

```
rightMoves[i] := # of moves to move all balls in boxes[i+1:n-1] to cell i
Base case:
    rightMoves[n-1] = 0 // 最右边没有ball可以移动，显然也就没有moves
Recurrence:
    rightMoves[i] = rightMoves[i+1] + right[i]*1
```

```
right[i] := total # of balls in boxes[i+1:n-1] // i位置的右边有多少个ball, 有ball才需要移动, 没有就不用管
Base case:
    right[n-1] = 0
Recurrence:
    right[i] = right[i+1] + 1 if boxes[i+1] == 1
             = right[i+1]     0/w
// 显然, right[]是可以用一个标量right简化, 实现时不需要array right[]
```

Time complexity = $O(n)$


## Resource
[【每日一题】1769. Minimum Number of Operations to Move All Balls to Each Box, 2/24/2021](https://www.youtube.com/watch?v=XrOc63_GSsY&ab_channel=HuifengGuan)
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
3. Binary String alternating 可以有两种pattern: `01` or `10`
     * `01`: 以0开头，重复01
     * `10`: 以1开头，重复10



#### 物理意义
*Note: `n=len(s)`*

*Note: 为了容易理解，这里提及的 `[x:y]` 都是左闭右闭*

```
leftFlips01[i] := # of flips to make s[0:i] alternate with 01 pattern (leftmost element starts with 0, next with 1)
leftFlips10[i] := # of flips to make s[0:i] alternate with 10 pattern (leftmost element starts with 1, next with 0)
rightFlips01[i] := # of flips to make s[i:n-1] alternate with 01 pattern (rightmost element starts with 0, next with 1)
rightFlips10[i] := # of flips to make s[i:n-1] alternate with 10 pattern (rightmost element starts with 1, next with 0)
```

* **什么时候使用rotate?**: 对每一个分界点`i`, 根据L段和R段各自远心端互异原则, 分两种情况:
     1. `leftFlips01[i] + rightFlips10[i+1]`
     ```
     0101...   ...010101
           i  i+1
     ```
     2. `leftFlips10[i] + rightFlips01[i+1]`
     ```
     1010...   ...101010
           i  i+1
     ```

* 分界点`i`为了不越界, 有效范围为`[0:n-2]`. 需要单独处理顶到最右边的情况, i.e. `i=n-1`
     * 顶到最右边的情况 = 不发生rotate = L段`[0:n-1]`一整个alternating 01/10 (`leftFlips01[n-1]` or `leftFlips10[n-1]`)

#### 代码整体结构总结
```
Step 1: 构造 leftFlips01[] 和 leftFlips10[]: index i 从左往右, 依次计算需要flip的次数 (用i的奇偶性判断当前位子i应该是0还是1)
Step 2: 构造 rightFlips01[] 和 rightFlips10[]: index i 从右往左, 依次计算需要flip的次数 (用i的奇偶性判断当前位子i应该是0还是1)
     Note: i从右往左走时, 数值是从大到小, 为了方便, 转换i, 以i从右起是第几个index来判断奇偶性 (i.e. 用n-1-i的奇偶性判断当前位子i应该是0还是1)
Step 3: 初始值=min(leftFlips01[n-1], leftFlips10[n-1]); iterate 分界点, 取 min(leftFlips01[i] + rightFlips10[i+1], leftFlips10[i] + rightFlips01[i+1])
```

Time complexity = $O(n)$

## Resource
[【每日一题】1888. Minimum Number of Flips to Make the Binary String Alternating, 6/9/2021](https://www.youtube.com/watch?v=shqRII8gvCo&ab_channel=HuifengGuan)
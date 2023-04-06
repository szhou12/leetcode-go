# [940. Distinct Subsequences II](https://leetcode.com/problems/distinct-subsequences-ii/description/)

## Solution idea

### DP
#### 要点总结
1. `DP` 的定义与往常不同
    * 一般的, `DP[i]` 尝试与 `s[0:i]` 也就是与第i的位置产生联系, 但是在这道题下这么定义的话, 显然我们还需要辅助的数据结构来记录以`s[i]`结尾的 subsequennces 分别都是谁。
    * 那么, 多加一个维度呢? `DP[i][c]` 表示 `s[0:i]`中以character `c`结尾的 subsequennces的个数。好像行得通。但是似乎可以优化: 因为`DP[i][...]` 只与 `DP[i-1][...]` 产生联系，表示第i位置的字符加与不加在尾巴上产生的 subsequennces的个数。这种`DP[i]`只与`DP[i-1]`产生联系的关系, 我们可以降维优化, `i`这个维度可以直接用 `for i:=0; i < n; i++` 的迭代来替换掉 而不用存起来。
    * 这样, 我们就有了 `DP`的定义: `DP[c]` 表示 迭代到第i次时 (=`s[0:i]`)，以字符`c`为结尾的 subsequennces的个数. `DP`长度为26，代表26个字母。
    * `DP[c]`的定义还有效地避免了 同一个subsequence的重复计算: 因为`DP[c]`的定义就是唯一以`c`为结尾的subsequennces的个数，`c`不同，自然也就不会有重复计算 (i.e. 在计算状态转移方程 (Recurrence) 时, 同时达到了去重的目的)
2. 在计算状态转移方程 (Recurrence) 时, `DP[]`的定义/物理意义 实际上发生了变化, 但是由于两种定义/物理意义 在数学上可证明为等价的 (证明在下面给出), 状态转移方程的计算仍可保证正确性.

#### DP整体结构
1. Definition
* `DP[c]` = at i-th iteration (i.e. `s[0:i]`), the number of distinct subsequences ending with character `c`
* `DP[s[i]]` = at i-th iteration (i.e. `s[0:i]`), the number of distinct subsequences ending with `s[i]`

2. Base Case
* Init `DP[c]`= 0

3. Reccurrence
* For n times, `DP[s[i]]` = $\sum_{c=0}^{25} $ `DP[c]` + 1
    * +1: `s[i]`单独作为一个subsequence，不加在任何其他character后面

4. 为什么可以避免重复? 即, $abc^1d$和$abc^2d$如何做区别？
* 注意到, Recurrence中每次迭代, `s[i]`对应的character结尾的subseq都被重新计算赋值; 又有, `DP[s[i]]` = `DP[c]`
* 可以理解为，每次迭代代表一个时刻t, 如果两个时刻t1和t2都有相同字母d结尾, 后一时刻t2的d的值会覆盖前一时刻t1的d的值
```
[a  b  c1  X X X c2] d
 0      ...      i-1 i
```
* $abc^1d$和$abc^2d$产生的情况如上图, 根据`DP[c]`的定义，以c为结尾的subseq个数在i-1轮迭代时只会被唯一计算，所以在第i轮迭代把d加在c屁股后面不会有$abc^1d$和$abc^2d$两种情况, 只会有 abcd

#### 数学证明

**Proof:** `DP[s[i]]` = `DP[c]` (# of distinct subseqs ending with `s[i]` = # of distinct subseqs ending with `c`)

($\Rightarrow $) **WTS** `DP[s[i]]` $\subseteq $ `DP[c]`
* 因为`s[i]`一定是26个字母中某一个`c`，所以`DP[s[i]]`一定属于`DP[c]`

($\Leftarrow $) **WTS** `DP[s[i]]` $\supseteq $ `DP[c]`
* 因为根据定义, `DP[c]`描述的是以`c`为结尾的subseq，我们并不关心`c`到底是在`s[0:i]`中具体哪个位置，那么，我们就可以认为`c`就是`s[i]`，所以，每个符合条件的`DP[c]`都可以对应到`DP[s[i]]`

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 940. Distinct Subsequences II](https://www.youtube.com/watch?v=2Io_meiaqng&ab_channel=HuifengGuan)
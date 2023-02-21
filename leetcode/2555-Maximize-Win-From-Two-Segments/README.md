# [2555. Maximize Win From Two Segments](https://leetcode.com/problems/maximize-win-from-two-segments/description/)

## Solution idea

### 3-Pass

#### 要点总结
1. 题眼: 2 intervals ==> 找中间的分界点
    * 提及对两个intervals操作的问题, 一个自然的思路是找中间的分界点
2. 遍历某一个点 (分界点): 看它左边可以包含最多prizes的segment在哪里，看它右边可以包含最多prizes的segment在哪里
3. 因为求max，所以是越多越好，尽量不重合就会越多; 除非整个数组长度容不下 2k, 才会取成重合的, 如果必然重合, 显然所有的prizes都能取到 (edge case)
4. 作图:
```
p: ------------------------
            mid
     ----          ----
     i   r         l   j
```
根据上图, 分界点`mid` 把 p 分成L, R三段. 显然, 可以使用 `3-Pass` 来解

使用两个`sliding window`思想的arrays来分别记录L段和R段的信息: `pre[]`, `post[]`

#### 物理意义
*Note: `n=len(p)`*

*Note: 为了容易理解，这里提及的 `[x:y]` 都是左闭右闭*
```
pre[r] := max # prizes that can be collected from k-len segment in p[0:r]
    Note: k-len segment that gives max # not necessarily ending at index r, can be any k-len seg in interval [0:r]
post[l] := max # prizes that can be collected from k-len segment in p[l:n-1]
    Note: k-len segment that gives max # not necessarily starting at index l, can be any k-len seg in interval [l:n-1]
```

* 构建`pre[]`, `post[]`
    * sliding window flex 版本: 以 `pre[]` 为例, 右端点`r`不停往前走, 当`p[r]-p[l] > k`时, 前进左端点
    * 注意: sliding window在刚开始滑的时候, 会存入 < k-len segment的信息, 但是这不要紧, 因为此题的input array is non-decreasing order, 所以在之后iterate 分界点并取最大值的时候, 一定会被淘汰掉 (i.e. < k=len segment 给的结果一定比 k-len segment给的结果小).

#### 代码整体结构总结
```
Step 1: 构造pre[]: 右端点不停往右走, length of segment > k 时, 左端点前进
Step 2: 构造post[]: 左端点不停往左走, length of segment > k 时, 右端点前进
Step 3: iterate 分界点, 取 max(pre[mid]+post[mid+1])
```

Time complexity = $O(n)+O(n)+O(n) = O(n)$

## Resource

[【每日一题】LeetCode 2555. Maximize Win From Two Segments](https://www.youtube.com/watch?v=0Tjuy464sP8&ab_channel=HuifengGuan)
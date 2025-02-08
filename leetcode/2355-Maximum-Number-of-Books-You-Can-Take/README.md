# [2355. Maximum Number of Books You Can Take](https://leetcode.ca/2022-09-04-2355-Maximum-Number-of-Books-You-Can-Take/)

## Solution idea
### DP + Stack
1. 见到Subarray的题，首先想想有哪些常用的算法？能想到的有: DP, Stack, Prefix Sum, Sliding Window, Segment Tree
    - DP: 能想到的是定义 `dp[i] := ` max subarray sum ending at i.
    - Stack: 能想到的是用单调栈 given `nums[i]`，找 prevSmaller, nextSmaller etc.
2. observation: 我们留意到一个性质 - 看`dp[i]`时，势必我们会把`books[i]`都取满。
3. 接下来的问题是，基于这个observation，我们怎么向前取，得到最大的subarray？我们考虑下面的例子：
    ```
                i
    ... 3 10 10 9
    ... 3][7 8] 9
    ```
    站在`i`处，如果前面的值更大，那我们没法取满，只能逐次递减地取，从而形成等差数列；如果前面的值更小，陡减，没法形成一个等差数列，这时候，我们可以取满，而取满就意味着我们向前遇到了`dp[j]`。这样，我们就发现，`dp[i]`代表的subarray可以大概分成两部分: `dp[j] + sum of arithmetic seq`
4. 好，发现这个pattern以后，我们就可以考虑怎么确定这个分界点`j`。发现没有，找这个`j`的过程，其实很像用单调递增栈找prevSmaller。只是，我们不能单纯地套用模版。因为，我们不能单纯地用`nums[i]`去找prevSmaller。比如下面这个例子:
    ```
                i
    2, 9, 4, 7, 5, 2
          ^
    1  2  3  4  5
    ```
    当在5这里时，它的prevSmaller是4，但是4并不是分界点`j`，我们没法取满4。所以，我们要额外留意，什么时候算是“前面的值更小”？
5. 答案是预留一段等差数列的空间之后，还小。即: 我们不用`nums[i]`本身去比较，而是做一个变换，使用`nums[i] - i`。

Time complexity = $O(n)$

## Resource
[【LeetCode】2355. Maximum Number of Books You Can Take](https://www.bilibili.com/video/BV1PG411n7FS/?spm_id_from=333.337.search-card.all.click&vd_source=0c02ef6f6e7a2b0959d7dd28e9e49da4)
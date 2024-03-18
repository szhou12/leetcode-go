# [2938. Separate Black and White Balls](https://leetcode.com/problems/separate-black-and-white-balls/description/)

## Solution idea
### Greedy
1. 理解题意：题目要求每次swap需要发生在相邻的两个位置上，且只有`10`需要变换成`01`。
2. 根据题意，可以总结出规律：任意一个位子上出现1，若要把它交换到右侧以至于它的右边全是1，那么与它右边的0交换的最少次数就是它右边的0的出现次数。想象成“挤泡泡” (泡泡=0)，要把1挤到右边，就要挤掉右手边所有的0。
3. 实现：从右往左遍历，每遇到一次0，累计0的个数；每遇到一次1，就把当前累计的0的个数加到结果里。

Time complexity = $O(n)$

## Recap: Rainbow Sort
* 0, 1代表两种颜色，如何sort使得相同颜色聚在一起？
    * 双指针相向而行
    * `[0 : l) := 0`, `(r : n-1] := 1`
    ```
    if l == 0 && r == 1 {
        l++, r--
    } else if l == 0 && r == 0 {
        l++
    } else if l == 1 && r == 1 {
        r--
    } else if l == 1 && r == 0 {
        swap(l, r)
        l++, r--
    }
    ```

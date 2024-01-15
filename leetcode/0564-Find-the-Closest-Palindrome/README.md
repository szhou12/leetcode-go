# [564. Find the Closest Palindrome](https://leetcode.com/problems/find-the-closest-palindrome/description/)

## Solution idea
### Math
1. String上进行十进制的加减操作：
    * 进位：加法运算时，从最低位开始往高位loop，如果有进位 (当前位+1 > 9)，记 carry = 1，进入下一循环；如果没有进位 (当前位+1 <= 9)，改 carry = 0，进入下一循环。loop 结束时，检查 corner case: 如果loop结束后，carry = 1，说明最高位也有进位，是这样的情况：`999` + 1 = `1000`，需要额外生成一个长度为 n+1 的字符串，首位 = 1，其余位 = 0。
    * 借位：减法运算时，从最低位开始往高位loop，如果有借位 (当前位-1 < 0)，记 carry = 1，进入下一循环；如果没有借位 (当前位-1 >= 0)，改 carry = 0，进入下一循环。loop 结束时，检查 corner case: 如果loop结束后，最高位显示0 并且 原先String长度 > 1，是这样的情况：`100` - 1 = `099`，存在leading 0，需要额外生成一个长度为 n-1 的字符串，所有位 = 9。注意：如果原先String长度 = 1，是这样的情况：`1` - 1 = `0`，可以保留leading 0，不需要额外生成字符串。 
2. String上把数字变成Palindrome：
    * 从两端向终点，镜像复制高位数字给低位数字，直到中间位置。
3. Closest Palindrome：
    * 如何找到最接近目标数字的回文串？
    * 高位尽量不动，调整低位。
    * 从中间位置为分界点，把数字切成两半：左半高位包含中点部分 first-half，右半低位部分 second-half。
        * 小技巧：中间位置的 index = `(n-1)/2` (奇数长正好在中间，偶数长在中间偏左)。
            * e.g. `[13]2`, `[12]31`
    * 最接近且比目标数字稍大的回文串 = first-half + 1 并把 first-half 镜像复制到 second-half。
        * e.g. `[13]2` -> 13+1 = 14 -> 1(4)~(4)1 -> 141 -> `[14]1`
    * 最接近且比目标数字稍小的回文串 = first-half - 1 并把 first-half 镜像复制到 second-half。
        * e.g. `[19]71` -> 19-1 = 18 -> 18~81 -> 1881 -> `[18]81`

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 564. Find the Closest Palindrome](https://www.youtube.com/watch?v=IdhV2dcvvSw&ab_channel=HuifengGuan)
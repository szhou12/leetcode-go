# [1221. Split a String in Balanced Strings](https://leetcode.com/problems/split-a-string-in-balanced-strings/)

## Solution idea

### Greedy Algorithm
* 从前向后遍历，只要遇到平衡子串，计数就+1，遍历一遍即可。
    * 如何判断平衡子串: 遇到 L, counter++; 遇到 R, counter--; 每当counter==0就是找到一个平衡子串
* 局部最优：从前向后遍历，只要遇到平衡子串 就统计
* 全局最优：统计了最多的平衡子串。
* 局部最优可以推出全局最优，举不出反例，那么就试试贪心。

Time complexity = $O(n)$

## Resource
[代码随想录-1221. 分割平衡字符串](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1221.%E5%88%86%E5%89%B2%E5%B9%B3%E8%A1%A1%E5%AD%97%E7%AC%A6%E4%B8%B2.md)
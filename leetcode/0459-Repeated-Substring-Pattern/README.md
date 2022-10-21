# [459. Repeated Substring Pattern](https://leetcode.com/problems/repeated-substring-pattern/)

## Solution idea

### My Solution
Naive Approach: 所有的以第一个字母为开始的substring挨个试一遍, 找到一个就返回true

Time complexity = $O(n^2)$

### Better Solution: 移动拼接
* `s+s`
* 如果 `s` 有重复的 `substr`, 那么既然前面有相同的子串，后面有相同的子串，用 `s + s`，这样组成的字符串中，后面的子串做前串，前面的子串做后串，就一定还能组成一个 `s`
* 所以判断字符串`s`是否有重复子串组成，只要两个`s`拼接在一起，里面还出现一个`s`的话，就说明是由重复子串组成。
* 注意！我们在判断 `s + s` 拼接的字符串里是否出现一个`s`的的时候，要刨除 `s + s` 的首字符和尾字符，这样避免在`s + s`中搜索出原来的`s`，我们要搜索的是中间拼接出来的`s`。

### Optimal Solution: KMP
* 在一个String中查找是否出现过另一个String，这是KMP的看家本领!

## Resource
移动匹配解法 & KMP解法: [代码随想录-459.重复的子字符串](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0459.%E9%87%8D%E5%A4%8D%E7%9A%84%E5%AD%90%E5%AD%97%E7%AC%A6%E4%B8%B2.md)
# [1002. Find Common Characters](https://leetcode.com/problems/find-common-characters/description/)

## Solution idea

### HashMap
* Step 1: 用一个 map 记录每个出现在第一个word里的字母分别在每个word中出现的次数
    * `map = [key=char, value = [counts0, counts1, ...]]`
    * 注意：这里以 `words[0]` 中出现的字母为基准，`words[0]`中没有出现过的字母在之后步骤中不做考虑
* Step 2: 每个字母对应的 count数组中 取最小值，就是这个字母在所有单词的common个数

Time complexity = $O(m) + O(n*m) + O(m*(n+m)) = O(m*(n+m))$ where $m=$ average length of a word, $n=$ length of input slice `words` (ie. total number of words)

## Resource
[代码随想录-1002. 查找常用字符](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1002.%E6%9F%A5%E6%89%BE%E5%B8%B8%E7%94%A8%E5%AD%97%E7%AC%A6.md)
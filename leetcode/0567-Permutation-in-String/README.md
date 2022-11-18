# [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/description/)

## Solution idea

### Sliding Window
* 整体code框架与 [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/description/) 完全一致
* 不同的地方在于: 此题找到一个合法的就可以返回，不用都找到
* 因为 sliding window 是定长, 写法也可以是单指针

Time complexity = $O(n)$

## Resource
[滑动窗口算法](https://labuladong.github.io/algo/2/20/27/)
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Hash/438.Find-All-Anagrams-in-a-String)
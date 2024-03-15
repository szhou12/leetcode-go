# [3045. Count Prefix and Suffix Pairs II](https://leetcode.com/problems/count-prefix-and-suffix-pairs-ii/description/)

## Solution idea
### KMP + Trie
1. Recap:
    * Trie适应症：在众多str中查找是否存在某一str
    * KMP适应症：在str中找某一substr 首次出现或者所有出现的位置。lsp具有传递性transitive。
Time complexity = $O(\sum{i} L_i)$ where $L_i$ is the length of `words[i]`

## Resource
[【每日一题】LeetCode 3045. Count Prefix and Suffix Pairs II](https://www.youtube.com/watch?v=a7zikGuzDjM&ab_channel=HuifengGuan)
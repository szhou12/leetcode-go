# [3045. Count Prefix and Suffix Pairs II](https://leetcode.com/problems/count-prefix-and-suffix-pairs-ii/description/)

## Solution idea
### KMP + Trie
1. Recap:
    * Trie适应症：在众多str中查找是否存在某一str
    * KMP适应症：a. 在str中找某一substr 首次出现或者所有出现的位置。b. lsp具有传递性transitive。
2. Loop backwards
    * 题目要求一个合法的pair`(i, j)` s.t. `i < j`. 因此，从后往前遍历每个word。对于当前word: `words[i]`，能与它匹配的`j`已经存在一个数据结构里，因此可以`O(n)`时间找到所有匹配的`j`。
3. KMP Step 1
    * 题目要求一个匹配需要同时match prefix和suffix。直接明示需要使用KMP。对于每一个word，使用KMP以`O(n)`时间找到最长、次长、第三长，... 的longest suffix-prefix (lsp)，将每一个符合lsp的前缀长度记录在一个数据结构里 (剧透：Trie)。将来每来一个新word，就在这个数据结构里找匹配的`j`的个数。
4. Trie
    * 使用Trie来储存每次看到的word，并用一个计数器(Why counter? 因为方便数出现的`j`的个数)在对应的节点标记出现符合lsp的前缀的结尾字符。
    * Trie的添加和查找操作都是`O(n)`时间。
    * 不使用Hashing的原因：容易发生collision。

Time complexity = $O(\sum{i} L_i)$ where $L_i$ is the length of `words[i]`

## Resource
[【每日一题】LeetCode 3045. Count Prefix and Suffix Pairs II](https://www.youtube.com/watch?v=a7zikGuzDjM&ab_channel=HuifengGuan)
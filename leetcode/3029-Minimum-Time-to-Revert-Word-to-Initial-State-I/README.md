# [3029. Minimum Time to Revert Word to Initial State I](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-i/description/)

## Solution idea
1. 破题：切入点在最后保留的尾巴 (后缀)。尾巴后面接续的部分不重要，因为可以放置任意字符。简化题意以后，就可以看出，这道题是求可以保留的最长后缀。在最终成型的字符串中，这个保留下来的尾巴会被顶到前面成为头部 (前缀)。所以，解题思路就是，找到可以保留的最长后缀，使得它等于前缀。
2. 3029 的数据量级并不大，所以可以使用暴力解法：直接比较在最少切几刀的情况下，使得后缀 = 前缀.
3. 实现细节：注意找准index来定为前缀和后缀。
    * 假设：切`t`刀时，找到后缀=前缀
    * 切掉的长度 = `tk`，保留的尾巴长度 = `n-tk`
    * 用index表示：切掉的长度 = `word[0 : tk-1]`，保留的尾巴长度 = `word[tk : n-1]`，需要相等的前缀 = `word[0 : n-tk-1]` (inclusive)
    * 所以判断条件：`word[0 : n-tk-1] == word[tk : n-1]`
4. 一图胜千言：
![autodraw 2_7_2024](https://github.com/szhou12/leetcode-go/assets/35708194/d2ec6b8c-273f-4c68-b0ce-9cd81dc2a15f)

Time complexity = $O(n)$


## Resource
[【每日一题】LeetCode 3031. Minimum Time to Revert Word to Initial State II](https://www.youtube.com/watch?v=ySjzFSCqLBI&ab_channel=HuifengGuan) - First 7.5 min

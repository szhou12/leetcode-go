# [3031. Minimum Time to Revert Word to Initial State II](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/description/)

## Solution idea
### KMP
1. 思路接续 [3029](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3029-Minimum-Time-to-Revert-Word-to-Initial-State-I)。但是因为量级增大，之前的解法不再适用。
2. 求最长后缀=前缀，使用 KMP 算法解决。
3. 使用 KMP STEP 1 可以求出**最长**的后缀=前缀 `lsp[(n-1)-j+1 : n-1]` (尾巴的长度=`j`)。**但是注意！**这里有一个问题：切`t`刀`k`长度的前缀，不一定刚好可以完整切掉头部，保留下最长的尾巴`j`。有可能切"多"了，剩下的尾巴长度不足`j`。此时，我们就需要寻找：第二长的后缀=前缀，第三长的后缀=前缀...
4. 如何找第二长，第三长？这里考察到 KMP算法中很重要的transitive思想!
    * 首先，我们看到`word[0:j-1]`这一段，其中，`lsp[j-1]`代表最长的长度`j'` s.t. `word[0:j'-1] == word[(j-1)-j'+1:j-1]`
    * 根据传递性(transitive)，我们可以得到 `word[0:j'-1] == word[(n-1)-j'+1 : n-1]`
    * 这个长度`j'`就是我们寻找的第二长度的后缀=前缀。
    * 以此类推，第三长度在`word[0:j'-1]`这一段中找...
    * 不停缩短，直到长度为0，说明我们找不到后缀=前缀，整个字符串都会被切掉。

### 一图胜千言

Time complexity = $O(n)$


## Resource
[【每日一题】LeetCode 3031. Minimum Time to Revert Word to Initial State II](https://www.youtube.com/watch?v=ySjzFSCqLBI&ab_channel=HuifengGuan)
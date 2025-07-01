# [472. Concatenated Words](https://leetcode.com/problems/concatenated-words/description/)

## Solution idea
### DP
Time complexity = $O(m^3 * n)$ where $m$ is the length of the longest word and $n$ is the number of words. Although we use HashSet, we need to consider the cost to calculate the hash value of a string internally which would be $O(m)$. So putting all words into the HashSet takes $O(n * m)$. For each word, the i and j loops take $O(m^2)$. The internal logic to take the substring and search in the HashSet needs to calculate the hash value for the substring too, and it should take another $O(m)$.

Space complexity = $O(m * n)$. This is just the space to save all words in the dictionary, if we don't take M as a constant.

### Trie + DFS (Memoization)

## Resource
[【每日一题】472. Concatenated Words, 4/27/2021](https://www.youtube.com/watch?v=dsnTJscs4BA&ab_channel=HuifengGuan)
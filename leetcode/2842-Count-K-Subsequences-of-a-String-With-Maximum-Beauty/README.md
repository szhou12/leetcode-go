# [2842. Count K-Subsequences of a String With Maximum Beauty](https://leetcode.com/problems/count-k-subsequences-of-a-string-with-maximum-beauty/description/)

## Solution idea
### DFS

#### 优化解法
* 用一个 array `counter[]` 记录每一个字母在输入字符串 `s` 中出现的次数 (`counter[]` 最长只有26)
* 将 `counter[]` 从大到小排序，前k个的和 就是能得到的最大beauty
* 再在排好序的 `counter[]` 上进行 DFS。每一次选择加入当前字母时，当前subsequence组合数就乘上加入字母对应的次数 (因为是从有次数个相同字母中选择一个)
* Time complexity = $O(n) + O(26) + O(26\log 26) + O(k) + O(26^k)$

#### 我的解法 (会超时:exclamation:)
* 不做任何变换，直接在输入字符串 `s` 上进行 DFS 穷举出所有符合要求的子序列
* Time complexity = $O(n^k)$
```go
var M = int(1e9+7)

func countKSubsequencesWithMaxBeauty(s string, k int) int {
    f := make([]int, 26)
    for i := 0; i < len(s); i++ {
        char := s[i] - 'a'
        f[char]++
    }

    visited := make([]bool, 26)
    beautyMap := make(map[int]int) // key=beauty value; value=count

    cur := ""
    dfs(s, k, cur, 0, visited, f, beautyMap)

    
    maxBeauty := 0
    for k, _ := range beautyMap {
        maxBeauty = max(maxBeauty, k)
    }
    return beautyMap[maxBeauty]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func dfs(s string, k int, cur string, curIndex int, visited []bool, f []int, beautyMap map[int]int) {
    if len(cur) == k {
        beauty := 0
        for _, char := range cur {
            beauty += f[int(char -'a')]
        }
        beautyMap[beauty]++
        beautyMap[beauty] %= M
        return
    }

    if curIndex >= len(s) {
        return
    }

    for i := curIndex; i < len(s); i++ {
        if visited[s[i] - 'a'] {
            continue
        }
        visited[s[i] - 'a'] = true
        dfs(s, k, cur+s[i:i+1], i+1, visited, f, beautyMap)
        visited[s[i] - 'a'] = false
    }
}
```


## Resource
[【每日一题】LeetCode 2842. Count K-Subsequences of a String With Maximum Beauty](https://www.youtube.com/watch?v=8pPPcODAWAA)
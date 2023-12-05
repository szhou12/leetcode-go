# [2842. Count K-Subsequences of a String With Maximum Beauty](https://leetcode.com/problems/count-k-subsequences-of-a-string-with-maximum-beauty/description/)

## Solution idea
### DFS

#### 我的解法 (会超时:exclamation:)
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

#### 优化解法

## Resource
[【每日一题】LeetCode 2842. Count K-Subsequences of a String With Maximum Beauty](https://www.youtube.com/watch?v=8pPPcODAWAA)
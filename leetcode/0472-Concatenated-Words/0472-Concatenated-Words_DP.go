package leetcode

// Solution 1: DP
// dp[i] := boolean whether word[0:i-1] (inclusive) can be creatd by concatenation. 
// ie. there exists j < i s.t. dp[j]=true and word[j:i] exists in the map.
// dp[0] = true
// dp[i] = true if dp[j] && dict[word[j:i]]=true for 0 < j < i
func findAllConcatenatedWordsInADict_dp(words []string) []string {
    dict := make(map[string]bool)
    for _, w := range words {
        dict[w] = true
    }

    res := make([]string, 0)
    for _, word := range words {
        // corner case
        if word == "" {
            continue
        }

        // mute current word itself
        dict[word] = false
        if canBreak(word, dict) {
            res = append(res, word)
        }
        // unmute current word for other words to check
        dict[word] = true

    }

    return res
}

func canBreak(word string, dict map[string]bool) bool {
    n := len(word)

    dp := make([]bool, n+1)
    

    // base case
    dp[0] = true // empty string can always be created by not selecting any words from dict

    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            if dp[j] && dict[word[j:i]] {
                dp[i] = true
                break
            }
        }
    }

    return dp[n]
}

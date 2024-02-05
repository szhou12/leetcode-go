# [2916. Subarrays Distinct Element Sum of Squares II](https://leetcode.com/problems/subarrays-distinct-element-sum-of-squares-ii/description/)

## Solution idea
### DP + Segment Tree
Let `DP[i] =` sum of squares of distinct counts of an subarray in `nums[0:i]` (inclusive) ending at `i`.

Define `count[a:b] =` count of distinct elements in `nums[a:b]` (inclusive); `square[a:b] = count[a:b]^2`

Assume `nums[k] = nums[i]` (也就是，`i`位出现的元素上一次出现是在`k`位).
```
DP[i] = sum{square[j:i]} for j = 0, 1, ..., i
      = sum{square[j:i]} for j = 0, 1, ..., k + sum{square[j:i]} for j = k+1, k+2, ..., i
            (i不贡献1)
      = sum{square[j:i-1]} for j = 0, 1, ..., k + sum{square[j:i]} for j = k+1, k+2, ..., i
            (1)                                       (2)
      = DP[i-1]             +        2*sum{count[j:i-1]} for j = k+1, ..., i-1    +    (i-1-k-1+1)   +     1
        (1)+sum{square[j:i-1]}
      = DP[i-1]        +        2*sum{count[j:i-1]} for j = k+1, ..., i-1        +        (i-k-1)        +        1
                                    segment tree
```

```
(2) = sum{(count[j:i-1] + 1)^2} for j = k+1, ..., i-1   +     1
      表示i对[j:i-1]贡献1个count                            表示i自己作为单独的subarray，与[j:i-1]无关
    = sum{count[j:i-1]^2 + 2*count[j:i-1] + 1} for j = k+1, ..., i-1   +     1
    = sum{square[j:i-1] + 2*count[j:i-1] + 1} for j = k+1, ..., i-1   +     1
    = sum{square[j:i-1]} for j = k+1, ..., i-1 + sum{2*count[j:i-1] + 1} for j = k+1, ..., i-1   +     1
```


## Resource
[【每日一题】LeetCode 2916. Subarrays Distinct Element Sum of Squares II](https://www.youtube.com/watch?v=Ga19fyvyyKk&ab_channel=HuifengGuan)
# [3434. Maximum Frequency After Subarray Operation](https://leetcode.com/problems/maximum-frequency-after-subarray-operation/description/)

## Solution idea
1. Key observation: for the subarray to be selected, we need to have 2 observations: 1. we will convert the mode value to k in order to result in most k counts; 2. this mode value is not k, this means after trasnformation, any k in the subarray will no longer be k.
2. Deriving from this observation, suppose we have selected a subarray, we look at its `net gain = the largest frequency (the mode) - the number of k's`. Our goal is to find the positive net gain and maximize it. However, if a subarray gives negative net gain, it means that converting this subarray makes us lose more k's than gain, it's not worth it. Therefore, we should throw away this subarray. 
3. We do this by keeping track of the leftendpoint (first occurrence) of each element in `nums`, and everytime we find another occurrence as rightendpoint, we check the net gain using this element as the mode minus k counts between leftendpoint and rightendpoint. If the net gain is not positive, we throw away by moving this element's leftendpoint to the current rightendpoint and reset its frequency to 1.
4. in the end, the max k counts = `prefixSum[n-1] + maxGain`
```
[ .. +a .. [.. X - b ..] ] nums
[ .. +a .. [... + b ...] ] prefixSum

a: k's frequency outside the selected subarray
b: k's frequency inside the selected subarray
X: the mode (the largest frequency element in the subarray)
```

Time complexity = $O(n)$

## Resource
[Prefix sum O(n)](https://leetcode.com/problems/maximum-frequency-after-subarray-operation/solutions/6330142/python-solution-prefix-sums-o-n-beats-100-in-contest/?envType=company&envId=amazon&favoriteSlug=amazon-thirty-days)
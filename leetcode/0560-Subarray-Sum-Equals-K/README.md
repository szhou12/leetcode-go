# [560. Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/description/)

## Solution idea
### Prefix Sum
1. Calculate prefix sum and try all subarrays

### Prefix Sum + Hash Map
1. Hash map to store the frequency of prefix sum
2. The idea behind this approach is as follows: If the cumulative sum(represented by sum[i] for sum up to i-th index) up to two indices is the same, the sum of the elements lying in between those indices is zero. Extending the same thought further, if the cumulative sum up to two indices, say i and j is at a difference of k i.e. if sum[i]−sum[j]=k, the sum of elements lying between indices i and j is k.

Based on these thoughts, we make use of a hashmap map which is used to store the cumulative sum up to all the indices possible along with the number of times the same sum occurs. We store the data in the form: `{prefix sum[0:i]: count}`. We traverse over the array nums and keep on finding the cumulative sum. Every time we encounter a new sum, we make a new entry in the hashmap corresponding to that sum. If the same sum occurs again, we increment the count corresponding to that sum in the hashmap. Further, for every sum encountered, we also determine the number of times the sum `sum−k` has occurred already, since it will determine the number of times a subarray with sum k has occurred up to the current index. We increment the count by the same amount.

After the complete array has been traversed, the count gives the required result.

Basically, the equality: count of k as of sum = count of (sum - k)

```
<---         B = sum     --->
[ .. A = (sum-k) ..] .. k .. ]
                   i         j
```

Note: we need a seed `dict[0] = 1`. It means to capture the case where `B = nums[0:j] = k`, the corresponding `A` is empty array. We should make `dict[sum - k] = dict[0]` valid.

Time complexity = $O(n)$

Space complexity = $O(n)$
# [2281. Sum of Total Strength of Wizards](https://leetcode.com/problems/sum-of-total-strength-of-wizards/description/)

## Solution idea
### Prefix Sum + Stack
1. for a given `i`, how to find all subarrays where `s[i]` is the smallest element?
```
prevSmaller[i] := previously smaller element index (Note: in this problem, we are actually finding previously smaller or equal element index)
nextSmaller[i] := next smaller element index
```
```
[idx] a X X X X i X X X b
        ---x---   --y--
        1 2 3 4   3 2 1  // # times included in subarrays   
```
- Implementation: Use Stack to find `prevSmaller` and `nextSmaller` in $O(n)$ time

2. how many such subarrays do we have?
- Number of left boundaries can choose: `a+1 -> i`
    - `x = i-a` // x number of left boudaries to choose
- Number of right boundaries can choose: `i -> b-1`
    - `y = b-i` // y number of right boundaries to choose
- Total # subarrays can form = `x * y`

3. how to calculate total strength of subarrays with `i` being smallest element?
- Notice that as an element gets closer to i, the more often it is included in a subarray
- specifically, given a fixed right boundary, 
    - leftmost element can be included once,
    - 2nd leftmost element can be included twice, 
    - 3rd leftmost element can be included thrice,
    - 4th leftmost element can be included four times.
- There are total y number of right boundaries to choose.
- Thus, total contribution of LHS to subarrays: `sum = (nums[a+1]*1 + nums[a+2]*2 + nums[a+3]*3 + nums[a+4]*4) * y`

- similarly, given a fixed left boundary,
    - rightmost element can be included once,
    - 2nd rightmost element can be included twice, 
    - 3rd rightmost element can be included thrice.
- There are total x number of left boundaries to choose.
- Thus, total contribution of RHS to subarrays: `sum = (nums[i+1]*3 + nums[i+2]*2 + nums[i+3]*1) * x`

- i-th element itself is included in all subarrays
- Thus, total contribution of i-th element to subarrays `sum = nums[i] * x*y`

- So, 
```
sum[i] = 
  (nums[a+1]*1 + nums[a+2]*2 + nums[a+3]*3 + nums[a+4]*4) * y
+ (nums[i+1]*3 + nums[i+2]*2 + nums[i+3]*1) * x
+ (nums[i]) * x*y

res[i] = sum[i] * nums[i]
```

4. How to efficiently calculate `(nums[a+1]*1 + nums[a+2]*2 + nums[a+3]*3 + nums[a+4]*4)`?
- Use **prefix sum** to calculate range sum `(nums[a+1]+nums[a+2]+nums[a+3]+nums[a+4])`
- can we use a similar concept?
- Define a new prefix sum: 
```
presum2[i] = nums[1]*1 + nums[2]*2 +...+ nums[i]*i

presum2[a+4] - presum2[a]
= nums[a+1]*(a+1) + nums[a+2]*(a+2) + nums[a+3]*(a+3) + nums[a+4]*(a+4)
= (nums[a+1]*1 + nums[a+2]*2 + nums[a+3]*3 + nums[a+4]*4) + a * (nums[a+1] + nums[a+2] + nums[a+3] + nums[a+4])
= what-we-want + a * (presum[a+4] - presum[a])

So, LHS = (presum2[a+4] - presum2[a]) - a * (presum[a+4] - presum[a])
Generically, LHS = (presum2[i-1] - presum2[a]) - a * (presum[b-1] - presum[a])

presum2[b-1] - presum2[i]
= nums[i+1]*(i+1) + nums[i+2]*(i+2) + nums[i+3]*(i+3)
= b * (presum[b-1] - presum[i]) - (nums[i+1]*3 + nums[i+2]*2 + nums[i+3]*1)
= b * (presum[b-1] - presum[i]) - what-we-want

So, RHS = b * (presum[b-1] - presum[i]) - (presum2[i+3] - presum2[i])
Generically, RHS = b * (presum[b-1] - presum[i]) - (presum2[b-1] - presum2[i])
```

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 2281. Sum of Total Strength of Wizards](https://www.youtube.com/watch?v=HGCm9PkFd58&ab_channel=HuifengGuan)
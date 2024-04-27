# [2918. Minimum Equal Sum of Two Arrays After Replacing Zeros](https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/description/)

## Solution idea
### Greedy Algorithm
1. Missing places 必须至少填1. 所以, 有missing places的array的array sum只会单增。
2. 切入点: 看每个array是否有missing places? 如果存在array没有missing places, 说明它的array sum是不可能再增加. 
    1. 如果两个array都没有missing places, 那么它们各自的array sum都不可能变动。如果此时两者的array sum相等，这就是唯一的equal sum。否则，-1 (得不到equal sum).
    2. 如果其中一个array没有missing places, 另一个有。那么，没有missing places的array sum不可能再增加，这个array sum本身就是唯一的equal sum. 如果它比另一个array所有missing places都填1以后的array sum还小，那么-1 (得不到equal sum).
    3. 如果两个array都有missing places, 那么一定能找到equal sum。谁填完1后的array sum**较大**，就是可以得到的最小equal sum.
    
Time complexity = $O(1)$

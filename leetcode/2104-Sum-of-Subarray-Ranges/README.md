# [2104. Sum of Subarray Ranges](https://leetcode.com/problems/sum-of-subarray-ranges/description/)

## Solution idea
### Stack - PrevSmaller, NextSmaller; prevGreater, nextGreater
1. array not sorted
2. given an array, # of subarrays?
    - O(n^2)
    - set a left-end, set a right-end, thus total # = n*(n-1)/2
3. can you efficiently find all subarrays? 
4. for a subarray, can you efficiently find its range?

sum of ranges = subarray1(max - min) + subarray2(max - min) + ...
    = {subarray1(max) + subarray2(max) + ...} - {subarray1(min) + subarray2(min) + ...}
    = sum of subarray maximums - sum of subarray minimums
    -> similar question: leetcode 907 Sum of Subarray Minimums

5. sum of subarray minimums
    -> instead of checking subarrays one by one and find their own mins
    -> 转换视角: given a num, how many subarrays regards this num as min?
    -> solver: stack | prevSmaller, nextSmaller
    -> given a num at i, 
    move in left until encounter the first smaller value L, 
    move in right until encounter the first smaller value R
    -> total subarrays with i being min = (i - L) * (R - i)
    -> 持续递增stack: 一旦来者 < 守门员，守门员出去，标记来者为守门员的nextSmaller
    -> 持续递减stack: 一旦来者 > 守门员，守门员出去，标记来者为守门员的nextGreater
    -> 守门员 = given a num at i
    -> nextSmaller, nextGreater: 从左往右撸
    -> prevSmaller, prevGreater: 从右往左撸
    -> 突发状况1: 如果撸到底都没有符合条件的来者, 
    dummy nextSmaller/nextGreater = n, dummy prevSmaller/nextGreater = -1
    -> 突发状况2: 
    多个元素是相同的最小值，约定最左边的为唯一指定的最小值
    prevSmaller 变成 nextSmallerOrEqual 
    nextSmaller 仍是 nextSmaller

## Resource
[【每日一题】LeetCode 2104. Sum of Subarray Ranges](https://www.youtube.com/watch?v=xba0NzSbuas&t=300s&ab_channel=HuifengGuan)
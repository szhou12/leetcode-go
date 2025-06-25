# [540. Single Element in a Sorted Array](https://leetcode.com/problems/single-element-in-a-sorted-array/description/)

## Solution idea
### Binary Search
We have to run in $O(\log n)$ time and $O(1)$ space. So we know for sure that we need to use binary search. The question is: how to find the monotonicity in this problem?
1. Where is the monotonicity? Observe that the problem is set up in a way that only one element will be single while all other elements will appear in pairs. So by looking at their indices, we can find out that: before the single element comes out, the first of a pair is even-indexed while the second is odd-indexed; after the single element, the first of a pair changes to odd-indexed while the second is even-indexed. So we can summarize: as long as this (even-indexed, odd-indexed) property remains, there is no answer to the left, we look its right side. Otherwise, (i.e., (odd-indexed, even-indexed) property shows up), `mid` is potentially the answer.
2. The next big question I may have is: when accessing `nums[mid+1]` and `nums[mid-1]`, why will it not cause out-of-bound error? The answer is: `nums` will always be odd-length and we only access `mid+1` when `mid` is even, and we only access `mid-1` when `mid` is odd. `mid=0` is always even and `mid=n-1` is always odd.

Time complexity = $O(\log n)$

Space complexity = $O(1)$`
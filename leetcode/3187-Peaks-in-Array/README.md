# [3187. Peaks in Array](https://leetcode.com/problems/peaks-in-array/description/)

## Solution idea
### Segment Tree (sum)
1. segment tree (sum) not on `nums[]`, but on `peak[]`
2. `pushDown()` not necessary for queryRange() bc no updateRange()
3. when `queryRange(l, r)`: if `l==r`, no peak for single element
4. when `queryRange(l, r)`: if `l<r`, don't count peaks on both ends of `subarray[l:r]` i.e. `peak[l]` and `peak[r]` e.g. `[0 0 1] => 0 peaks`
5. because we can't count peaks on both ends of `subarray[l:r]`, we need to keep updating `peak[]`, meaning only updating `nums[]` is not enough
6. every time `nums[i]` is updated, need to check 3 indices: `i-1`, `i`, `i+1`, whether their corresponding numbers create new peaks or slience peaks, respectively.

Time complexity = $O(Q*\log n)$ where $Q = $ number of queries, $n = $ length of `nums[]`

## Resource
[Segment Tree Solution](https://leetcode.com/problems/peaks-in-array/solutions/5320039/segment-tree/)
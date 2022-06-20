# [973. K Closest Points to Origin](https://leetcode.com/problems/k-closest-points-to-origin/)

## Solution idea:

### Idea 1
1) sort input array by distance
2) output first k elements
Time = $O(n\log n)$

### Idea 2
Use a max-heap whose value considers the distance
Keep inserting elements and pop the remained k elements
Time = $O(n\log n)$

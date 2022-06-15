# [825. Friends Of Appropriate Ages](https://leetcode.com/problems/friends-of-appropriate-ages/)

## Solution idea:
### Brute force method
step 1: sub-function `makeRequest(x,y)` that checks if `y` will make a request to `x` by the listed conditions

step 2: use a map `countMap` that counts number of people at each unique age because age only ranges from $[1,120]$

step 3: nested-for loop. For any two keys `x, y` of `countMap` that satisfies `makeRequest(x,y)`:

case 1: if `x==y`, then # of requests = `countMap[x] * (countMap[y] - 1)` ($-1$ to exclude friend request itself)

case 2: if `x!=y`, then # of requests = `countMap[x] * countMap[y]`

Time complexity = $O(n^2)$
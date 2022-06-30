# [2320. Count Number of Ways to Place Houses](https://leetcode.com/problems/count-number-of-ways-to-place-houses/)

## Solution idea
### Key Oberservation 1
街两边怎么放房子是完全独立 且 对称的，故只需要算一侧怎么放房子.

### Key Oberservation 2
This problem is very similar to 0198. House Robber.

Difference:

This problem: total numbers for a placement plan where no two adjacent houses.

House Robber: max sum for a robbing plan where no two adjacent houses.

### Key Observation 3
Unlike House Robber where it can be solved by 1-D DP, this problem asks for the total # of valid plans for any location $i$, we need to consider:

Case 1: we place a house at $i$

Case 2: we don't/can't place a house at $i$

How to capture two cases at the same time?

Answer is to use 2-D DP[][2]

Define $DP[i][0]$ as total # of valid plans when no house at $i$-th place; 

Define $DP[i][1]$ as total # of valid plans when placing a house at $i$-th place.

Base Case:

$DP[0][0]=1$ means no house placed at non-existing place, it's considered a valid plan.

$DP[0][1]=0$ means a house placed at non-existing place, logically impossible, hence, no valid plan.

Recurrence:

$DP[i][1] = DP[i-1][0]$ 

(if we place a house at $i$, then we must not place a house at $i-1$)

$DP[i][0] = DP[i-1][1] + DP[i-1][0]$ 

(if we don't place a house at $i$, it means either we already placed a hosue at $i-1$ or we also didn't place a house at $i-1$)

Time complexity = $O(n)$

## Resources
[【每日一题】LeetCode 2320. Count Number of Ways to Place Houses](https://www.youtube.com/watch?v=Q5JUafX3iMc&t=1114s)


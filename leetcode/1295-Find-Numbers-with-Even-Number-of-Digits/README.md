# [1295. Find Numbers with Even Number of Digits](https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/)

## Solution idea
### Math
1. Divide by 10 until 0, count the number of times

Time complexity = $O(n \log_{10} m)$ where $m$ is the max number in nums. It will conduct $\log_{10} m$ times of division.

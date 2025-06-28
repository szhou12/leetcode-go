# [1482. Minimum Number of Days to Make m Bouquets](https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/description/)

## Solution idea
### Binary Search Guess Min
1. Monotonicity: the more days we wait, the more flowers will bloom, more likely to make m bonquets.
2. Note: problem requires that one bonquet is made of k-adjecent (consecutive) flowers. This guides us to design `isOk()`.

Time complexity = $O(n \log D)$ where $D$ is the max day to bloom.
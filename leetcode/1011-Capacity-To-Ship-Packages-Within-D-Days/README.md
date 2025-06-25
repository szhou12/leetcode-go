# [1011. Capacity To Ship Packages Within D Days](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)

## Solution idea
### Binary Search 猜答案
1. 寻找单调性 Monotonicity: higher the capacity, the fewer days needed to ship, more likely to ship all within `days`.

Time complexity = $O(n \log m)$ where $n$ is the number of weights, $m$ is the sum of weights.
# [1889. Minimum Space Wasted From Packaging](https://leetcode.com/problems/minimum-space-wasted-from-packaging/description/)

## Solution idea
### Binary Search - Find Upper Bound
#### 要点总结
站在每个供应商(`boxes[j]`)的角度，按从小到大的顺序遍历其提供的 box size。对于每个 box size，利用binary search确定它能够覆盖的 packages 区间 (upper bound)，利用**前缀和**快速计算该区间内 package 的总体积，从而求出对应的 wasted space。将所有 box size 的浪费空间累加即可得到该供应商的总浪费空间。注意边界情况：若该供应商最大的 box size 仍无法覆盖最大的 package，则说明该供应商不可行，直接跳过。最终在所有可行供应商的总浪费空间中取最小值即为答案。

Time complexity = $O(k * (mlogm + mlogn)) = O(km * logn)$ where $k =$ length of `boxes`, $m =$ average length of `boxes[j]`, $n =$ length of `packages`.


## Resource
[【每日一题】1889. Minimum Space Wasted From Packaging, 6/10/2021](https://www.youtube.com/watch?v=55EpLON-Xdo)
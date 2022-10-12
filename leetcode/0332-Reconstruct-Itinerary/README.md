# [332. Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/)

## Solution idea

### DFS

* 暴力搜索所有可行路径，返回字典序最小的可行解

* 终止条件：根据题意，每一个ticket只能用一次。每个ticket包含两个地点，一个可行解经过每一个地点有且只有一次，
所以，终止条件是 append进的地点数 = ticket数+1

* 先对每一个出发地对应的抵达地的list进行字典序排序，那么，这样DFS找到的第一个可行解就是字典序最小的解，所以，找到第一个就可以结束了

## Resource
[代码随想录-332.重新安排行程](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0332.%E9%87%8D%E6%96%B0%E5%AE%89%E6%8E%92%E8%A1%8C%E7%A8%8B.md)

最优解-欧拉路径：
[【每日一题】332. Reconstruct Itinerary, 6/28/2020](https://www.youtube.com/watch?v=5yM3H0UgXTo&ab_channel=HuifengGuan)


# [407. Trapping Rain Water II](https://leetcode.com/problems/trapping-rain-water-ii/description/)

## Solution idea
### BFS + PQ
#### 要点总结
1. 海水倒灌: “剥洋葱”似的，从最外层开始，用PQ找到最低的位置h，海水涨到h的海平面就可以往内陆进行倒灌；如果PQ弹出来的最低位置h'都已经比当前海平面高了，海平面就要升高到h'。每一次Pop: 1. 更新海平面 if 当前cell的高度 > 当前海平面; 2. 计算一次倒灌的海水高度(接到的雨水)=当前海平面-当前cell的高度

time complexity = $O(mn \log mn)$

## Resource
[【每日一题】407. Trapping Rain Water II , 2/17/2021](https://www.youtube.com/watch?v=uupOnJJxPbI&t=1s&ab_channel=HuifengGuan)
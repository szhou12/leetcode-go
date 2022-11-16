# [1004. Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/description/)

## Solution idea

### Two Pointers 同向而行
* **Key Idea**: 固定一个边界，不停延伸另一个边界至最长
* **实现方式之一**: 固定左边界，不停延伸右边界至最长
    * **此实现方式的注意事项**:  
        1. 延伸右边界有可能出现出界的情况，每次延伸之前一定要检查!!!
        2. 优化: 每一轮挡板l, 挡板r不从挡板l处重新开始, 而是从上一轮r停留处继续往前探


Time complexity = $O(n)$

## Resource
[【每日一题】1004. Max Consecutive Ones III, 9/17/2020](https://www.youtube.com/watch?v=Ti9_4NVDzdg&ab_channel=HuifengGuan)
* 参考答案固定右边界，不停缩短左边界。优点：实现时不用处理右边界出界的情况；缺点：有点反直觉
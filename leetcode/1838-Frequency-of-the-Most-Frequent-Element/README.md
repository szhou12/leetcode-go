# [1838. Frequency of the Most Frequent Element](https://leetcode.com/problems/frequency-of-the-most-frequent-element/description/)

## Solution idea
### Sliding Window (非模版)
#### 思路总结
1. 提升frequency的目标num一定是input array中原本就存在的一个num。e.g. `[1, 2, 4]` 目标num一定是 1, 2 分别加到2，而不会是加到3。因为拿一个已经存在的num作为目标所需要的提升frequency操作一定 <= 拿一个不存在的num作为目标所需要的提升frequency操作。
2. 挑哪一个num作为目标?
    * 挨个试一遍
3. 目前，挨个试的时候有一个问题。input array是无序的，哪些位子上的数字头上要提升freq不固定，这样不好统计。怎么办？
    * 从小到大排序。这样就可以使用滑窗。对任一个目标 `nums[right]`，只需要往前找到左边界 `left`
    * `X  X  [left  X  X  X  right-1  right]  X  X`
    * 滑窗的机制就确定了：依次固定右边界作为目标num, 收缩左边界直到满足约束条件 (提升frequency的操作数 <= k)
4. 每次固定一个新的右边界，左边界都需要从头开始吗？
    * 不需要。这里要用到 Induction 思想，累加操作数
    ```
    // 假设已经算好了右边界 = r:
    target = nums[r] => 滑窗 = [l : r]
    操作数 = count
    // 吃 - 移动右边界 = r + 1:
    // 增量 = 所有窗口内元素增加前后右边界的差值
    target = nums[r + 1] => 滑窗 = [l : r+1]
    操作数 = count += (r - l + 1) * (nums[r+1] - nums[r])
    // 吐 - 收缩左边界:
    // 减量 = 左边界元素以当前右边界为目标时的增量
    操作数 = count -= nums[r+1] - nums[l]
    ```

Time complexity = $O(n\log n) + O(n) = O(n\log n)$

## Resource
[【每日一题】1838. Frequency of the Most Frequent Element, 4/28/2021](https://www.youtube.com/watch?v=c-Jt5jFKNv0&ab_channel=HuifengGuan)
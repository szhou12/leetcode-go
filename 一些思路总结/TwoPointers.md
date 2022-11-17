# Two Pointers

## 题目归类

### 一个数组 + 双指针同向而行

* 移动所有0并且保序: [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/)
    * 在需要**下落物体类**游戏中是非常有用的helper function

* 求满足条件的Subarray个数: [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)

* 有翻转次数找最长连续1: [1004. Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/description/)
    * 总体思路: 固定一个边界，不停延伸另一个边界至最长
    * i.e. 固定左边界，不停延伸右边界至最长

### 一个数组 + 双指针相向而行

* 判断唯一山峰，左右两边单调递增: [941. Valid Mountain Array](https://leetcode.com/problems/valid-mountain-array/description/)

* 平方一个有序数组: [977. Squares of a Sorted Array](https://leetcode.com/problems/squares-of-a-sorted-array/description/)

### 两个数组 + 双指针同向而行

* 跳过重复元素，匹配subsequence: [925. Long Pressed Name](https://leetcode.com/problems/long-pressed-name/description/)

### X-Sum
* [1. Two Sum](https://leetcode.com/problems/two-sum/)
* [15. 3Sum](https://leetcode.com/problems/3sum/)
    * 第一步: 排序
    * 第二步: 对每个元素，找符合的two-sum
* [16. 3Sum Closest](https://leetcode.com/problems/3sum-closest/)
* [18. 4Sum](https://leetcode.com/problems/4sum/)
* [454. 4Sum II](https://leetcode.com/problems/4sum-ii/)


## 解法归类
### Subarray 类型

* 一般像是求满足条件的subarray个数

* 总体思路:
    * 求 Subarray 只需要确定左、右边界，左右边界圈出来的长度就是Subarray的个数

* 具体实现的思路:
    * 如何确定左右边界？
    * 用快慢指针
    * **固定一个边界，另一个边界不停探索至极限**
    * 到达极限后，收缩固定的那个边界，另一个边界再不停探索至极限
    * 循环往复，直至出界

* 实现的注意事项:
    * 如果选择的实现是：固定左边界，右边界不停探索
        1. 检查右边界**出界**的情况
        2. 检查左边界**超过**右边界的情况

* 求满足条件的Subarray个数: [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)
    * 固定左边界，右边界不停探索至极限
    * 到达极限后，收缩左边界(右移一位)，右边界再探索至极限
    * **注意！！！** 探索前先检查 slow超过fast的情况

### Sliding Window
* **解题思路**: **吃了吐**
* i.e. 分别分析清楚: **吃进去**的各种情况，**吐出来**的各种情况

* **滑动窗口算法框架**
```
func slidingWindow(s string) {
    window := make(map[byte]int)
    left, right := 0, 0
    for left < len(s) {
        rightElement := s[right] // 吃: rightElement 是将移入窗口的字符
        right++ // 增大窗口
        [...] // 进行窗口内数据的一系列更新

        /*** debug 输出的位置 ***/
        // 注意是左闭右开
        fmt.Println("window: [" + left + ", " right + ")")

        // 判断左侧窗口是否要收缩
        for (window needs shrink) {
            leftElement := s[left] // 吐: leftElement 是即将移出窗口的字符
            left-- // 缩小窗口
            [...] // 进行窗口内数据的一系列更新
        }

        [update result (can be anywhere in the loop)]
    }
}
```


* 最多满意客户: [1052. Grumpy Bookstore Owner](https://leetcode.com/problems/grumpy-bookstore-owner/description/)

* 找子串位置: [28. Find the Index of the First Occurrence in a String](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/)
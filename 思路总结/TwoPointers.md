# Two Pointers
## 目录

## 题目归类
* [Sliding Window](#sliding-window)

### 一个数组 + 双指针同向而行

* 移动所有0并且保序: [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/)
    * 物理意义:
        * `[0, left)`: 只接纳扔过来的非0元素
        * `left`: 当前为0, 下一个非0元素要放的位置
        * `[left, right)`: 0元素
        * `[right, n)`: 未知元素
    * 在需要**下落物体类**游戏中是非常有用的helper function

* 求满足条件的Subarray个数: [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)

* 有翻转次数找最长连续1: [1004. Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/description/)
    * 总体思路: 固定一个边界，不停延伸另一个边界至最长
    * i.e. 固定左边界，不停延伸右边界至最长

* 找差值为K的所有不同pair: [532. K-diff Pairs in an Array](https://leetcode.com/problems/k-diff-pairs-in-an-array/description/)
    * 结合了 **sort**
    * 再进行两个指针同向而行

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

### Sliding Window
* **解题思路**: **吃了吐**
* i.e. 分别分析清楚: **吃进去**的各种情况，**吐出来**的各种情况
* **Sliding Window 题型**可分为两大类:
    1. **固定的窗口**长度
    2. **可变的窗口**长度 (注意！可变的窗口长度类型 等价于 Subarray 类型, 可用同一个思路/模版解)

#### 两套模版

* **模版Fix**
    * 难点主要是以下几点:
        1. 吃进新元素/吐出旧元素 时，如何更新窗口内数据？一旦确定，吃和吐是镜面对称的，实现起来就很容易
        2. 更新result的条件是什么？在哪里更新result？
        3. 窗口收缩的条件是什么？(如果是定长窗口，比较容易，一般是维护定长)
```go
func slidingWindow(s string) {
    window := make(map[byte]int)
    left, right := 0, 0
    for right < len(s) {
        rightElement := s[right] // 吃: rightElement 是将移入窗口的字符
        right++ // 增大窗口
        [...] // A: 吃进新元素后, 进行窗口内数据的一系列更新 (与 B段 写法一般是镜面对称的)

        /*** debug 输出的位置 ***/
        // 注意是左闭右开
        fmt.Println("window: [" + left + ", " right + ")")

        // 判断左侧窗口是否要收缩
        for (window needs shrink) { // 如果window定长，这里条件一般是维护定长: (right - left >= 定长)
            
            [update result (can be anywhere in the loop)] (可能在这里left收缩前 update result)

            leftElement := s[left] // 吐: leftElement 是即将移出窗口的字符
            left-- // 缩小窗口
            [...] // B: 吐出旧元素后, 进行窗口内数据的一系列清理和更新 (与 A段 写法一般是镜面对称的)
        }

        [update result (can be anywhere in the loop)] (也有可能在这里 update result)
    }
}
```

* **模版Flex**
    * 核心思路: 固定左边界，右边界不停探索至极限
```go
func slidingWindow(s string) int {
    window := make(map[byte]int)
    right := 0
    res := 0

    for left := 0; left < len(s); left++ {
        // 固定左边界，延伸右边界至极限
        // 至极限条件:  (right未出界) && (right依然可延伸的条件)
        for right < len(s) && [right依然可延伸的条件] {
            rightElement := s[right]
            window[rightElement] do something // A: 吃进新元素后, 进行窗口内数据的一系列更新
            right++
        }

        // NOTE: 这里通常需要check哪种情况下才update result
        // 一般的, right停下来不能继续延伸时, 才能update result
        // 有的题目, right出界时right可能依然可延伸(不满足停下来的条件), 此时不能update result
        if [right停下来不能继续延伸时] {
            res = [update result]
        }

        // 吐
        leftElement := s[left]
        window[leftElement] do something // B: 吐出旧元素后, 进行窗口内数据的一系列清理和更新
    }

    return res

}
```

#### Sliding Window 长度固定

* **模版Fix** 比较适合解决 Sliding Window 为定长的题目 (容易写出正确的code), 所以 定长的题目 优先考虑 **模版Fix**.
    * **模版Fix** 也可以解决 长度可变 的题目，但是实际操作中实现上不容易保证正确

* 最多满意客户: [1052. Grumpy Bookstore Owner](https://leetcode.com/problems/grumpy-bookstore-owner/description/)

* 找子串位置: [28. Find the Index of the First Occurrence in a String](https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/)

* 找到字符串中所有字母异位词: [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/description/)

* 字符串的排列子串: [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/description/)

* 定长为k的subarray最大和: [2461. Maximum Sum of Distinct Subarrays With Length K](https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/)




#### Sliding Window 长度可变 == Subarray 类型

* **模版Flex** 比较适合解决 Sliding Window 为长度可变的题目 (容易写出正确的code), 所以 长度可变的题目 优先考虑 **模版Flex**.

* Subarray 类型
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
    * 不是很常规的题, 因为右边界延伸过程中的都是所求, 要注意处理右边界延伸的初始状态下就不满足要求的情况 (会导致左边界**超过**右边界)
    * 固定左边界，右边界不停探索至极限
    * 到达极限后，收缩左边界(右移一位)，右边界再探索至极限
    * **注意！！！** 探索前先检查 slow超过fast的情况

* 求满足条件的二进制Subarray个数： [930. Binary Subarrays With Sum](https://leetcode.com/problems/binary-subarrays-with-sum/description/)
    * 需要检查右边界**出界**的情况
    * 固定左边界，右边界不停探索至极限
    * 预处理: 每个元素后面跟多少个0

* 有翻转次数找最长连续1: [1004. Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/description/)
    * 固定左边界，不停延伸右边界至最长

* 有翻转次数找最长连续T/F: [2024. Maximize the Confusion of an Exam](https://leetcode.com/problems/maximize-the-confusion-of-an-exam/description/)
    * 做两遍 *LC 1004*

* 无重复字符的最长子串: [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/)

* 最小覆盖子串: [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/description/)

* 最小变换使数组变连续: [2009. Minimum Number of Operations to Make Array Continuous](https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/description/)
    * 预处理: 先去重，再增序排列

* 每个字母至少k个的最长子串: [395. Longest Substring with At Least K Repeating Characters](https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/)
    * 为防止右边界无限延伸，需要额外增加一个条件: sliding window 内 固定只框住 `m` 个不同的字母 ($m = 1 \cdots 26$)

* :yellow_circle: 包含所有类型元素的字串个数: [1358. Number of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/description/)
    * 典型题
    * 两个易错点：
        1. 每次update result是数多少个子串? 
        2. 无条件update result吗？

* :red_circle: 左右两头包含每个元素至少k个: [2516. Take K of Each Character From Left and Right](https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/description/)
    * 转化思想: 左右两头 取反 :arrow_right: 中间滑窗

* :red_circle: 最长"好"子数组: [2401. Longest Nice Subarray](https://leetcode.com/problems/longest-nice-subarray/description/)

* :red_circle: :secret: 可以得到bitwise OR最大值的最短子数组长度: [2411. Smallest Subarrays With Maximum Bitwise OR](https://leetcode.com/problems/smallest-subarrays-with-maximum-bitwise-or/description/)
    *  左边界从左往右移动。固定左边界，右边界尽量缩短。
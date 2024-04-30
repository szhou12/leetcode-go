# Two Pointers

## 目录
* [X-Sum](#x-sum)
* [Sliding Window](#sliding-window)
    * [套路总结](#Methodology)
    * [Fix 模版题](#fix-模版题)
    * [Flex 模版题](#flex-模版题)
    * [非 模版题](#非-模版题)
    * [需要同时满足两个及以上条件](#同时满足两个及以上条件)

## 一个数组 + 双指针同向而行

* 移动所有0并且保序: [283. Move Zeroes](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0283-Move-Zeroes)
    * 物理意义:
        * `[0, left)`: 只接纳扔过来的非0元素
        * `left`: 当前为0, 下一个非0元素要放的位置
        * `[left, right)`: 0元素
        * `[right, n)`: 未知元素
    * 在需要**下落物体类**游戏中是非常有用的helper function

* 求满足条件的Subarray个数: [713. Subarray Product Less Than K](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0713-Subarray-Product-Less-Than-K)

* 有翻转次数找最长连续1: [1004. Max Consecutive Ones III](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1004-Max-Consecutive-Ones-III)
    * 总体思路: 固定一个边界，不停延伸另一个边界至最长
    * i.e. 固定左边界，不停延伸右边界至最长

* 找差值为K的所有不同pair: [532. K-diff Pairs in an Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0532-K-diff-Pairs-in-an-Array)
    * 结合了 **sort**
    * 再进行两个指针同向而行

## 一个数组 + 双指针相向而行

* 判断唯一山峰，左右两边单调递增: [941. Valid Mountain Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0941-Valid-Mountain-Array)

* 平方一个有序数组: [977. Squares of a Sorted Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0977-Squares-of-a-Sorted-Array)

### 两个数组 + 双指针同向而行

* 跳过重复元素，匹配subsequence: [925. Long Pressed Name](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0925-Long-Pressed-Name)

## X-Sum
* 两数之和: [1. Two Sum](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0001-Two-Sum)
* 三数之和: [15. 3Sum](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0015-3Sum)
    * Step 1: 排序
    * Step 2: 对每个元素，找符合的 Two-Sum
* 最接近的三数之和: [16. 3Sum Closest](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0016-3Sum-Closest)
* 四数之和: [18. 4Sum](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0018-4Sum)
* 四数之和II: [454. 4Sum II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0454-4Sum-II)



## Sliding Window

### Methodology
* **解题思路: "吃了吐"**
* i.e. 分别分析清楚: **吃进去**的各种情况，**吐出来**的各种情况
* **Sliding Window 题型**可分为两大类:
    1. **固定的窗口**长度
    2. **可变的窗口**长度 (注意！可变的窗口长度类型 等价于 Subarray 类型, 可用同一个思路/模版解)

#### 两套模版

* **模版Fix**: 定长窗口滑窗
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

* **模版Flex**: 伸缩窗口滑窗
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

        // 这里通常需要check不同情况下对应的update result
        //  情况1: right出界 但 right依然可延伸，如何update result
        //  情况2: right没出界 但 right不可延伸，如何update result
        //  情况3: right出界 且 right不可延伸，如何update result
        if [哪种情况] {
            res = [update result]
        }

        // 吐
        leftElement := s[left]
        window[leftElement] do something // B: 吐出旧元素后, 进行窗口内数据的一系列清理和更新
    }

    return res

}
```

### Fix 模版题

* **模版Fix** 比较适合解决 Sliding Window 为定长的题目 (容易写出正确的code), 所以 定长的题目 优先考虑 **模版Fix**.
    * **模版Fix** 也可以解决 长度可变 的题目，但是实际操作中实现上不容易保证正确

* :yellow_circle: 最短可能的子串的bit OR总和>=k: [3097. Shortest Subarray With OR at Least K II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3097-Shortest-Subarray-With-OR-at-Least-K-II)
    * Binary Search 猜答案 + Sliding Window + Bitwise OR Sum
    * 定长SLiding Window没有按照模版实现

* 最多满意客户: [1052. Grumpy Bookstore Owner](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1052-Grumpy-Bookstore-Owner)

* 找子串位置: [28. Find the Index of the First Occurrence in a String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0028-Implement-strStr)

* 找到字符串中所有字母异位词: [438. Find All Anagrams in a String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0438-Find-All-Anagrams-in-a-String)

* 字符串的排列子串: [567. Permutation in String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0567-Permutation-in-String)

* 定长为k的subarray最大和: [2461. Maximum Sum of Distinct Subarrays With Length K](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2461-Maximum-Sum-of-Distinct-Subarrays-With-Length-K)


* :red_circle: :secret: 满足XOR值的最短子串: [2564. Substring XOR Queries](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2564-Substring-XOR-Queries)
    * Sliding Window (Fix) + **Bitwise XOR**

* :red_circle: :secret: [2747. Count Zero Request Servers](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2747-Count-Zero-Request-Servers)
    * 类 Sliding Window (Fix)


### Flex 模版题

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


* 求满足条件的Subarray个数: [713. Subarray Product Less Than K](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0713-Subarray-Product-Less-Than-K)
    * 不是很常规的题, 因为右边界延伸过程中的都是所求, 要注意处理右边界延伸的初始状态下就不满足要求的情况 (会导致左边界**超过**右边界)
    * 固定左边界，右边界不停探索至极限
    * 到达极限后，收缩左边界(右移一位)，右边界再探索至极限
    * **注意！！！** 探索前先检查 slow超过fast的情况

* 求满足条件的二进制Subarray个数： [930. Binary Subarrays With Sum](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0930-Binary-Subarrays-With-Sum)
    * 需要检查右边界**出界**的情况
    * 固定左边界，右边界不停探索至极限
    * 预处理: 每个元素后面跟多少个0

* 有翻转次数找最长连续1: [1004. Max Consecutive Ones III](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1004-Max-Consecutive-Ones-III)
    * 固定左边界，不停延伸右边界至最长

* 有翻转次数找最长连续T/F: [2024. Maximize the Confusion of an Exam](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2024-Maximize-the-Confusion-of-an-Exam)
    * 做两遍 *LC 1004*

* 无重复字符的最长子串: [3. Longest Substring Without Repeating Characters](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0003-Longest-Substring-Without-Repeating-Characters)

* 最小覆盖子串: [76. Minimum Window Substring](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0076-Minimum-Window-Substring)

* 最小变换使数组变连续: [2009. Minimum Number of Operations to Make Array Continuous](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2009-Minimum-Number-of-Operations-to-Make-Array-Continuous)
    * 预处理: 先去重，再增序排列

* 每个字母至少k个的最长子串: [395. Longest Substring with At Least K Repeating Characters](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0395-Longest-Substring-with-At-Least-K-Repeating-Characters)
    * 为防止右边界无限延伸，需要额外增加一个条件: sliding window 内 固定只框住 `m` 个不同的字母 ($m = 1 \cdots 26$)

* :yellow_circle: 包含所有类型元素的字串个数: [1358. Number of Substrings Containing All Three Characters](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1358-Number-of-Substrings-Containing-All-Three-Characters)
    * 典型题
    * 两个易错点：
        1. 每次update result是数多少个子串? 
        2. 无条件update result吗？

* :red_circle: 左右两头包含每个元素至少k个: [2516. Take K of Each Character From Left and Right](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2516-Take-K-of-Each-Character-From-Left-and-Right)
    * 转化思想: 左右两头 取反 :arrow_right: 中间滑窗

* :red_circle: :secret: 最长"好"子数组: [2401. Longest Nice Subarray](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2401-Longest-Nice-Subarray)
    * Sliding Window (Flex) + **Bitwise AND**

* :red_circle: :secret: 可以得到bitwise OR最大值的最短子数组长度: [2411. Smallest Subarrays With Maximum Bitwise OR](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2411-Smallest-Subarrays-With-Maximum-Bitwise-OR)
    * Sliding Window (Flex) + **Bitwise OR**
    * Trick 1: 左边界从右往左移动。固定左边界，右边界尽量缩短。
    * Trick 2: 判断一个整数 第k bit位 是否有1: `num>>k & 1 == 1`


* :red_circle: 窗口内统一字母的最长子串: [424. Longest Repeating Character Replacement](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0424-Longest-Repeating-Character-Replacement)
    * 解法一: 26个字母挨个锁定为目标
        * 相同思路的题目: [1004. Max Consecutive Ones III](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1004-Max-Consecutive-Ones-III), [2024. Maximize the Confusion of an Exam](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2024-Maximize-the-Confusion-of-an-Exam)
    * 解法二: 窗口总长度 - 窗口占多数的字母次数 <= k

* :yellow_circle: "好"子串的个数: [2537. Count the Number of Good Subarrays](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2537-Count-the-Number-of-Good-Subarrays)
    * 逆向思维：滑窗物理意义 = 题目要求取反
    * 分类讨论：更新结果前需要分类讨论

* :bell: :yellow_circle: [992. Subarrays with K Different Integers](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0992-Subarrays-with-K-Different-Integers)
    * 2537 的进阶版

* :yellow_circle: 最长的"半重复"子串: [2730. Find the Longest Semi-Repetitive Substring](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2730-Find-the-Longest-Semi-Repetitive-Substring)


* :red_circle: 最长的元素相同的子串: [2831. Find the Longest Equal Subarray](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2831-Find-the-Longest-Equal-Subarray)

* :orange_circle: 在无限增殖的数组中找sum为target的最短子数组: [2875. Minimum Size Subarray in Infinite Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2875-Minimum-Size-Subarray-in-Infinite-Array)
    * Sliding Window: 需要modify. Sliding Window 是找subarray sum的总个数，这里是找其中最短的那一个subarray
    * 无限增殖 = 所有元素按顺序倍数增加 = cyclic = modulo 来简化

### 非 模版题
* :red_circle: :secret: k步内收集最多果子: [2106. Maximum Fruits Harvested After at Most K Steps](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2106-Maximum-Fruits-Harvested-After-at-Most-K-Steps)
    * Sliding Window (非) + Prefix Sum

* :red_circle: :secret: 有限次增量下使相同元素最多: [1838. Frequency of the Most Frequent Element](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1838-Frequency-of-the-Most-Frequent-Element)

### 同时满足两个及以上条件
* 思考方法：要求同时满足两个/多个条件的滑窗 -> "挑软柿子捏"：挑一个容易的条件分割出若干区间，则每个区间内都符合该条件，再看每个区间内有多少子区间符合另一条件

* :red_circle: “完备”子串的个数: [2953. Count Complete Substrings](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2953-Count-Complete-Substrings)
    * Flex 模版 + Fix 模版
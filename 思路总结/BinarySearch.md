# Binary Search

## 目录
* [方法论](#方法论)
* [模版 - upperBound, lowerBound](#模版-find-lower-bound--find-upper-bound)
    * [写法一: sort.Search()](#写法一-sortsearch)
    * [写法二: sort.SearchInts()](#写法二-sortsearchints)
    * [写法三: 自己实现](#写法三-自己实现)
* [经典题](#经典题)
* [Lower Bound & Upper Bound](#lower-bound--upper-bound)
* [First/Last Occurrence](#find-first-occurrence--last-occurrence)
* [Closet Element Greater/Less](#find-closest-element-greatersmaller-than-target)
* [BS猜答案](#binary-search-猜答案)
    * [Guess K-th Element](#guess-k-th-element)
    * [Guess + Math](#guess--math)
    * [Guess Min/Max](#guess-minmax)
* [非常规](#非常规)

## 方法论
* 统一循环条件:
```
for left < right {...}
```
* 终止条件/终止时指针位置状态：
    * Case 1: `left == right`
    * Case 2: `left > right` (当只剩两个元素时，此条件会在下一轮循环触发)

* 根据上述的循环条件，就要自己推导**当只剩两个元素时**, 1. 死循环的情况, 要针对着更改计算 `mid` 的方式; 2. 左、右指针出界的情况，要针对着选择不会出界的那个指针.
    1. `mid = left + (right - left)/2`
        1. `mid` 在两个元素中靠左
            * `[0 (left), 1 (right)]` 中, `mid`会选择 `0 (left)` 指向的元素
            * 收敛时，若有 `left = mid`, 则**会造成死循环** (因为上一轮`left`和本轮的`left`都在同一个位置没有移动)
        2. 收敛时，若有 `right = mid - 1`, 则`right`**会出现出界**的情况， 
            * `[0, 1]` 中, `right`会变成`mid-1=-1` 而出界 (因为此时 `mid` 指向`0`位置)
            * 然而，`left` 一般是 `=mid`/`=mid+1`, 故不会出现出界的情况。所以，返回 `left` 保险
    2. `mid = right - (right - left)/2`
        1. `mid` 在两个元素中靠右
            * `[0 (left), 1 (right)]` 中会选择 `1 (right)` 指向的元素
            * 收敛时，若有 `right = mid`, 则**会造成死循环** (因为上一轮`right`和本轮的`right`都在同一个位置没有移动)
        2. 收敛时，若有 `left = mid + 1`, 则`left`**会出现出界**的情况， 
            * `[0, 1]` 中, `left`会变成`mid+1=2` 而出界 (因为此时 `mid` 指向`1`位置)
            * 然而，`right` 一般是 `=mid`/`=mid-1`, 故不会出现出界的情况。所以，返回 `right` 保险

* :rotating_light: CAUTION!!! :rotating_light: CAUTION!!! :rotating_light: CAUTION!!! :rotating_light:
    * 如果题目涉及左右边界的初始值为`MinInt`和`MaxInt`时，在Go中直接初始为`left, right = math.MinInt, math.MaxInt`会导致溢出。
    * 原因：Go中的`math.MinInt`和`math.MaxInt`会按照操作系统允许的最多位数来表示。e.g. 32位系统中，`math.MinInt`为$-2^{31}$，`math.MaxInt`为$2^{31}-1$。64位系统中，`math.MinInt`为$-2^{63}$，`math.MaxInt`为$2^{63}-1$。
    * 不良结果：因为这个实现，会在计算`mid`的`(right-left)`这一步造成overflow。
    * 解决方法：`left, right := math.MinInt32, math.MaxInt32` (64位系统中) 或者 `left, right := math.MinInt / 2, math.MaxInt / 2`。

* 什么时候需要 Post-Processing ?
    * TLDR: 搜索范围内不一定有解。Post-Processing额外检查loop跳出后给的答案是否符合题意。
    * 使用 Binary Search 时, 只有当题目的**input不一定包含答案target的时候**需要 post-processing, 否则 Binary Search 一定可以找到答案, 也就是不用post-processing
    * `lowerBound(target)` 和 `upperBound(target)` 都需要 post-processing. 因为搜索区间内不一定存在满足target的左侧边界 和 满足target的右侧边界. 
        * e.g. 搜索区间内的值都 `> target`，那就找不到 满足target的左侧边界 (ie. 第一个 `== target` 的元素)
        * e.g. 搜索区间内的值都 `<= target`, 那就找不到 满足target的右侧边界 (ie. 第一个 `> target` 的元素)


## 模版: find lower bound & find upper bound

**Assume input array in ASCENDING order !!!**

### 写法一: sort.Search()
- [func Search | Offical Document](https://pkg.go.dev/sort#Search)
    - TL;DR: `sort.Search(n, f)` returns the first (smallest) **index** that makes condition function `f()` true
```go
func lowerBound(nums []int, target int) int {
    res := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
    return res
}


func upperBound(nums []int, target int) int {
    res := sort.Search(len(nums), func(i int) bool { return nums[i] > target })
    return res
}
```

### 写法二: sort.SearchInts()
- [func SearchInts | Offical Document](https://pkg.go.dev/sort#SearchInts)
```go
func lowerBound(nums []int, target int) int {
    res := sort.SearchInts(nums, target)
    return res
}

func upperBound(nums []int, target int) int {
    res := sort.SearchInts(nums, target+1)
    return res
}
```

### 写法三: 自己实现 
1. `lowerBound()`: Find the **first index** whose element in the array has the **value >= target** (It follows the same concept as `lower_bound()` method in C++)
    1. 情况一: target 不在数组范围中，target 小于数组中所有元素，在**左边界**以外 :arrow_right: 返回数组的第一个元素的index
    2. 情况二: target 不在数组范围中，target 大于数组中所有元素，在**右边界**以外 :arrow_right: 返回数组的最后一个元素的index+1
    3. 情况三: target 在数组范围中，只是数组中不存在 :arrow_right: 返回数组中第一个 > target的元素的index
    4. 情况四: target 在数组范围中，并且数组中存在 :arrow_right: 返回数组中第一个 = target的元素的index
```go
func lowerBound(nums []int, target int) int {
    left := 0
    right := len(nums) - 1

    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] < target { // < target, mid一定不是答案
            left = mid + 1
        } else { // nums[mid] >= target, mid有可能是答案
            right = mid
        }
    }

    // 返回left: 因为left一定不会出界 (一定指向array中的某一个元素)

    // post-processing 简化版
    if nums[left] >= target { // left符合题意, 直接返回left
        return left
    } else { // nums[left] < target, 返回left+1 (情况二)
        return left + 1
    }

    // post-processing 完整版
    // if nums[left] == target { // 情况四
    //     return left
    // } else if nums[left] > target { // 情况一 + 情况三
    //     return left
    // } else { // 情况二
    //     return left + 1
    // }
 }
```
2. `upperBound()`: Find the **first index** whose element in the array has the **value > target** (It follows the same concept as `upper_bound()` method in C++)
    1. 情况一: target 不在数组范围中，target 小于数组中所有元素，在**左边界**以外 :arrow_right: 返回数组的第一个元素的index
    2. 情况二: target 不在数组范围中，target 大于数组中所有元素，在**右边界**以外 :arrow_right: 返回数组的最后一个元素的index+1
    3. 情况三: target 在数组范围中，只是数组中不存在 :arrow_right: 返回数组中第一个 > target的元素的index
    4. 情况四: target 在数组范围中，并且数组中存在 :arrow_right: 返回数组中第一个 > target的元素的index
```go
func upperBound(nums []int, target int) int {
    left := 0
    right := len(nums) - 1

    for left < right {
        mid := left + (right - left) / 2
        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }

    // 返回left: 因为left一定不会出界 (一定指向array中的某一个元素)

    // post-processing 简化版
    if nums[left] > target { // left符合题意, 直接返回left
        return left
    } else { // nums[left] <= target, 返回left+1 (情况二+情况四)
        return left + 1
    }

    // post-processing 完整版
    // if nums[left] == target { // 情况四
    //     return left + 1
    // } else if nums[left] < target { // 情况二
    //     return left + 1
    // } else { // 情况三 + 情况一
    //     return left
    // }
}
```

* **备注**: `lowerBound()`, `upperBound()` 的物理意义也可以根据具体题目进行更改, 并根据具体的物理意义微调代码. 
    * e.g. `lowerBound()`表示搜索区间内 最大的/最后一个 小于target的元素
    * e.g. `upperBound()`表示搜索区间内 最大的/最后一个 等于target的元素


## 经典题
* Classic Binary Search: [704. Binary Search](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0704-Binary-Search)


## Lower Bound & Upper Bound
* :red_circle: 找所有的“美丽”index II: [3008. Find Beautiful Indices in the Given Array II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3008-Find-Beautiful-Indices-in-the-Given-Array-II)
    * 同样的题还有: [3006. Find Beautiful Indices in the Given Array I](https://leetcode.com/problems/find-beautiful-indices-in-the-given-array-i/description/)
    * 与 KMP算法的结合题

* :red_circle: 求符合题意的配对数量: [2563. Count the Number of Fair Pairs](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2563-Count-the-Number-of-Fair-Pairs)
    * 求区间范围: 用到 `upperBound()` 和 `lowerBound()`



## Find First Occurrence / Last Occurrence

* 在排序数组中查找元素的第一个和最后一个位置: [34. Find First and Last Position of Element in Sorted Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0034-Find-First-and-Last-Position-of-Element-in-Sorted-Array)

**一定！一定！一定要明确**target可能在数组中**没有出现**的三种情况，才能正确地写出post-processing中边界的查看条件!!!

* 情况一: target 不在数组范围中，target 过于小，在**左边界**以外. e.g. 数组{3, 4, 5}，target为2
* 情况二: target 不在数组范围中，target 过于大，在**右边界**以外，e.g. 数组{3, 4, 5}，target为6
* 情况三: target 在数组范围中，只是数组中不存在. e.g. 数组{3, 6, 7}, target为5
* 情况四: target 在数组范围中，并且数组中存在. Binary Search 一定能找到.



## Find Closest Element Greater/Smaller Than Target

* 此类题都要在最后注意 Post-process 找不到的情况:
    1. target > 数组中所有值
    2. target < 数组中所有值
    3. target 在数组范围中，只是数组中不存在该值

* :red_circle: 多少种subarray的删法使得剩余元素严格递增: [2972. Count the Number of Incremovable Subarrays II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2972-Count-the-Number-of-Incremovable-Subarrays-II)

* 找最小的大于Target的元素: [744. Find Smallest Letter Greater Than Target](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0744-Find-Smallest-Letter-Greater-Than-Target)




## Binary Search 猜答案

**算法题的解题思路可以概括为两大类：**
1. 正面硬刚：用合适的算法 (DP, DFS, BFS, etc.) 直接解决
2. 二分搜值：猜出答案, 每一轮看看是猜大了还是猜小了？以达到逼近最终答案的目的
    * 此法适用于 题目要求 return single value (scalar) 并且 题目有单调性 + 涉及 max/min value

**要点**: 什么时候可以用？
* 当我们从题目中找到一种 **单调性** 的规律时，一般可以尝试用二分猜答案

**猜答案的具体步骤**
1. 确定搜索空间：也就是确定Binary Search的下界/左边界 和 上界/右边界
    * 因为是在搜索空间里猜答案，所以 左/右边界的物理意义通常就是 return值本身 (i.e. 也就是说，不是index而是value本身)
2. 确定收敛条件：什么条件让我们确定可以舍弃掉一半的空间？舍弃掉小的那一半还是大的那一半？包不包括边界的值？



**经典例题**

### Guess K-th Element

* :red_circle: 第k小配对的绝对差值: [719. Find K-th Smallest Pair Distance](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0719-Find-K-th-Smallest-Pair-Distance)
    * 嵌套二分法: Binary Search 猜一个值 + Upper Bound to count pairs

* :red_circle: 第k小的组合数: [3116. Kth Smallest Amount With Single Denomination Combination](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3116-Kth-Smallest-Amount-With-Single-Denomination-Combination)
    * Binary Search + Grosper's Hack + Exclusion-Inclusion Principle + Divisors

* :red_circle: 有不同元素的subarrays的个数的中位数: [3134. Find the Median of the Uniqueness Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3134-Find-the-Median-of-the-Uniqueness-Array)
    * Binary Search + Sliding Window
    * 联动 [992. Subarrays with K Different Integers](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0992-Subarrays-with-K-Different-Integers)

* :red_circle: Search in Sorted Matrix: [240. Search a 2D Matrix II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0240-Search-a-2D-Matrix-II)
    * 解法本身不是Binary Search, 但是**解法思路可以应用于所有 sorted matrix**, 并且和 Binary Search 结合组成更高效的算法
    * 从 bottom-left / top-right 出发做 matrix traversal

* :red_circle: 无序数组中第k大的元素: [215. Kth Largest Element in an Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0215-Kth-Largest-Element-in-an-Array)
    * 每猜一个数`mid`，统计`nums[]`中大于等于`mid`的元素个数，然后根据个数与`k`的关系，调整搜索区间。

* 非递减矩阵中找第k小的元素: [378. Kth Smallest Element in a Sorted Matrix](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0378-Kth-Smallest-Element-in-a-Sorted-Matrix)
    * 用 max heap 也可以解，但是 Binary Search 更效率

* 乘法表中找第k小的元素: [668. Kth Smallest Number in Multiplication Table](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0668-Kth-Smallest-Number-in-Multiplication-Table)

* 找k个最小和: [373. Find K Pairs with Smallest Sums](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0373-Find-K-Pairs-with-Smallest-Sums)
    * 用 max heap 会超时！！！只能 Binary Search 提高效率

* 非递减数组中找第k小的乘积: [2040. Kth Smallest Product of Two Sorted Arrays]()
    * 找前k小的个数的思路与前面题型不同. main idea: 分类讨论

* :red_circle: :secret: 从矩阵中找第k小 array sum: [1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2040-Kth-Smallest-Product-of-Two-Sorted-Arrays)
    * 找前k小的个数的思路与前面题型不同. main idea: 结合**DFS**
    * 用 **DFS - All Combo Sum** 来找满足当前猜的数 (`mid`) 的array 的个数是否 >= k 个

* :red_circle: :secret: 找第k小 subarray sum: [1918. Kth Smallest Subarray Sum](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1918-Kth-Smallest-Subarray-Sum)
    * 找前k小的个数的思路与前面题型不同. main idea: 转化 subarray sum 为 prefixSum diff, 再结合**Two Pointers**
    * 因为 prefixSum array 是单增的, 所以可以通过 **Two Pointers** 来找满足当前猜的数 (`mid`) 的 subarray 个数是否 >= k 个

### Guess + Math

* :star: 搜索区间为自然数，由于自然数递增排序，所以可以天然地使用二分搜索

* :red_circle: 最大满足特定条件的自然数: [3007. Maximum Number That Sum of the Prices Is Less Than or Equal to K](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3007-Maximum-Number-That-Sum-of-the-Prices-Is-Less-Than-or-Equal-to-K)
    * Binary Search 猜答案 + Digital Counting

* 寻找丑数III: [1201. Ugly Number III](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1201-Ugly-Number-III)
    * *1201 与 1539 思路相近*
    * 需要用到一些数学定理: 容斥原理, LCM, GCD

* 找第k个缺失的自然数: [1539. Kth Missing Positive Number](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1539-Kth-Missing-Positive-Number)
    * * *1201 与 1539 思路相近*

* 阶乘函数后 K 个零: [793. Preimage Size of Factorial Zeroes Function](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0793-Preimage-Size-of-Factorial-Zeroes-Function)

### Guess Min/Max

* :red_circle: 切出最小可允许的identical substring的次数: [3399. Smallest Substring With Identical Characters II]()

* :red_circle: 预算内允许的最多数量robots: [2398. Maximum Number of Robots Within Budget](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2398-Maximum-Number-of-Robots-Within-Budget)
    * Binary Search 猜答案 + Sliding Window (Subarray Sum + Max Deque)

* :yellow_circle: 最少时间修完所有车辆: [2594. Minimum Time to Repair Cars](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2594-Minimum-Time-to-Repair-Cars)
    * 不需要考虑具体每个车床分配多少辆车，计算一个上限足以作为判定条件

* :yellow_circle: 可制造最多的合金数: [2861. Maximum Number of Alloys](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2861-Maximum-Number-of-Alloys)
    * Binary Search 猜答案
    * 注意Binary Search上界定得太高会造成溢出

* :yellow_circle: 最短可能的子串的bit OR总和>=k: [3097. Shortest Subarray With OR at Least K II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3097-Shortest-Subarray-With-OR-at-Least-K-II)
    * Binary Search 猜答案 + Sliding Window + Bitwise OR Sum

* :red_circle: 最短时间Mark所有位置II: [3049. Earliest Second to Mark Indices II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3049-Earliest-Second-to-Mark-Indices-II)
    * Binary Search 猜答案 + Regret Greedy (MinHeap)

* :red_circle: 最短时间Mark所有位置I: [3048. Earliest Second to Mark Indices I](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3048-Earliest-Second-to-Mark-Indices-I)
    * Binary Search 猜答案 + Greedy

* 抢房子4: [2560. House Robber IV](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2560-House-Robber-IV)
    * Binary Search 猜答案 + DP

* 从不能整除的元素中找最小可能的最大值: [2513. Minimize the Maximum of Two Arrays](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2513-Minimize-the-Maximum-of-Two-Arrays)

* 找最小可裁定/删除的subarray长度: [2565. Subsequence With the Minimum Score](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2565-Subsequence-With-the-Minimum-Score)
    * 结合 **3-Pass** 思想
    * 重点在于要清楚明白各物理量的物理意义

## 非常规

* :red_circle: 二维矩阵中寻找峰值II: [1901. Find a Peak Element II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1901-Find-a-Peak-Element-II)

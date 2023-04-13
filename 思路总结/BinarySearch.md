# Binary Search

## Contents
* [方法论](#方法论)
* [经典题](#经典题)
* [First/Last Occurrence](#find-first-occurrence--last-occurrence)
* [Closet element greater/less](#find-closest-element-greatersmaller-than-target)
* [BS猜答案](#binary-search-猜答案)
    * [Guess K-th Element](#guess-k-th-element)
    * [Guess + Math](#guess--math)
    * [Guess Min/Max](#guess-minmax)

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

* 什么时候需要 Post-Processing ?
    * 使用 Binary Search 时, 只有当题目的**input不一定包含答案target的时候**需要 post-processing, 否则 Binary Search 一定可以找到答案, 也就是不用post-processing
    * `lowerBound(target)` 和 `upperBound(target)` 都需要 post-processing. 因为搜索区间内不一定存在满足target的左侧边界 和 满足target的右侧边界. 
        * e.g. 搜索区间内的值都 `> target`，那就找不到 满足target的左侧边界 (ie. 第一个 `== target` 的元素)
        * e.g. 搜索区间内的值都 `<= target`, 那就找不到 满足target的右侧边界 (ie. 第一个 `> target` 的元素)

* **模版 - find lower bound & right bound**

**Assume input array non-decreasing order:**

1. `lowerBound()`: Find the **first index** whose element in the array has the value >= target (It follows the same concept as `lower_bound()` method in C++)
    1. 情况一: target 不在数组范围中，target 小于数组中所有元素，在**左边界**以外 $\Rightarrow $ 返回数组的第一个元素的index
    2. 情况二: target 不在数组范围中，target 大于数组中所有元素，在**右边界**以外 $\Rightarrow $ 返回数组的最后一个元素的index+1
    3. 情况三: target 在数组范围中，只是数组中不存在 $\Rightarrow $ 返回数组中第一个 > target的元素的index
    4. 情况四: target 在数组范围中，并且数组中存在  $\Rightarrow $ 返回数组中第一个 = target的元素的index
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

    // post-processing
    // if nums[left] == target { // 情况四
    //     return left
    // } else if nums[left] > target { // 情况一 + 情况三
    //     return left
    // } else { // 情况二
    //     return left + 1
    // }

    // post-processing 简化版
    if nums[left] >= target { // left符合题意, 直接返回left
        return left
    } else { // nums[left] < target, 返回left+1 (情况二)
        return left + 1
    }
 }
```

2. `upperBound()`: Find the **first index** whose element in the array has the value > target (It follows the same concept as `upper_bound()` method in C++)
    1. 情况一: target 不在数组范围中，target 小于数组中所有元素，在**左边界**以外 $\Rightarrow $ 返回数组的第一个元素的index
    2. 情况二: target 不在数组范围中，target 大于数组中所有元素，在**右边界**以外 $\Rightarrow $ 返回数组的最后一个元素的index+1
    3. 情况三: target 在数组范围中，只是数组中不存在 $\Rightarrow $ 返回数组中第一个 > target的元素的index
    4. 情况四: target 在数组范围中，并且数组中存在  $\Rightarrow $ 返回数组中第一个 > target的元素的index
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

    // post-processing
    // if nums[left] == target { // 情况四
    //     return left + 1
    // } else if nums[left] < target { // 情况二
    //     return left + 1
    // } else { // 情况三 + 情况一
    //     return left
    // }

    // post-processing 简化版
    if nums[left] > target { // left符合题意, 直接返回left
        return left
    } else { // nums[left] <= target, 返回left+1 (情况二+情况四)
        return left + 1
    }
}
```

* **备注**: `lowerBound()`, `upperBound()` 的物理意义也可以根据具体题目进行更改, 并根据具体的物理意义微调代码. 
    * e.g. `lowerBound()`表示搜索区间内 最大的/最后一个 小于target的元素
    * e.g. `upperBound()`表示搜索区间内 最大的/最后一个 等于target的元素


## 经典题
* Classic Binary Search: [704. Binary Search](https://leetcode.com/problems/binary-search/)


## Find First Occurrence / Last Occurrence

* 在排序数组中查找元素的第一个和最后一个位置: [34. Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

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



* 找最小的大于Target的元素: [744. Find Smallest Letter Greater Than Target](https://leetcode.com/problems/find-smallest-letter-greater-than-target/)




## Binary Search 猜答案

**算法题的解题思路可以概括为两大类：**
1. 正面硬刚：用合适的算法 (DP, DFS, BFS, etc.) 直接解决
2. 二分搜值：猜出答案, 每一轮看看是猜大了还是猜小了？以达到逼近最终答案的目的
    * 此法适用于 题目要求 return single value (scalar) 并且 涉及 max/min value

**要点**: 什么时候可以用？
* 当我们从题目中找到一种 **单调性** 的规律时，一般可以尝试用二分猜答案

**猜答案的具体步骤**
1. 确定搜索空间：也就是确定Binary Search的下界/左边界 和 上界/右边界
    * 因为是在搜索空间里猜答案，所以 左/右边界的物理意义通常就是 return值本身 (i.e. 也就是说，不是index而是value本身)
2. 确定收敛条件：什么条件让我们确定可以舍弃掉一半的空间？舍弃掉小的那一半还是大的那一半？包不包括边界的值？



**经典例题**

### Guess K-th Element

* :red_circle: Search in Sorted Matrix: [240. Search a 2D Matrix II](https://leetcode.com/problems/search-a-2d-matrix-ii/description/)
    * 解法本身不是Binary Search, 但是**解法思路可以应用于所有 sorted matrix**, 并且和 Binary Search 结合组成更高效的算法
    * 从 bottom-left / top-right 出发做 matrix traversal

* :red_circle: 无序数组中第k大的元素: [215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

* 非递减矩阵中找第k小的元素: [378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/)
    * 用 max heap 也可以解，但是 Binary Search 更效率

* 乘法表中找第k小的元素: [668. Kth Smallest Number in Multiplication Table](https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/description/)

* 找k个最小和: [373. Find K Pairs with Smallest Sums](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/)
    * 用 max heap 会超时！！！只能 Binary Search 提高效率

* 非递减数组中找第k小的乘积: [2040. Kth Smallest Product of Two Sorted Arrays](https://leetcode.com/problems/kth-smallest-product-of-two-sorted-arrays/description/)
    * 找前k小的个数的思路与前面题型不同. main idea: 分类讨论

* :red_circle: :secret: 从矩阵中找第k小 array sum: [1439. Find the Kth Smallest Sum of a Matrix With Sorted Rows](https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/description/)
    * 找前k小的个数的思路与前面题型不同. main idea: 结合**DFS**
    * 用 **DFS - All Combo Sum** 来找满足当前猜的数 (`mid`) 的array 的个数是否 >= k 个

* 找第k小 subarray sum: [1918. Kth Smallest Subarray Sum](https://leetcode.ca/2021-08-01-1918-Kth-Smallest-Subarray-Sum/)
    * 找前k小的个数的思路与前面题型不同. main idea: 转化 subarray sum 为 prefixSum diff, 再结合**Two Pointers**
    * 因为 prefixSum array 是单增的, 所以可以通过 **Two Pointers** 来找满足当前猜的数 (`mid`) 的 subarray 个数是否 >= k 个

### Guess + Math

* :star: 搜索区间为自然数，由于自然数递增排序，所以可以天然地使用二分搜索

* 寻找丑数III: [1201. Ugly Number III](https://leetcode.com/problems/ugly-number-iii/description/)
    * *1201 与 1539 思路相近*
    * 需要用到一些数学定理: 容斥原理, LCM, GCD

* 找第k个缺失的自然数: [1539. Kth Missing Positive Number](https://leetcode.com/problems/kth-missing-positive-number/description/)
    * * *1201 与 1539 思路相近*

* 阶乘函数后 K 个零: [793. Preimage Size of Factorial Zeroes Function](https://leetcode.com/problems/preimage-size-of-factorial-zeroes-function/description/)

### Guess Min/Max

* 抢房子4: [2560. House Robber IV](https://leetcode.com/problems/house-robber-iv/description/)
    * Binary Search 猜答案 + DP

* 从不能整除的元素中找最小可能的最大值: [2513. Minimize the Maximum of Two Arrays](https://leetcode.com/problems/minimize-the-maximum-of-two-arrays/description/)

* 找最小可裁定/删除的subarray长度: [2565. Subsequence With the Minimum Score](https://leetcode.com/problems/subsequence-with-the-minimum-score/)
    * 结合 **3-Pass** 思想
    * 重点在于要清楚明白各物理量的物理意义
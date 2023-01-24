# Binary Search

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

## 经典题
* Classic Binary Search: [704. Binary Search](https://leetcode.com/problems/binary-search/)




## Find First Occurrence / Last Occurrence

[34. Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

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

* **模版**

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

1. `upperBound()`: Find the **first index** whose element in the array has the value > target (It follows the same concept as `upper_bound()` method in C++)
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

* 找最小的大于Target的元素: [744. Find Smallest Letter Greater Than Target](https://leetcode.com/problems/find-smallest-letter-greater-than-target/)




## Binary Search 猜答案

**算法题的解题思路可以概括为两大类：**
1. 正面硬刚：用合适的算法 (DP, DFS, BFS, etc.) 直接解决
2. 二分搜值：猜出答案, 每一轮看看是猜大了还是猜小了？以达到逼近最终答案的目的
    * 此法适用于 题目要求 return single value (scalar) 并且 涉及 max/min value

**猜答案的具体步骤**
1. 确定搜索空间：也就是确定Binary Search的下界/左边界 和 上界/右边界
    * 因为是在搜索空间里猜答案，所以 左/右边界的物理意义通常就是 return值本身 (i.e. 也就是说，不是index而是value本身)
2. 确定收敛条件：什么条件让我们确定可以舍弃掉一半的空间？舍弃掉小的那一半还是大的那一半？包不包括边界的值？

**经典例题**

### Guess K-th Element

* Search in Sorted Matrix: [240. Search a 2D Matrix II](https://leetcode.com/problems/search-a-2d-matrix-ii/description/)
    * 解法本身不是Binary Search, 但是**解法思路可以应用于所有 sorted matrix**, 并且和 Binary Search 结合组成更高效的算法
    * 从 bottom-left / top-right 出发做 matrix traversal

* 非递减矩阵中找第K小的元素: [378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/)
    * 用 max heap 也可以解，但是 Binary Search 更效率

* 乘法表中找第K小的元素: [668. Kth Smallest Number in Multiplication Table](https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/description/)

* 找k个最小和: [373. Find K Pairs with Smallest Sums](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/)
    * 用 max heap 会超时！！！只能 Binary Search 提高效率


### 其他
* 从不能整除的元素中找最小可能的最大值: [2513. Minimize the Maximum of Two Arrays](https://leetcode.com/problems/minimize-the-maximum-of-two-arrays/description/)
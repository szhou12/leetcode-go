# [925. Long Pressed Name](https://leetcode.com/problems/long-pressed-name/description/)

## Solution idea

### Two Pointers - 两个数组双指针同向而行

* **这道题有很多坑**
    * 看题意好像找到common subsequence就行，但实际上不是！！！
    * `typed` 里只能出现 `name`中出现的字母，并且相对顺序要与`name`中一致
        * Case 1: `name=alex`, `typed=aaleelx` 这种就不行因为`l`的顺序没有保持一致，第二个`l`出现在了 `e` 后头
        * Case 2: `name=alex`, `typed=aaleexdd` 这种也不行因为`d`没有在`name`中出现过 (虽然在最后也不行)

* 解法：
    * 两个数组双指针`i, j`同向而行
    * 双指针`i, j`指向的元素每一轮必须随时保持相同，直到双指针都走完了各自的数组
    * 如果当前轮`i, j`指向的元素相同，`i, j`同时右移直到指向不同元素；此时`j`右移跳过所有重复的元素直到指向不同元素；进入下一轮
    * 如果下一轮`i, j`指向了不相同元素，直接返回false (出现Case 1); 否则，重复上一步
    * 最后，如果 `i, j`没有全部走完各自的数组，返回false (出现Case 2)


Time complexity = $O(m+n)$

## Resource

[halfrost/LeetCode-Go](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0925.Long-Pressed-Name)

[代码随想录-925.长按键入](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0925.%E9%95%BF%E6%8C%89%E9%94%AE%E5%85%A5.md)
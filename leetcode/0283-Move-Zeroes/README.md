# [283. Move Zeroes](https://leetcode.com/problems/move-zeroes/)

## Solution idea

### Two Pointers - 同向而行
**注意**题意要求保留相对顺序，所以双指针相向而行不可行

* 快慢指针同向而行

* 物理意义:
    * **left 挡板的左边**都是非零元素
    * **left 指针**永远逗留(指向)0 (除了初始位置可能不是0)

* 机制:
    * right指针不停向前走
        * 如果right指针当前指向非0元素，和left指针互换元素; left往前走一格
        * 如果right指针当前指向0元素，什么都不用管，进入下一轮

Time complexity = $O(n)$
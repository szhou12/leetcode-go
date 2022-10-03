# [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)

## Solution idea

1. Step 1: 同 `LC-141`. 双指针同向而行, fast指针走两步，slow指针走一步. 判断：1. 是否有环. 2. 若有环，fast和slow的相遇节点

2. Step 2: fast指针从头节点开始，slow指针从 *Step 1*的相遇节点开始。双指针同向而行, fast指针走一步，slow指针走一步. fast和slow的相遇节点即为环的起始节点。

* Why?
    * 设：头节点到环的起始节点的节点数(路程)为 `x`, 环的起始节点到*Step 1*相遇节点的节点数(路程)为 `y`, *Step 1*相遇节点到环的起始节点的节点数(路程)为 `z`
    * 在*Step 1*中：
        * **Key 1**: fast 走到相遇节点花的路程数为：`x+y+n(y+z)`, $n \geq 1$.
        * **Key 2**: slow 走到相遇节点花的路程数为：`x+y`.
        * **Key 3**: fast 的速度为 2， slow 的速度为 1.
        * **Key 4**: fast 和 slow 会相遇，说明他们走的时间 t 是相同的
        * 由以上 4 点，我们可以推导出 fast 和 slow 的关系：
        ```
        fast: t = (x+y+n(y+z)) / 2
        slow: t = (x+y) / 1
        => (x+y+n(y+z)) / 2 = (x+y) / 1
        => x+y+n(y+z) = 2(x+y)
        => n(y+z) = x+y
        => (n-1)(y+z) + z = x
        => z = x for n == 1; z = x for n > 1 as well bc (y+z) is whole cycle
        ```
    * 所以，在*Step 2*中，只需要一个指针从头节点走，另一个指针从*Step 1*相遇节点走，两人同时以相同的速度走，相遇的地方就是环的起始节点

Time comlexity = $O(n)$

## Resource

[代码随想录-142.环形链表II](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0142.%E7%8E%AF%E5%BD%A2%E9%93%BE%E8%A1%A8II.md)

[halfrost-142.环形链表II](https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/0142.Linked-List-Cycle-II)

# [649. Dota2 Senate](https://leetcode.com/problems/dota2-senate/)

## Solution idea

### 模拟: two queues storing index of senates
* 两个Queue (R, D), 分别存R议员的index, D议员的index
* 每一轮：
    * 两个Queue分别pop头一个元素, 谁的index值小, 说明该议员排位靠前, 可以ban掉另一个敌方议员
    * 排位靠前议员存活，可以进入下一轮，排到队尾, **index 变成 index + n**
* 只要有一个queue空了，说明这一队所有人都被消灭; 结束loop, 此时，哪一个queue不为空，即是赢家

Time complexity = $O(n)$, Space complexity = $O(n)$

### Greedy
* 难以想到，并且难以实现
* **Idea**: 消灭的策略是，尽量消灭自己后面的对手，因为前面的对手已经使用过权利了，而后序的对手依然可以使用权利消灭自己的同伴

Time complexity = $O(n)$

## Resource
* 模拟-两个queue存index, 容易理解: [贾考博 LeetCode 649. Dota2 Senate](https://www.youtube.com/watch?v=GqNlKwelclc)

* Greedy: [代码随想录-649. Dota2 参议院](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0649.Dota2%E5%8F%82%E8%AE%AE%E9%99%A2.md)
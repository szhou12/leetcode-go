# [1950. Maximum of Minimum Values in All Subarrays](https://leetcode.ca/2021-08-11-1950-Maximum-of-Minimum-Values-in-All-Subarrays/)

## Solution idea
### Stack - prevSmaller, nextSmaller + MaxHeap
1. 题目要求：给定一个定长`k` (`1 <= k <= n`)，从`nums[]`中找到符合该定长的所有subarrays，并从每个subarray中挑出最小值，然后从这群最小值中挑出最大的那个。
2. 首先，容易想到的是：subarray minimums。即，使用递增栈，找到每个元素的 prevSmallerOrEqual 和 nextSmaller。注意，这道题是需要去重的，所以是 prevSmallerOrEqual。
3. 既然我们已经知道一个元素作为最小值时的subarray最长有多长`L`，那么很自然地，可以想到，我们拿这个长度跟`k`比较：如果`L >= k`，那么这个元素是一个合格的候选人，因为我们可以保证在长度`k`时，我们可以找到这样的一个subarray以当前这个元素为最小值 (当然，我们尚未知道这个元素是否是“鹤立鸡群”的那个)。相反，如果 `L < k`，那么，显然这个元素压根没法入选。
4. 那么，剩下要解决的问题就是如何在众多候选人中找到那个最大值？这时就需要盘一个逻辑：
    - `k`越小时，全局越大的值越容易成为这个“鹤立鸡群的人”。比如：当`k=1`时，每个subarray都是单一元素，这时**全局最大值**就成为了我们要找的那个人。当`k=2`时，每个subarray都只有两个元素，这时，**全局第二大**有可能成为这个“鹤立鸡群的人” (如果全局第二大刚好挨着全局最大值)，或者一个更小的值成为“鹤立鸡群的人” (如果全局第二大不挨着全局最大值，也就是说它挨着比它更小的值)。以此类推，我们可以发现一个**单调性**。
    - 当`k`变大时，“鹤立鸡群的人”的值就变得越小。这是因为`k`越大的subarray，越容易包含更多较小元素，使得subarray的最小值(候选人)变得越小，从而时最后的结果(“鹤立鸡群的人”)变小。
    - 具体实现：我们用一个MaxHeap来存每个元素的值和index。我们逐步从小到大增长`k=1...n`，拿出MaxHeap栈顶的这个元素，看是否存在以它作为**最小值**的subarray，并且这个subarray长度`>= k`：
        - 如果存在，显然这个元素就是我们要找的当前`k`长度下“鹤立鸡群的人”（首先，全局没有比它更大的了，其次，它还是某个subarray的最小值-势必会被选为候选人）。这时，我们不急着把它排出MaxHeap，因为它还有可能在下一次`k+1`时，再度入选。我们只是记录这个答案。
        - 如果不存在，即以它做最小值的subarray `< k`，那么，首先，它在当前轮就没法入选，其次，当`k`增长时，它更没法入选。这时，我们就可以把它从MaxHeap中排出。排出后，再看此时栈顶元素是否满足要求。
        - 以这个流程，我们就可以头也不回地找出，`k`从小变到大时，每一轮的“鹤立鸡群的人”。

Time complexity = $O(n\log n)$

## Resource
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Stack/1950.Maximum-of-Minimum-Values-in-All-Subarrays)
# [1358. Number of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/description/)

## Solution idea
### Two Pointers
#### 思路总结
1. 本题求满足条件的 substring 个数，很明显属于 Two Pointers 中的 *Sliding Window 长度可变 == Subarray 类型* :arrow_right: 使用 模版Flex
2. 两个难点 (我在解题时, 导致bug的两个点):
    1. 每次right停下来时, substring 个数怎么数？
        * 需要明确 right 停下来时，滑窗的物理意义是什么：此时，滑窗内表示a, b, c每个字母至少包含一个，换句话说，这表示锚定当前left，right直至末尾的所有substring都满足条件。所以，substring 个数 = (n-1) - (right-1) + 1
        ```
        [X   X   X   X]   X   X   X   X   X   X
                     --------------------------
        left           right                 n-1
        ```
        * 我一开始写的是 `个数+=1`，相当于只算了滑窗圈起来的这一个substring，肯定少算了
    2. 每次right停下来，需要有条件地 update result
        * right停下来的条件可以是: 1. right越界了；2. 滑窗内条件满足了
        * 需要注意的是，right越界时，并不一定滑窗内条件满足了，所以不能直接update result
        * 所以，需要check此时滑窗内条件是否满足，只有满足时才update result

Time complexity = $O(n)$

## Resource
[【每日一题】LeetCode 1358. Number of Substrings Containing All Three Characters](https://www.youtube.com/watch?v=NimL6uuhBYI&ab_channel=HuifengGuan)
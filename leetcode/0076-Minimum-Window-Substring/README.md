# [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/description/)

## Solution idea

### Sliding Window
* 对于每个吃进来的元素`s[right]`，首先更新该字符出现次数的 `window[s[right]]++`。如果更新后，`window[s[right]]` 等于需要出现的次数 `need[s[right]]`，则计数器 `count++`，说明有一个字符满足了出现次数的要求.
* 当 `count` 等于 `t` 中的字符类型数 `len(need)` 时，说明任务已经实现。此时，让 `left`指针不断右移，相应的 `window[s[left]]` 就要自减，一旦 `window[s[left] < need[s[left]]`，则 `count` 需要自减1从而不再满足 `len(need)`，说明需要继续加入新元素才能满足任务. 从而 `right` 才可以右移继续遍历。
* 在这个过程中，每次满足条件 `count==len(need)`，都需要不断更新和记录结果。

## Resource
[滑动窗口算法](https://labuladong.github.io/algo/2/20/27/)
[wisdompeak/LeetCode](https://github.com/wisdompeak/LeetCode/tree/master/Two_Pointers/076.Minimum-Window-Substring)
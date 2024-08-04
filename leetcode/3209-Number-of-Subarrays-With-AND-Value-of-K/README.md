# [3209. Number of Subarrays With AND Value of K](https://leetcode.com/problems/number-of-subarrays-with-and-value-of-k/description/)

## Solution idea
### Bitwise AND + HashTable (Suffix Hashing)
1. Bitwise AND的性质：AND越多，结果越小。单调递减。Non-increasing.
    - 如何利用Bitwise AND的性质？注意到Constraint中的`nums[i] <= 10^9`，而`10^9`差不多是`2^31`。也就是说，每个`nums[i]`的二进制表示最多有31位全是1。此时，结合Bitwise AND的性质，我们可以发现，从全是1降到全是0，一共会经历31次变换 (每一位`1 -> 0`，总共31位。i.e., 有一个0，有两个0，... 有31个0)。
2. 这时候，你可能会想：等一下！相同数量的0在不同bit位的可能性为什么不算了？比如说，4-bit with one 0: `1110`, `1101`, `1011`, `0111`。这不是一下子就有四种可能吗？
3. 你再仔细想想，在这道题的语境下，有没有所有上述情况都出现的可能？其实不可能，因为要结合 "subarray" 这个关键词。对于所有以`nums[i]`结尾的subarrays: `nums[i-1 : i], ..., nums[1:i], nums[0:i]`，subarray越长，bitwise AND的数越多，那么，最后结果里的0的数量只可能越来越多。唯一的例外是0数量不变的情况，只有在往前AND的时候，前面的数都等于`nums[i]`时会发生。e.g. `1110 & 1110 = 1110`。只要往前“并”时，数不同于`nums[i]`，那么，这个subarray的AND结果中0的数量一定是越来越多的。也就是说，不可能发生 `X & 1110 = 1101`这种例子。
4. 总结：对于所有以`nums[i]`结尾的subarrays，无论它最长有多少元素，这些subarrays的bitwise AND的数值种类最多只有31种。
5. 实现：用一个map来存所有以`nums[i-1]`结尾的subarrays的AND结果的种类数。key是AND结果 (最多31种)，value是该结果出现次数count。在更新以`nums[i]`结尾的subarrays时，我们只需要遍历map里的key，都与`nums[i]`进行AND，得到的新"key"累加count。注意，`nums[i]`单一元素也可以是一个subarray，所以，`nums[i]`本身也是一个key，对应count要累加1。遍历完以后，只要检查有没有key=k，有的话，说明有以`nums[i]`结尾的subarray的AND结果是k，累加对应的count到最终结果里。

Time complexity = $O(31* n) = O(n)$

## Resource
[【每日一题】LeetCode 3209. Number of Subarrays With AND Value of K](https://www.youtube.com/watch?v=MPDkaH9x9Yo&ab_channel=HuifengGuan)
# [454. 4Sum II](https://leetcode.com/problems/4sum-ii/)

## Solution idea

### My Solution - break into 2 two-sums
1. 两两数组先找出所有可能的sum
    * 这里任意两个数组pair up都可以，因为加法满足交换律
2. 此时，四个数组变成了两个数组 A, B. 然后再用常规的 `2Sum` 解，即，数组B转换成 map, 然后看 Target 与 数组A每个元素的差 是否是map的key.
如果是, 则找到一组配对，把key对应的value加进result.
    * 注意：这里 `result += value` 而不是 `result += 1`, 这是因为 Step 1 找出的sum有可能重复，但是重复的应该当作distinct elements因为构成他们的元素index并不一样

Time complexity = $O(n^2)$

Step 1 = $O(n^2)$

Step 2 = $O(n^2)$ 因为两两数组所有可能的sum个数最多 $n^2$ 个 distinct sum

## Resource

[代码随想录-第454题.四数相加II](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0454.%E5%9B%9B%E6%95%B0%E7%9B%B8%E5%8A%A0II.md)

Follow-up: 如果本题想难度升级：就是给出一个数组（而不是四个数组），在这里找出四个元素相加等于0，答案中不可以包含重复的四元组。

# Prefix Sum & Difference Array

## Difference Array (差分数组)

* 差分数组: `diff[i] = nums[i] - nums[i-1]`, `diff[0] = nums[0]`

* **规律1** 对差分数组求第i个元素的前缀和:
```
sum[i] = diff[0] + diff[1] + ... + diff[i]
       = (nums[0]) + (nums[1] - nums[0]) + ... + (nums[i] - nums[i-1])
       = nums[i]
```

* **规律1总结**: 差分数组第i个元素的前缀和 = 原数组第i个元素的值

* **规律2** 对原数组`nums[l...r]` (双闭区间) 每个元素都加上 `val` 时，映射到差分数组上就是: `diff[l]+val` 和 `diff[r+1]-val`

* **规律2总结**:
    * 思路: 对原数组`nums`的某个subarray`nums[l...r]`的每个元素都加上一个值 `val`, 实际上是对差分数组的 `diff[l]+val` 和 `diff[r+1]-val`
    * 结果: 对变化后的差分数组`diff`求每个元素的前缀和 $\Rightarrow $ 变化后的原数组

* 举个例子:
```
nums: 1,   3,  5,  4,  8  => nums[0...1] += 10
     11,  13,  5,  4,  8
diff: 1,   2,  2, -1,  4
     11,   2, -8, -1,  4  => find that only diff[0], diff[1+1] changes, diff[0] += 10, diff[1+1] -= 10 
```
# Prefix Sum & Difference Array

## 目录
* [一维差分](#一维差分-1-d-difference-array)
* [二维差分](#二维差分-2-d-difference-array)
* [扫描线](#扫描线-sweep-line)
* [整体区间的增减](#整体区间的增减)

## Prefix Sum (前缀和)
* 前缀和的定义:

```
prefixSum[0] = nums[0],
prefixSum[i] = nums[0] + nums[1] + ... + nums[i]
```

## Difference Array (差分数组)

### 1-D Difference Array

* 差分数组的定义:

```
diff[0] = nums[0] - 0,
diff[i] = nums[i] - nums[i-1]
```

* **性质 1**: 差分数组中求第i个元素的前缀和:
    * **性质 1 总结**: 差分数组第i个元素的前缀和 = 原数组第i个元素的值
```
presum[i] = presum[i-1] + diff[i]
          = diff[0]   + diff[1]             + ... + diff[i]
          = (nums[0]) + (nums[1] - nums[0]) + ... + (nums[i] - nums[i-1])
          = nums[i]

presum[0] = diff[0]
```

* **性质 2**: 对原数组`nums[l...r]` (双闭区间) 每个元素都加上 `val` 时，映射到差分数组上就是: `diff[l]+val` 和 `diff[r+1]-val`
    * **性质 2 总结**:
        * 思路: 对原数组`nums`的某个区间/subarray`nums[l...r]`的每个元素都加上一个值 `val`, 实际上是对差分数组的 `diff[l]+val` 和 `diff[r+1]-val`
        * 结果: 对变化后的差分数组`diff`求每个元素的前缀和 $\Rightarrow $ 变化后的原数组
    * 举个例子:
```
idx:   0,   1,  2,  3,  4
nums:  1,   3,  5,  4,  8  => 对 nums[0...1] + 10
   => 11,  13,  5,  4,  8
diff:  1,   2,  2, -1,  4
      11,   2, -8, -1,  4  => 只有 diff[0], diff[1+1] 发生变化: diff[0] += 10, diff[1+1] -= 10 
```

* **实现细节**
    1. 由**性质 1**可知，实现代码中需要开辟两个数组`diff[]`和`presum[]`，`diff[]`是差分数组，`presum[]`用来计算**差分数组的前缀和**。
    2. 在计算差分数组时，要时刻注意越界的情况。e.g. `diff[r+1]-val`中的 `r+1`当 `r`指向原数组最后一个元素时就会发生越界。
    3. 一般长度为`n`的原数组`nums[]`，生成长度为`n+1`的`diff[]`和`presum[]`。
        1. 方便规避index越界的情况
        2. 方便计算`diff[0]`


## 一维差分 (1-D Difference Array)

* 航班预订统计: [1109. Corporate Flight Bookings](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1109-Corporate-Flight-Bookings)

* 拼车: [1094. Car Pooling](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1094-Car-Pooling)

## 二维差分 (2-D Difference Array)

* 子矩阵加1: [2536. Increment Submatrices by One](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2536-Increment-Submatrices-by-One)

## 扫描线 (Sweep Line)

* :red_circle: 统计不开会的“好日子”: [3169. Count Days Without Meetings]()

* :red_circle: "好"分隔的总数: [2963. Count the Number of Good Partitions](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2963-Count-the-Number-of-Good-Partitions)

## 整体区间的增减

* :red_circle: 是否可以消掉所有元素为0: [2772. Apply Operations to Make All Array Elements Equal to Zero](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2772-Apply-Operations-to-Make-All-Array-Elements-Equal-to-Zero)
# [1356. Sort Integers by The Number of 1 Bits](https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/description/)

## Solution idea

### Naive solution
* convert a number to its binary representation
* sort by number of 1's

### Bit Operation
* 循环n的二进制表示中1的个数次
```
int bitCount(int n) {
    int count = 0;
    while (n) {
        n &= (n - 1); // 清除最低位的1
        count++;
    }
    return count;
}
```

## Resource
[代码随想录-1356. 根据数字二进制下 1 的数目排序](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/1356.%E6%A0%B9%E6%8D%AE%E6%95%B0%E5%AD%97%E4%BA%8C%E8%BF%9B%E5%88%B6%E4%B8%8B1%E7%9A%84%E6%95%B0%E7%9B%AE%E6%8E%92%E5%BA%8F.md)
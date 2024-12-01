# HashTable

## 目录
* [Two Sum](#Two-Sum)
* [Prefix Hashing](#Prefix-Hashing)

## Two Sum
* yellow_circle: 找出数组中最大的outlier: [3371. Identify the Largest Outlier in an Array]()

* :red_circle: 能组成24倍数的配对数: [3185. Count Pairs That Form a Complete Day II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3185-Count-Pairs-That-Form-a-Complete-Day-II)
    - Two Sum + HashMap = {Remainder : Frequency}

* :red_circle: 两数之和: Two Sum: [1. Two Sum](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0001-Two-Sum)
    - 法一: HashMap
    - 法二: Two Pointers

## Prefix Hashing
* :red_circle: 多少子数组的bitwise AND结果是k: [3209. Number of Subarrays With AND Value of K](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3209-Number-of-Subarrays-With-AND-Value-of-K)
    * AND性质: 连续 bitwise AND 所得值一定越变越小。
    * 常数个key (31种)
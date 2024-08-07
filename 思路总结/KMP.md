# KMP

## 目录
* [应用场景](#应用场景)
* [算法概述](#算法概述)
* [一图胜千言](#一图胜千言)
* [题目](#题目)

## 应用场景
1. 在一个字符串 (haystack) 中找到一个子串 (needle) **首次出现的位置** 或者 **所有出现的位置**。

## 算法概述
* 模版：
    1. Find first appearance (index) of a substring in a string: [28. Find the Index of the First Occurrence in a String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0028-Implement-strStr)
    2. Find all appearances (indices) of a substring in a string: [3008. Find Beautiful Indices in the Given Array II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3008-Find-Beautiful-Indices-in-the-Given-Array-II)
* 实现逻辑以Guan神的讲解为最优先
* 其他资源：[动态规划之 KMP 算法详解](https://mp.weixin.qq.com/s/r9pbkMyFyMAvmkf4QnL-1g)


## 一图胜千言
![kmp](https://github.com/szhou12/leetcode-go/assets/35708194/2f038b22-54c2-4ecc-bade-19769da60cec)

## 题目

* :red_circle: 前后缀同时匹配的配对数: [3045. Count Prefix and Suffix Pairs II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3045-Count-Prefix-and-Suffix-Pairs-II)
    * KMP STEP 1 + Trie

* :red_circle: 最少时间恢复字符串原样II：[3031. Minimum Time to Revert Word to Initial State II](https://leetcode.com/problems/minimum-time-to-revert-word-to-initial-state-ii/description/)
    * 考察KMP算法中重要的传递性(transitive)思想

* :red_circle: 最少时间恢复字符串原样I：[3029. Minimum Time to Revert Word to Initial State I](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3029-Minimum-Time-to-Revert-Word-to-Initial-State-I)
    * 暴力解法：直接比较在最少切几刀的情况下，使得后缀 = 前缀.

* :red_circle: 找最长自相关最长前缀后缀数组: [1392. Longest Happy Prefix](https://github.com/szhou12/leetcode-go/tree/main/leetcode/1392-Longest-Happy-Prefix)
    * **优先练习** 考察KMP算法 Step 1 的实现

* :red_circle: 找第一次出现的匹配: [28. Find the Index of the First Occurrence in a String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0028-Implement-strStr)
    * **优先练习** 考察KMP算法 Step 1 + 2 的实现
    * 字符串中找子串**第一次出现**的位置

* :red_circle: 找所有的“美丽”index II: [3008. Find Beautiful Indices in the Given Array II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3008-Find-Beautiful-Indices-in-the-Given-Array-II)
    * KMP: 字符串中找子串**所有出现**的位置
    * Binary Search: `upperBound()` 和 `lowerBound()`

* :yellow_circle: 找所有的“美丽”index I: [3006. Find Beautiful Indices in the Given Array I](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3006-Find-Beautiful-Indices-in-the-Given-Array-I)
    * 与 3008 完全一样

* :purple_circle: 有限次切字符串找匹配: [2851. String Transformation](https://github.com/szhou12/leetcode-go/tree/main/leetcode/2851-String-Transformation)

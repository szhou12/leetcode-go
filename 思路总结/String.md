# String

## 目录
* [翻转String](#翻转String)
* [Rotate/Circular String](#Rotate/Circular-String)
* [Compress String](#Compress-String)


## 翻转String
 * [151. Reverse Words in a String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0151-Reverse-Words-in-a-String)

 * [189. Rotate Array](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0189-Rotate-Array)

 * [541. Reverse String II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0541-Reverse-String-II)
	* 每 `2*k` 一跳，每一个 `2*k`区间内翻转 前`k` 个字符

## Rotate/Circular String
### 技巧: 复制一倍. 同一个String/Array拼接在一起 `s+s`

* :green_circle: 构建加密字符串: [3210. Find the Encrypted String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/3210-Find-the-Encrypted-String)
    * Modulo + string复制一倍

* Rotate题: [796. Rotate String](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0796-Rotate-String)
	* `source str + source str`, 那么所有 rotated string 就都会被包含在里面
	* The 'shift' operation can be regarded as a 'sliding' operation on `source str + source str`

* Circular Array: [503. Next Greater Element II](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0503-Next-Greater-Element-II)
	* 思路: 从后往前看，挤掉栈顶矮个子
    * 用 取模 (`%`) 的方法 **模拟** **复制一倍**的方法

## Compress String
* 字符串压缩: [443. String Compression](https://github.com/szhou12/leetcode-go/tree/main/leetcode/0443-String-Compression)
    * **Two Pointers**题，但是细节非常多
    * slow物理意义：slow挡板左边不包括slow表示compressed string
    * 遇到相同字母时，移动fast，slow不动，并计算遇到的相同字母个数
    * 遇到不同字母时：
		1. 把当时在看的字母 抛给 slow, slow走一格
		2. 把字母个数append到后面挨着 (count==1 则跳过)，slow走count多个格子. 此时，slow和fast在同一个格子

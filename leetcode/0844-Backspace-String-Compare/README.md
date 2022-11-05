# [844. Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/description/)

## Solution idea

### Stack - 对称匹配类型题
* 如果是一般字母，入栈
* 如果是 #, 栈顶元素出栈
* 最后栈内剩下的元素就是做完backspace得到的字符串

Time complexity = $O(m+n)$, Space complexity = $O(m+n)$

## Resource
[代码随想录-844.比较含退格的字符串](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0844.%E6%AF%94%E8%BE%83%E5%90%AB%E9%80%80%E6%A0%BC%E7%9A%84%E5%AD%97%E7%AC%A6%E4%B8%B2.md)


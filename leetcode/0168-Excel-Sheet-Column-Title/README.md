# [168. Excel Sheet Column Title]()

## Solution idea

26进制：模26得一个字母，并从末位开始添加；每添加一轮，除26砍掉添加过的字母

Time complexity = $O(n/26) = O(n)$ where $n$ is `columnNumber`
# [93. Restore IP Addresses](https://leetcode.com/problems/restore-ip-addresses/)

## Solution idea

### DFS

```
# levels = 4
# branches = 3
```

level 表示把整个 `s` 分割成四份

每个branch 表示取 一位数，两位数，三位数

**注意**：开头为0的两位/三位数不合法 eg. 01, 012

Time complexity = O((n-1) choose 3) where n is length of `s`
package leetcode

import "strconv"

// DFS - My version
// # levels = 4
// # branches = 3
func restoreIpAddresses(s string) []string {
	var res []string
	ip := ""

	dfs(s, 0, 0, ip, &res)
	return res
}

func dfs(s string, index int, level int, ip string, res *[]string) {
	// Base Case: 如果成功分割成四份 并且 所有数字都考虑进去了 -> 找到一个可行解
	if level == 4 && index == len(s) {
		*res = append(*res, ip[:len(ip)-1])
		return
	}

	for i := 1; i < 4; i++ {
		if index+i <= len(s) && valid(s[index:index+i]) {
			ip = ip + s[index:index+i] + "."
			dfs(s, index+i, level+1, ip, res)
			ip = ip[:len(ip)-i-1]
		}
	}
}

func valid(s string) bool {
	// 注意：开头为0的两位/三位数不合法 eg. 01, 012
	val, _ := strconv.Atoi(s)

	if 0 <= val && val <= 9 && len(s) == 1 {
		return true
	}

	if 10 <= val && val <= 99 && len(s) == 2 {
		return true
	}

	if 100 <= val && val <= 255 && len(s) == 3 {
		return true
	}
	return false
}

//           101023
//  /           |            \
// 1|01023    10|1023     101|023
//   /|\
// 0|1023

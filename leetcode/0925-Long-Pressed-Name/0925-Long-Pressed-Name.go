package leetcode

func isLongPressedName(name string, typed string) bool {
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		for i < len(name) && j < len(typed) && name[i] == typed[j] {
			i++
			j++
		}
		for j < len(typed) && typed[j-1] == typed[j] {
			j++
		}
	}

	return i == len(name) && j == len(typed)

}

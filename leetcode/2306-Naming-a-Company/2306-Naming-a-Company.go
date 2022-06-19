package leetcode

// Time = O(26^2 * n)
func distinctNames(ideas []string) int64 {
	head2str := make([]map[string]bool, 26)
	for i := 0; i < len(head2str); i++ {
		head2str[i] = make(map[string]bool)
	}
	for _, idea := range ideas {
		head2str[idea[0]-'a'][idea[1:]] = true
	}

	var res int64
	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {
			x := int64(len(head2str[i]))
			y := int64(len(head2str[j]))
			common := int64(0)
			for key, _ := range head2str[i] {
				if head2str[j][key] {
					common++
				}
			}
			res += (x - common) * (y - common) * 2

		}

	}
	return res
}

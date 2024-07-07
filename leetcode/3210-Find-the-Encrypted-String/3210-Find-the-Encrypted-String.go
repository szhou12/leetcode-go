package leetcode

func getEncryptedString(s string, k int) string {
	n := len(s)

	s2 := s + s

	j := k % n

	return s2[j : j+n]
}

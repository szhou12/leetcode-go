package leetcode

func firstUniqChar(s string) int {
    idxMap := make(map[int]int) // {charInt : count}
    for _, char := range s {
        charInt := int(char - 'a')
        idxMap[charInt]++
    }

    for i, char := range s {
        charInt := int(char-'a')
        if idxMap[charInt] == 1 {
            return i
        }
    }
    return -1
}
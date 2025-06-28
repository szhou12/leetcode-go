package leetcode

func isValidSudoku(board [][]byte) bool {
    n := 9
    rows := make([]map[byte]bool, n)
    cols := make([]map[byte]bool, n)
    boxes := make([]map[byte]bool, n)
    for i := 0; i < n; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        boxes[i] = make(map[byte]bool)
    }

    for r := 0; r < n; r++ {
        for c := 0; c < n; c++ {
            val := board[r][c]
            if val == '.' {
                continue
            }

            // check if row has val already
            if rows[r][val] {
                return false
            }
            rows[r][val] = true

            // check if col has val already
            if cols[c][val] {
                return false
            }
            cols[c][val] = true

            // check if box has val already
            idx := (r/3)*3 + (c/3)
            if boxes[idx][val] {
                return false
            }
            boxes[idx][val] = true
        }
    }

    return true
}
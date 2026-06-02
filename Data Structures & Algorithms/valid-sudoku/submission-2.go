func isValidSudoku(board [][]byte) bool {

	// Создаём map-ы, в которых будем хранить нормальные строки

	var rows [9]map[byte]bool
	var cols [9]map[byte]bool
	var tables [3][3]map[byte]bool

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tables[i][j] = make(map[byte]bool)
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			elem := board[i][j]
			if elem == '.' {
				continue
			}

			boxR := i / 3
			boxC := j / 3

			if rows[i][elem] || cols[j][elem] || tables[boxR][boxC][elem] {
				return false
			}
			rows[i][elem] = true
			cols[j][elem] = true
			tables[boxR][boxC][elem] = true
		}
	}
	return true

}
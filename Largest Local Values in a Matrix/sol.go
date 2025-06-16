func largestLocal(grid [][]int) [][]int {
	sum := 0
	rows := len(grid) - 2
	cols := len(grid[0]) - 2
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}
	//IX, IY := 0, 0
	for Y := 0; Y < len(grid); Y++ {
		for X := 0; X < len(grid[Y]); X++ {
			if Y != 0 && Y != len(grid[Y])-1 && X != 0 && X != len(grid[Y])-1 {
				sum = max(grid[Y-1][X-1], grid[Y-1][X], grid[Y-1][X+1], grid[Y][X-1], grid[Y][X], grid[Y][X+1], grid[Y+1][X-1], grid[Y+1][X], grid[Y+1][X+1])
				result[Y-1][X-1] = sum
			}

		}

	}

	return result

}
package main

// recursively calculate pool size
func Pool(heightmap [][]int, row, col, sum int) ([][]int, int) {
	if heightmap[row][col] == -1 || heightmap[row][col] == 9 {
		return heightmap, sum
	}

	//-1 means processed
	old := heightmap[row][col]
	heightmap[row][col] = -1
	sum++

	hm := HeightMap{heightmap}
	r := row - 1
	c := col
	if hm.ValidPos(r, c) && heightmap[r][c] > old {
		heightmap, sum = Pool(heightmap, r, c, sum)
	}

	r = row + 1
	if hm.ValidPos(r, c) && heightmap[r][c] > old {
		heightmap, sum = Pool(heightmap, r, c, sum)
	}

	r = row
	c = col - 1
	if hm.ValidPos(r, c) && heightmap[r][c] > old {
		heightmap, sum = Pool(heightmap, r, c, sum)
	}

	c = col + 1
	if hm.ValidPos(r, c) && heightmap[r][c] > old {
		heightmap, sum = Pool(heightmap, r, c, sum)
	}

	return heightmap, sum
}

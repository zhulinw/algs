package testdata

func JumpGameTestData() ([][]int, []bool) {
	inputTmp := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
		{1, 3, 2},
		{3},
		{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
		{2, 0, 6, 9, 8, 4, 5, 0, 8, 9, 1, 2, 9, 6, 8, 8, 0, 6, 3, 1, 2, 2, 1, 2, 6, 5, 3, 1, 2, 2, 6, 4, 2, 4, 3, 0, 0, 0, 3, 8, 2, 4, 0, 1, 2, 0, 1, 4, 6, 5, 8, 0, 7, 9, 3, 4, 6, 6, 5, 8, 9, 3, 4, 3, 7, 0, 4, 9, 0, 9, 8, 4, 3, 0, 7, 7, 1, 9, 1, 9, 4, 9, 0, 1, 9, 5, 7, 7, 1, 5, 8, 2, 8, 2, 6, 8, 2, 2, 7, 5, 1, 7, 9, 6},
	}
	except := []bool{
		true,
		false,
		true,
		true,
		true,
		false,
	}
	return inputTmp, except
}

func ClimbStairsTestData() ([]int, []int) {
	input := []int{2,3,5,23,44}
	except := []int{2,3,8,46368,1134903170}
	return input, except
}

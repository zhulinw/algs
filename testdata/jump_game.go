package testdata

func JumpGameTestData() ([][]int, []bool) {
	inputTmp := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
	}
	except := []bool{
		true,
		false,
	}
	return inputTmp, except
}

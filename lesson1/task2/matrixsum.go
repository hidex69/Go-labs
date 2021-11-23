package task2

func MatrixSum(matrix [][]int) int {
	sum := 0

	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j] == 0 {
				break
			} else {
				sum += matrix[i][j]
			}
		}
	}

	return sum
}

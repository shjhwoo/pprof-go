package sol

func Solution(arr1 [][]int, arr2 [][]int) [][]int {
	result := [][]int{}
	for i := range arr1 {
		resRow := []int{}
		for h := 0; h < len(arr2[0]); h++ {
			sum := 0
			for j := range arr1[i] {
				sum += arr1[i][j] * arr2[j][h]
			}
			resRow = append(resRow, sum)
		}
		result = append(result, resRow)
	}
	return result
}

package sol

import "testing"

// func BenchmarkSolution(b *testing.B) {
// 	arr1 := [][]int{{1, 4}, {3, 2}, {4, 1}}
// 	arr2 := [][]int{{3, 3}, {3, 3}}
// 	for i := 0; i < b.N; i++ {
// 		Solution(arr1, arr2)
// 	}
// }

func BenchmarkSolutionLargeInput(b *testing.B) {
	arr1 := [][]int{{2, 3, 2}, {4, 2, 4}, {3, 1, 4}}
	arr2 := [][]int{{5, 4, 3}, {2, 4, 1}, {3, 1, 1}}
	for i := 0; i < b.N; i++ {
		Solution(arr1, arr2)
	}
}

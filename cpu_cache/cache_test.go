package cpu_cache

import "testing"

func makeMatrix(n int64) [][]int64 {
	graph := make([][]int64, n)
	var i int64
	for i = 0; i < n; i++ {
		graph[i] = make([]int64, n)
	}
	return graph
}

var num int64 = 128

func BenchmarkMatrixCombination_Sequential_128(b *testing.B) {
	matrix1 := makeMatrix(num)
	matrix2 := makeMatrix(num)
	var i int64
	var j int64
	for n := 0; n < b.N; n++ {
		for i = 0; i < num; i++ {
			for j = 0; j < num; j++ {
				matrix1[i][j] = matrix1[i][j] + matrix2[i][j]
			}
		}
	}
}
func BenchmarkMatrixCombination_NotSequential_128(b *testing.B) {
	matrix1 := makeMatrix(num)
	matrix2 := makeMatrix(num)
	var i int64
	var j int64
	for n := 0; n < b.N; n++ {
		for i = 0; i < num; i++ {
			for j = 0; j < num; j++ {
				matrix1[i][j] = matrix1[i][j] + matrix2[j][i]
			}
		}
	}
}

var num2 int64 = 256

func BenchmarkMatrixCombination_Sequential_256(b *testing.B) {
	matrix1 := makeMatrix(num2)
	matrix2 := makeMatrix(num2)
	var i int64
	var j int64
	for n := 0; n < b.N; n++ {
		for i = 0; i < num2; i++ {
			for j = 0; j < num2; j++ {
				matrix1[i][j] = matrix1[i][j] + matrix2[i][j]
			}
		}
	}
}
func BenchmarkMatrixCombination_NotSequential_256(b *testing.B) {
	matrix1 := makeMatrix(num2)
	matrix2 := makeMatrix(num2)
	var i int64
	var j int64
	for n := 0; n < b.N; n++ {
		for i = 0; i < num2; i++ {
			for j = 0; j < num2; j++ {
				matrix1[i][j] = matrix1[i][j] + matrix2[j][i]
			}
		}
	}
}

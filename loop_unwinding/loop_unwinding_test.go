package loop_unwinding

import "testing"

type a struct {
	age int
}

func (aa a) incre(i int) {
	aa.age += i
}

func BenchmarkWinding(b *testing.B) {
	aa := a{age: 0}
	for i := 0; i < b.N; i++ {
		aa.incre(i)
	}
}
func BenchmarkUnWinding10(b *testing.B) {
	aa := a{age: 0}
	for i := 0; i < b.N; i += 10 {
		aa.incre(i)
		aa.incre(i + 1)
		aa.incre(i + 2)
		aa.incre(i + 3)
		aa.incre(i + 4)
		aa.incre(i + 5)
		aa.incre(i + 6)
		aa.incre(i + 7)
		aa.incre(i + 8)
		aa.incre(i + 9)
	}
}

func BenchmarkUnWinding30(b *testing.B) {
	aa := a{age: 0}
	for i := 0; i < b.N; i += 30 {
		aa.incre(i)
		aa.incre(i + 1)
		aa.incre(i + 2)
		aa.incre(i + 3)
		aa.incre(i + 4)
		aa.incre(i + 5)
		aa.incre(i + 6)
		aa.incre(i + 7)
		aa.incre(i + 8)
		aa.incre(i + 9)
		aa.incre(i + 10)
		aa.incre(i + 11)
		aa.incre(i + 12)
		aa.incre(i + 13)
		aa.incre(i + 14)
		aa.incre(i + 15)
		aa.incre(i + 16)
		aa.incre(i + 17)
		aa.incre(i + 18)
		aa.incre(i + 19)
		aa.incre(i + 20)
		aa.incre(i + 21)
		aa.incre(i + 22)
		aa.incre(i + 23)
		aa.incre(i + 24)
		aa.incre(i + 25)
		aa.incre(i + 26)
		aa.incre(i + 27)
		aa.incre(i + 28)
		aa.incre(i + 29)
	}
}

package shared_val

import (
	"fmt"
	"sync"
	"testing"
)

var gNum = 10

func BenchmarkSharedValRead(b *testing.B) {
	var val2 int
	for l := 0; l < b.N; l++ {
		val := 33
		wg := sync.WaitGroup{}
		b.StartTimer()
		for i := 0; i < gNum; i++ {
			wg.Add(1)
			go func() {
				// read
				val2 = val
				wg.Done()
			}()
		}
		wg.Wait()
		b.StopTimer()
	}
	fmt.Printf("val2 %d\n", val2)
}

func BenchmarkSharedValWrite(b *testing.B) {
	var val2 int = 398
	val := 0
	for l := 0; l < b.N; l++ {
		val = 33
		wg := sync.WaitGroup{}
		b.StartTimer()
		for i := 0; i < gNum; i++ {
			wg.Add(1)
			go func() {
				// write
				val = val2
				wg.Done()
			}()
		}
		wg.Wait()
		b.StopTimer()
	}
	fmt.Printf("val %d\n", val)
}

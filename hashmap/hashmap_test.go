package hashmap

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/mox692/tdata/array"
	"github.com/mox692/tdata/hashmap"
)

func BenchmarkArrayLinear(b *testing.B) {
	var len = 10000
	a := array.NewArrayInt(len)
	for i := 0; i < b.N; i++ {
		// rand.Seed(time.Hour.Nanoseconds())
		want := rand.Int() % len
		b.StartTimer()
		for i := 0; i < len; i++ {
			if a[i] == want {
				goto aaa
			}
		}
	aaa:
		b.StopTimer()
	}
}

func search(a []int, want int) (bool, int) {
	cur_len := len(a)
	cur_pos := 0
	for {
		if cur_len%2 == 0 {
			if cur_len == 0 {
				return false, 0
			}
			center := (cur_len >> 1) - 1
			if a[center] > want {
				a = a[0 : center-1]
				cur_len = (cur_len >> 1) - 1
				continue
			} else if a[center] < want {
				a = a[cur_len>>1:]
				cur_pos += cur_len >> 1
				cur_len = cur_len >> 1
				continue
			} else if a[center] == want {
				return true, cur_pos + center
			} else {
				return false, 0
			}
		} else {
			if cur_len == 0 {
				return false, 0
			}
			center := (cur_len - 1) >> 1
			if a[center] > want {
				a = a[0 : center-1]
				cur_len = (cur_len >> 1) - 1
				continue
			} else if a[center] < want {
				a = a[center+1:]
				cur_pos += center
				cur_len = (cur_len - 1) >> 1
				continue
			} else if a[center] == want {
				return true, cur_pos + center
			} else {
				return false, 0
			}
		}
	}

}

type SearchTestCase struct {
	input  []int
	want   int
	expect struct {
		have  bool
		index int
	}
}

func TestSearch(t *testing.T) {
	testCase := []SearchTestCase{
		{
			input: []int{22, 44, 55, 666},
			want:  55,
			expect: struct {
				have  bool
				index int
			}{
				have:  true,
				index: 2,
			},
		},
	}
	for _, v := range testCase {
		have, index := search2(v.input, v.want)
		if have != v.expect.have || index != v.expect.index {
			t.Errorf("Fail, expect %+v, got %+v, %+v", v.expect, have, index)
		}
	}
}

func BenchmarkArrayBinary(b *testing.B) {
	var len = 10
	a := array.NewArrayInt(len)
	for i := 0; i < b.N; i++ {
		rand.Seed(time.Hour.Nanoseconds())
		want := rand.Int() % len
		b.StartTimer()

		// sort
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

		for i := 0; i < len; i++ {
			if a[i] == want {
				goto aaa
			}
		}
	aaa:
		b.StopTimer()
	}
}

func BenchmarkHashMap(b *testing.B) {
	var len = 10000
	m := hashmap.NewHashMaps("key", len)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		_ = m["key33"]
		b.StopTimer()
	}
}

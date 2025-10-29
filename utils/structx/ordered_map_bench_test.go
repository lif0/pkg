package structx_test

import (
	"math/rand"
	"testing"

	"github.com/lif0/pkg/utils/structx"
)

func makeInts(n int) ([]int, []int) {
	keys := make([]int, n)
	vals := make([]int, n)
	r := rand.New(rand.NewSource(42))
	for i := 0; i < n; i++ {
		keys[i] = i
		vals[i] = r.Int()
	}
	return keys, vals
}

func makeStrs(n int) ([]string, [][]string) {
	keys := make([]string, n)
	vals := make([][]string, n)
	for i := 0; i < n; i++ {
		keys[i] = "k_" + string(rune('a'+(i%26))) + "_" + itoa(i)
		vals[i] = []string{"v", itoa(i)}
	}
	return keys, vals
}

func makeStrEmpties(n int) ([]string, []complexStruct) {
	keys := make([]string, n)
	vals := make([]complexStruct, n)
	for i := 0; i < n; i++ {
		keys[i] = "k_" + itoa(i)
		vals[i] = complexStruct{Val: i, Nums: []int{n}}
	}
	return keys, vals
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	buf := [20]byte{}
	pos := len(buf)
	neg := i < 0
	u := uint64(i)
	if neg {
		u = uint64(-i)
	}
	for u > 0 {
		pos--
		buf[pos] = byte('0' + u%10)
		u /= 10
	}
	if neg {
		pos--
		buf[pos] = '-'
	}
	return string(buf[pos:])
}

func Benchmark_OrderedMapIntInt(b *testing.B) {
	const N = 10_000
	keys, vals := makeInts(N)

	b.Run("put/orderedMap", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for n := 0; n < b.N; n++ {
			m := structx.NewOrderedMap[int, int](N)
			for i := 0; i < N; i++ {
				m.Put(keys[i], vals[i])
			}
		}
	})

	b.Run("put/builtin", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for n := 0; n < b.N; n++ {
			m := make(map[int]int, N)
			for i := 0; i < N; i++ {
				m[keys[i]] = vals[i]
			}
		}
	})

	b.Log("------------------------------")

	b.Run("get_hit/orderedMap", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		m := structx.NewOrderedMap[int, int](N)
		for i := 0; i < N; i++ {
			m.Put(keys[i], vals[i])
		}
		b.ResetTimer()
		var sink int

		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				v, _ := m.Get(keys[i])
				sink ^= v
			}
		}
		_ = sink
	})

	b.Run("get_hit/builtin", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		m := make(map[int]int, N)
		for i := 0; i < N; i++ {
			m[keys[i]] = vals[i]
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				v, _ := m[keys[i]]
				sink ^= v
			}
		}
		_ = sink
	})

	b.Log("------------------------------")

	b.Run("delete/orderedMap", func(b *testing.B) {
		b.ReportAllocs()
		m := structx.NewOrderedMap[int, int](N)
		for i := 0; i < N; i++ {
			m.Put(keys[i], vals[i])
		}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				structx.Delete(m, keys[i])
			}
		}
	})

	b.Run("delete/builtin", func(b *testing.B) {
		b.ReportAllocs()
		m := make(map[int]int, N)
		for i := 0; i < N; i++ {
			m[keys[i]] = vals[i]
		}

		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				delete(m, keys[i])
			}
		}
	})

	b.Log("------------------------------")

	b.Run("iterate_values/orderedMap", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		m := structx.NewOrderedMap[int, int](N)
		for i := 0; i < N; i++ {
			m.Put(keys[i], vals[i])
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for _, v := range m.Iter() {
				sink ^= v
			}
		}
		_ = sink
	})

	b.Run("iterate_values/builtin_range", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		m := make(map[int]int, N)
		for i := 0; i < N; i++ {
			m[keys[i]] = vals[i]
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for _, v := range m {
				sink ^= v
			}
		}
		_ = sink
	})
}

func Benchmark_OrderedMap_vs_Builtin_StringSlice(b *testing.B) {
	const N = 1_0000
	keys, vals := makeStrs(N)

	b.Run("put/ordered", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			m := structx.NewOrderedMap[string, []string](N)
			for i := 0; i < N; i++ {
				m.Put(keys[i], vals[i])
			}
		}
	})

	b.Run("put/builtin", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			m := make(map[string][]string, N)
			for i := 0; i < N; i++ {
				m[keys[i]] = vals[i]
			}
		}
	})

	b.Log("------------------------------")

	b.Run("get_hit/ordered", func(b *testing.B) {
		b.ReportAllocs()
		m := structx.NewOrderedMap[string, []string](N)
		for i := 0; i < N; i++ {
			m.Put(keys[i], vals[i])
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				v, _ := m.Get(keys[i])
				if len(v) > 0 {
					sink ^= len(v[0])
				}
			}
		}
		_ = sink
	})

	b.Run("get_hit/builtin", func(b *testing.B) {
		b.ReportAllocs()
		m := make(map[string][]string, N)
		for i := 0; i < N; i++ {
			m[keys[i]] = vals[i]
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				v, _ := m[keys[i]]
				if len(v) > 0 {
					sink ^= len(v[0])
				}
			}
		}
		_ = sink
	})

	b.Log("------------------------------")

	b.Run("delete/ordered", func(b *testing.B) {
		b.ReportAllocs()
		m := structx.NewOrderedMap[string, []string](N)
		for i := 0; i < N; i++ {
			m.Put(keys[i], vals[i])
		}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				structx.Delete(m, keys[i])
			}
		}
	})

	b.Run("delete/builtin", func(b *testing.B) {
		b.ReportAllocs()
		m := make(map[string][]string, N)
		for i := 0; i < N; i++ {
			m[keys[i]] = vals[i]
		}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				delete(m, keys[i])
			}
		}
	})

	b.Log("------------------------------")

	b.Run("iterate_values/ordered", func(b *testing.B) {
		b.ReportAllocs()
		m := structx.NewOrderedMap[string, []string](N)
		for i := 0; i < N; i++ {
			m.Put(keys[i], vals[i])
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for range m.Iter() {
				sink ^= 1
			}
		}
		_ = sink
	})

	b.Run("iterate_values/builtin_range", func(b *testing.B) {
		b.ReportAllocs()
		m := make(map[string][]string, N)
		for i := 0; i < N; i++ {
			m[keys[i]] = vals[i]
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for _, v := range m {
				if len(v) > 1 {
					sink ^= len(v[1])
				}
			}
		}
		_ = sink
	})
}

func Benchmark_OrderedMap_vs_Builtin_StringComplexStruct(b *testing.B) {
	const N = 1_0000
	keys, vals := makeStrEmpties(N)

	b.Run("put/ordered", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			m := structx.NewOrderedMap[string, complexStruct](N)
			for i := 0; i < N; i++ {
				m.Put(keys[i], vals[i])
			}
		}
	})

	b.Run("put/builtin", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			m := make(map[string]complexStruct, N)
			for i := 0; i < N; i++ {
				m[keys[i]] = vals[i]
			}
		}
	})

	b.Log("------------------------------")

	b.Run("get_hit/ordered", func(b *testing.B) {
		b.ReportAllocs()
		m := structx.NewOrderedMap[string, complexStruct](N)
		for i := 0; i < N; i++ {
			m.Put(keys[i], vals[i])
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				_, ok := m.Get(keys[i])
				if ok {
					sink ^= 1
				}
			}
		}
		_ = sink
	})

	b.Run("get_hit/builtin", func(b *testing.B) {
		b.ReportAllocs()
		m := make(map[string]complexStruct, N)
		for i := 0; i < N; i++ {
			m[keys[i]] = vals[i]
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				_, ok := m[keys[i]]
				if ok {
					sink ^= 1
				}
			}
		}
		_ = sink
	})

	b.Log("------------------------------")

	b.Run("delete/ordered", func(b *testing.B) {
		b.ReportAllocs()
		m := structx.NewOrderedMap[string, complexStruct](N)
		for i := 0; i < N; i++ {
			m.Put(keys[i], vals[i])
		}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				structx.Delete(m, keys[i])
			}
		}
	})

	b.Run("delete/builtin", func(b *testing.B) {
		b.ReportAllocs()
		m := make(map[string]complexStruct, N)
		for i := 0; i < N; i++ {
			m[keys[i]] = vals[i]
		}
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			for i := 0; i < N; i++ {
				delete(m, keys[i])
			}
		}
	})

	b.Log("------------------------------")

	b.Run("iterate_values/ordered", func(b *testing.B) {
		b.ReportAllocs()
		m := structx.NewOrderedMap[string, complexStruct](N)
		for i := 0; i < N; i++ {
			m.Put(keys[i], vals[i])
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for range m.Iter() {
				sink ^= 1
			}
		}
		_ = sink
	})

	b.Run("iterate_values/builtin_range", func(b *testing.B) {
		b.ReportAllocs()
		m := make(map[string]complexStruct, N)
		for i := 0; i < N; i++ {
			m[keys[i]] = vals[i]
		}
		b.ResetTimer()
		var sink int
		for n := 0; n < b.N; n++ {
			for range m {
				sink ^= 1
			}
		}
		_ = sink
	})
}

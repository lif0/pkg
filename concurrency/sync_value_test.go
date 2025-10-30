package concurrency_test

import (
	"sync"
	"testing"
	"time"

	"github.com/lif0/pkg/concurrency"
	"github.com/stretchr/testify/assert"
)

type complexStruct struct {
	Flag  bool
	Nums  []int
	Index map[string]int
}

func Test_Snapshot(t *testing.T) {
	t.Parallel()

	t.Run("ok/int_snapshot_not_shared", func(t *testing.T) {
		t.Parallel()
		// arrange
		sv := concurrency.NewSyncValue[int]()
		// act
		sv.MutateValue(func(v *int) {
			*v = 12
		})
		// assert
		var got int
		sv.ReadValue(func(v *int) {
			got = *v
		})
		assert.Equal(t, 12, got)

		sv.MutateValue(func(v *int) {
			*v = 20
		})
		assert.Equal(t, 12, got) // check that underlining value is not shared
	})

	t.Run("ok/slice_snapshot_not_shared_deep_copy", func(t *testing.T) {
		t.Parallel()
		// arrange
		sv := concurrency.NewSyncValue[[]int]([]int{1, 2, 3})
		var snap []int
		// act
		sv.ReadValue(func(v *[]int) {
			snap = make([]int, len(*v))
			copy(snap, *v)
		})
		sv.MutateValue(func(v *[]int) {
			*v = append(*v, 99)
			(*v)[0] = 777
		})
		// assert
		assert.Equal(t, []int{1, 2, 3}, snap)
	})

	t.Run("ok/map_snapshot_not_shared_deep_copy", func(t *testing.T) {
		t.Parallel()
		// arrange
		sv := concurrency.NewSyncValue[map[string]int](map[string]int{"a": 1, "b": 2})
		snap := make(map[string]int)
		// act
		sv.ReadValue(func(v *map[string]int) {
			for k, val := range *v {
				snap[k] = val
			}
		})
		sv.MutateValue(func(v *map[string]int) {
			(*v)["a"] = 10
			(*v)["c"] = 3
		})
		// assert
		assert.Equal(t, map[string]int{"a": 1, "b": 2}, snap)
	})

	t.Run("ok/complexStruct_snapshot_not_shared_deep_copy", func(t *testing.T) {
		t.Parallel()
		// arrange
		init := complexStruct{
			Flag:  true,
			Nums:  []int{1, 2, 3},
			Index: map[string]int{"x": 1},
		}
		sv := concurrency.NewSyncValue[complexStruct](init)

		var snap complexStruct
		// act
		sv.ReadValue(func(v *complexStruct) {
			snap.Flag = v.Flag
			snap.Nums = append([]int(nil), v.Nums...)
			snap.Index = make(map[string]int, len(v.Index))
			for k, val := range v.Index {
				snap.Index[k] = val
			}
		})
		sv.MutateValue(func(v *complexStruct) {
			v.Flag = false
			v.Nums[0] = 999
			v.Nums = append(v.Nums, 42)
			v.Index["x"] = 2
			v.Index["y"] = 3
		})
		// assert
		assert.True(t, snap.Flag)
		assert.Equal(t, []int{1, 2, 3}, snap.Nums)
		assert.Equal(t, map[string]int{"x": 1}, snap.Index)
	})
}

func Test_IntValue(t *testing.T) {
	t.Parallel()

	t.Run("ok/basic_set_get", func(t *testing.T) {
		t.Parallel()
		// arrange
		sv := concurrency.NewSyncValue[int]()
		// act
		sv.MutateValue(func(v *int) {
			*v = 12
		})
		// assert
		var got int
		sv.ReadValue(func(v *int) {
			got = *v
		})
		assert.Equal(t, 12, got)
	})

	t.Run("race/concurrent_increments", func(t *testing.T) {
		t.Parallel()
		// arrange
		const writers = 4
		const readers = writers
		const perWorker = 1000
		sv := concurrency.NewSyncValue[int](0)

		var wg sync.WaitGroup
		wg.Add(writers + readers)

		// act
		go func() {
			for i := 0; i < writers; i++ {
				go func() {
					defer wg.Done()
					for j := 0; j < perWorker; j++ {
						sv.MutateValue(func(v *int) {
							*v++
						})
					}
				}()
			}
		}()

		go func() {
			time.Sleep(time.Millisecond * 5)
			for i := 0; i < readers; i++ {
				go func() {
					defer wg.Done()
					var out int
					for j := 0; j < perWorker; j++ {
						sv.ReadValue(func(v *int) {
							out = *v
						})
					}

					assert.True(t, out > 0)
				}()
			}
		}()

		wg.Wait()
		// assert
		var got int
		sv.ReadValue(func(v *int) {
			got = *v
		})
		assert.Equal(t, writers*perWorker, got)
	})
}

func Test_SliceValue(t *testing.T) {
	t.Parallel()

	t.Run("ok/basic_append_and_read", func(t *testing.T) {
		t.Parallel()
		// arrange
		sv := concurrency.NewSyncValue[[]int]([]int{1, 2})
		// act
		sv.MutateValue(func(v *[]int) {
			*v = append(*v, 3)
		})
		// assert
		var got []int
		sv.ReadValue(func(v *[]int) {
			got = append([]int(nil), (*v)...)
		})
		assert.Equal(t, []int{1, 2, 3}, got)
	})

	t.Run("race/concurrent_append_and_read_len", func(t *testing.T) {
		t.Parallel()
		// arrange
		sv := concurrency.NewSyncValue[[]int]()
		const writers = 4
		const readers = writers
		const perWriter = 300

		var wg sync.WaitGroup
		wg.Add(writers + readers)
		// act
		go func() {
			for i := 0; i < writers; i++ {
				go func(base int) {
					defer wg.Done()

					for j := 0; j < perWriter; j++ {
						val := base*perWriter + j
						sv.MutateValue(func(v *[]int) {
							*v = append(*v, val)
						})
					}
				}(i)
			}
		}()

		go func() {
			time.Sleep(time.Millisecond * 5)

			for i := 0; i < readers; i++ {
				go func() {
					defer wg.Done()
					var len_ int
					for j := 0; j < perWriter; j++ {
						sv.ReadValue(func(v *[]int) {
							len_ = len(*v)
						})
					}

					assert.True(t, len_ > 1)
				}()
			}
		}()
		wg.Wait()
		// assert
		var gotLen int
		sv.ReadValue(func(v *[]int) {
			gotLen = len(*v)
		})
		assert.Equal(t, writers*perWriter, gotLen)
	})
}

func Test_MapValue(t *testing.T) {
	t.Parallel()

	t.Run("ok/basic_put_get", func(t *testing.T) {
		t.Parallel()
		// arrange
		sv := concurrency.NewSyncValue[map[string]int](map[string]int{})
		// act
		sv.MutateValue(func(v *map[string]int) {
			(*v)["a"] = 1
		})
		// assert
		var got int
		sv.ReadValue(func(v *map[string]int) {
			got = (*v)["a"]
		})
		assert.Equal(t, 1, got)
	})

	t.Run("race/concurrent_increments", func(t *testing.T) {
		t.Parallel()
		// arrange
		sv := concurrency.NewSyncValue[map[string]int](map[string]int{"x": 0})
		const writers = 4
		const readers = writers
		const perWorker = 400

		var wg sync.WaitGroup
		wg.Add(writers + readers)
		// act
		go func() {
			for i := 0; i < writers; i++ {
				go func() {
					defer wg.Done()
					for j := 0; j < perWorker; j++ {
						sv.MutateValue(func(v *map[string]int) {
							(*v)["x"] = (*v)["x"] + 1
						})
					}
				}()
			}
		}()

		go func() {
			time.Sleep(time.Millisecond * 5)
			for i := 0; i < readers; i++ {
				go func() {
					defer wg.Done()

					var x int
					for j := 0; j < perWorker; j++ {
						sv.ReadValue(func(v *map[string]int) {
							x = (*v)["x"]
						})
					}

					assert.True(t, x > 0)
				}()
			}
		}()
		wg.Wait()
		// assert
		var got int
		sv.ReadValue(func(v *map[string]int) {
			got = (*v)["x"]
		})
		assert.Equal(t, writers*perWorker, got)
	})
}

func Test_StructValue(t *testing.T) {
	t.Parallel()

	t.Run("ok/mutate_multiple_fields", func(t *testing.T) {
		t.Parallel()
		// arrange
		init := complexStruct{Flag: false, Nums: []int{1, 2}, Index: map[string]int{"a": 1}}
		sv := concurrency.NewSyncValue[complexStruct](init)
		// act
		sv.MutateValue(func(v *complexStruct) {
			v.Flag = true
			v.Nums = append(v.Nums, 3)
			v.Index["b"] = 2
		})
		// assert
		sv.ReadValue(func(v *complexStruct) {
			assert.True(t, v.Flag)
			assert.Equal(t, []int{1, 2, 3}, v.Nums)
			assert.Equal(t, 2, v.Index["b"])
		})
	})

	t.Run("race/concurrent_read_write", func(t *testing.T) {
		t.Parallel()
		// arrange
		sv := concurrency.NewSyncValue[complexStruct](complexStruct{
			Flag:  false,
			Nums:  []int{},
			Index: map[string]int{},
		})

		const writers = 4
		const readers = writers
		const perWriter = 200
		var wg sync.WaitGroup
		wg.Add(writers + readers)

		// act
		go func() {
			for i := 0; i < writers; i++ {
				go func(id int) {
					defer wg.Done()
					for j := 0; j < perWriter; j++ {
						sv.MutateValue(func(v *complexStruct) {
							v.Flag = true
							v.Nums = append(v.Nums, id*perWriter+j)
							v.Index["cnt"] = v.Index["cnt"] + 1
						})
					}
				}(i)
			}
		}()

		go func() {
			time.Sleep(time.Millisecond * 5)
			for i := 0; i < readers; i++ {
				go func(id int) {
					defer wg.Done()
					for j := 0; j < perWriter; j++ {
						sv.ReadValue(func(v *complexStruct) {
							_ = v.Flag
							_ = len(v.Nums)
							_ = v.Nums[0]
							_ = v.Index["cnt"]
						})
					}
				}(i)
			}
		}()
		wg.Wait()
		// assert
		sv.ReadValue(func(v *complexStruct) {
			assert.True(t, v.Flag)
			assert.Equal(t, writers*perWriter, len(v.Nums))
			assert.Equal(t, writers*perWriter, v.Index["cnt"])
		})
	})
}

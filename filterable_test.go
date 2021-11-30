package filterable

import (
	"testing"
)

func TestBuildObject(t *testing.T) {

	origin_slice := []int{7, 2, 3, 4}
	my_slice := New(origin_slice)

	if my_slice == nil {
		t.Fatalf("new slice is nil")
	}

	inner_slice := my_slice.Get().([]int)

	for index, int := range inner_slice {
		if int != origin_slice[index] {
			t.Fatalf("expected to find %d but found %d", origin_slice[index], int)
		}
	}
}

func xTestBuildObjectVaraidic(t *testing.T) {

	// origin_slice := []int{1, 2, 3}
	// my_slice := New(1, 2, 3)

	// if my_slice == nil {
	// 	t.Fatalf("new slice is nil")
	// }

	// inner_slice := my_slice.Get().([]int)

	// for index, int := range inner_slice {
	// 	if int != origin_slice[index] {
	// 		t.Fatalf("expected to find %d but found %d", origin_slice[index], int)
	// 	}
	// }
}

func TestCanFilter(t *testing.T) {

	origin_slice := []int{100, 2, 3, 4}
	my_slice := New(origin_slice)

	only_small_numbers := my_slice.Filter(func(i interface{}) bool {
		return i.(int) < 100
	}).Get()

	expected := []int{2, 3, 4}
	for index, int := range only_small_numbers.([]interface{}) {
		if int != expected[index] {
			t.Fatalf("expected to find %d but found %d instead", expected[index], int)
		}
	}

}

func TestCanFilterMultipleTimes(t *testing.T) {

	origin_slice := []int{100, 2, 3, 4}
	my_slice := New(origin_slice)

	only_small_and_even_numbers := my_slice.
		Filter(
			func(i interface{}) bool {
				return i.(int) < 100
			}).
		Filter(
			func(i interface{}) bool {
				return i.(int)%2 == 0
			}).Get()

	expected := []int{2, 4}
	for index, int := range only_small_and_even_numbers.([]interface{}) {
		if int != expected[index] {
			t.Fatalf("expected to find %d but found %d instead", expected[index], int)
		}
	}

}

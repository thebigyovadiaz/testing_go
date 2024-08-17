package array_slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum numbers positive", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertExpectedMessage(t, AssertAttr{Got: got, Want: want, Arr: numbers})
	})

	t.Run("sum numbers positive and negative", func(t *testing.T) {
		numbers := []int{1, -2, 3, -4, 5}
		got := Sum(numbers)
		want := 3
		assertExpectedMessage(t, AssertAttr{Got: got, Want: want, Arr: numbers})
	})

	t.Run("sum numbers - empty slices", func(t *testing.T) {
		var numbers []int
		got := Sum(numbers)
		want := 0
		assertExpectedMessage(t, AssertAttr{Got: got, Want: want, Arr: numbers})
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{3, 4}, []int{8})
		want := []int{6, 7, 8}
		assertSliceExpectedMessage(t, AssertSliceAttr{Got: got, Want: want})
	})

	t.Run("sum all - empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{}, []int{})
		var want []int
		assertSliceExpectedMessage(t, AssertSliceAttr{Got: got, Want: want})
	})
}

func TestSumaAllTail(t *testing.T) {
	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3}, []int{3, 4}, []int{8})
		want := []int{5, 4, 0}
		assertSliceExpectedMessage(t, AssertSliceAttr{Got: got, Want: want})
	})

	t.Run("sum all tails - negatives and positives", func(t *testing.T) {
		got := SumAllTail([]int{1, 2, 3, -6}, []int{3, 4, -3}, []int{8, -2})
		want := []int{-1, 1, -2}
		assertSliceExpectedMessage(t, AssertSliceAttr{Got: got, Want: want})
	})

	t.Run("sum all tails - negatives and positives", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{})
		want := []int{0, 0}
		assertSliceExpectedMessage(t, AssertSliceAttr{Got: got, Want: want})
	})
}

type AssertAttr struct {
	Got  int
	Want int
	Arr  []int
}

func assertExpectedMessage(t testing.TB, asserAttr AssertAttr) {
	t.Helper()
	if asserAttr.Got != asserAttr.Want {
		t.Errorf("got %q want %q given, %v", asserAttr.Got, asserAttr.Want, asserAttr.Arr)
	}
}

type AssertSliceAttr struct {
	Got  []int
	Want []int
	Arr  []int
}

func assertSliceExpectedMessage(t testing.TB, asserSliceAttr AssertSliceAttr) {
	t.Helper()
	if !reflect.DeepEqual(asserSliceAttr.Got, asserSliceAttr.Want) {
		t.Errorf("got %v want %v", asserSliceAttr.Got, asserSliceAttr.Want)
	}
}

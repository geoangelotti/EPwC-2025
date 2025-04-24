package epwc_test

import (
	"epwc"
	"testing"
)

func TestEvaluatePolynomial(t *testing.T) {
	input := []int{1, 3, 1, 4, 1, 5}
	unique1 := epwc.UniqueCountSortSlice(input)
	unique2 := epwc.UniqueCountSlicesSort(input)
	unique3 := epwc.UniqueCountMap(input)
	res := 4
	if res != unique1 {
		t.Errorf("Expected %v but got %v", res, unique1)
	}
	if res != unique2 {
		t.Errorf("Expected %v but got %v", res, unique2)
	}
	if res != unique3 {
		t.Errorf("Expected %v but got %v", res, unique3)
	}
}

func BenchmarkUniqueSortSlice(b *testing.B) {
	input := []int{1, 3, 1, 4, 1, 5}
	for b.Loop() {
		_ = epwc.UniqueCountSortSlice(input)
	}
}

func BenchmarkUniqueSlicesSort(b *testing.B) {
	input := []int{1, 3, 1, 4, 1, 5}
	for b.Loop() {
		_ = epwc.UniqueCountSlicesSort(input)
	}
}

func BenchmarkUniqueMap(b *testing.B) {
	input := []int{1, 3, 1, 4, 1, 5}
	for b.Loop() {
		_ = epwc.UniqueCountMap(input)
	}
}

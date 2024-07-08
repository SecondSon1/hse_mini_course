package hw1

import (
	"slices"
	"testing"
)

// Util
func SlicesEqual[T comparable](a []T, b []T) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRemoveDuplicates(t *testing.T) {
	empty := RemoveDuplicates([]int{})
	if len(empty) != 0 {
		t.Fatalf("Expected empty slice on empty input, got %v", empty)
	}
	unique_input := []int{100000, -200000, 30, 69}
	unique := RemoveDuplicates(unique_input)
	if !SlicesEqual(unique_input, unique) {
		t.Fatalf("Fail:\n  Expected: %v\n  Got: %v\n",
			unique_input, unique)
	}
	random_input := []int{1, 1, -3, 10000000, 1, 3, 10000001}
	random := RemoveDuplicates(random_input)
	if !SlicesEqual(unique_input, unique) {
		t.Fatalf("Fail:\n  Expected: %v\n  Got: %v\n",
			random_input, random)
	}
}

func TestBubbleSort(t *testing.T) {
	test1 := []int{-14, 75, 12, -8080, 134}
	test1_sorted := BubbleSort(test1)
	slices.Sort(test1)
	if !SlicesEqual(test1, test1_sorted) {
		t.Fatalf("Fail:\n  Expected: %v\n  Got: %v\n",
			test1, test1_sorted)
	}
	test2 := []int{0, 0, -1441425, 741, 1042, 1314, -4193, 987, 987, 741}
	test2_sorted := BubbleSort(test2)
	slices.Sort(test2)
	if !SlicesEqual(test2, test2_sorted) {
		t.Fatalf("Fail:\n  Expected: %v\n  Got: %v\n",
			test2, test2_sorted)
	}
}

func TestFibonacci(t *testing.T) {
	fib := Fibonacci(10)
	true_fib := []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	if !SlicesEqual(fib, true_fib) {
		t.Fatalf("Fail:\n  Expected: %v\n  Got: %v\n",
			true_fib, fib)
	}
}

func TestCountOccurences(t *testing.T) {
	arr := []int{14, 15, 73, -21, -1550, 13, 73, 73}
	res := CountOccurences(arr, 73)
	if res != 3 {
		t.Fatalf("Fail:\n  Expected: %d\n  Got: %d\n",
			3, res)
	}
}

func TestArrayIntersection(t *testing.T) {
	arr1 := []int{10, 414, 1452, 313, 998, -1449}
	arr2 := []int{10, -10000, 414, -414, -10, 999, 1452, 10}
	result := ArrayIntersection(arr1, arr2)
	slices.Sort(result)
	expected := []int{10, 414, 1452}
	if !SlicesEqual(result, expected) {
		t.Fatalf("Fail:\n  Expected: %v\n  Got: %v\n",
			expected, result)
	}
}

func TestAreAnagrams(t *testing.T) {
	test1_a := "aBoBa"
	test1_b := "AAbbo"
	if !AreAnagrams(test1_a, test1_b) {
		t.Fatalf("Fail:\n  Input \"%s\" and \"%s\" are expected to be anagrams (case insensitive)\n",
			test1_a, test1_b)
	}
	test2_a := ""
	test2_b := "bebra"
	if AreAnagrams(test2_a, test2_b) {
		t.Fatalf("Fail:\n  Input \"%s\" and \"%s\" are not expected to be anagrams\n",
			test2_a, test2_b)
	}

	test3_a := "ooo"
	test3_b := "oAo"
	if AreAnagrams(test3_a, test3_b) {
		t.Fatalf("Fail:\n  Input \"%s\" and \"%s\" are not expected to be anagrams\n",
			test3_a, test3_b)
	}
}

func TestMerge(t *testing.T) {
	arr1 := []int{1, 2, 5, 9, 11}
	arr2 := []int{2, 4, 5, 8, 12, 15}
	merged := Merge(arr1, arr2)
	arr1 = append(arr1, arr2...)
	slices.Sort(arr1)
	if !SlicesEqual(merged, arr1) {
		t.Fatalf("Fail:\n  Expected: %v\n  Got: %v\n",
			arr1, merged)
	}
}

func TestBinarySearch(t *testing.T) {
	test1 := []int{1, 5, 100, 9990, 234567}
	res := BinarySearch(test1, 105)
	expected := 3
	if res != expected {
		t.Fatalf("Fail:\n  Expected: %d\n  Got: %d\n",
			expected, res)
	}

	test2 := []int{-414, -45, -1, 0, 13, 155, 987, 9147}
	res = BinarySearch(test2, 11111111)
	expected = len(test2)
	if res != expected {
		t.Fatalf("Fail:\n  Expected: %d\n  Got: %d\n",
			expected, res)
	}
}

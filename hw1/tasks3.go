package hw1

import (
	"unicode"
)

func cloneSlice[T any](arr []T) []T {
	new_arr := make([]T, len(arr))
	copy(new_arr, arr)
	return new_arr
}

func RemoveDuplicates(array []int) []int {
	set := make(map[int]bool)
	var result []int
	for _, element := range array {
		if set[element] {
			continue
		}
		set[element] = true
		result = append(result, element)
	}
	return result
}

func BubbleSort(array []int) []int {
	n := int(len(array))
	result := cloneSlice(array)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

func Fibonacci(n uint) []uint64 {
	result := make([]uint64, n)
	if n >= 1 {
		result[0] = 1
	}
	if n >= 2 {
		result[1] = 1
	}
	var i uint
	for i = 2; i < n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result
}

func CountOccurences[T comparable](arr []T, element T) uint {
	var count uint = 0
	for _, el := range arr {
		if el == element {
			count++
		}
	}
	return count
}

func ArrayIntersection(a, b []int) []int {
	elems := make(map[int]bool)
	for _, el := range a {
		elems[el] = true
	}
	var result []int
	for _, el := range b {
		if elems[el] {
			result = append(result, el)
			elems[el] = false // so no duplicates
		}
	}
	return result
}

func AreAnagrams(a, b string) bool {
	cnt := make(map[rune]int)
	for _, ch := range a {
		cnt[unicode.ToLower(ch)]++
	}
	for _, ch := range b {
		cnt[unicode.ToLower(ch)]--
	}
	ok := true
	for _, el_cnt := range cnt {
		if el_cnt != 0 {
			ok = false
			break
		}
	}
	return ok
}

func Merge(a, b []int) []int {
	result := make([]int, len(a)+len(b))
	l := 0
	r := 0
	for l < len(a) && r < len(b) {
		if a[l] < b[r] {
			result[l+r] = a[l]
			l++
		} else {
			result[l+r] = b[r]
			r++
		}
	}
	for ; l < len(a); l++ {
		result[l+r] = a[l]
	}
	for ; r < len(b); r++ {
		result[l+r] = b[r]
	}
	return result
}

// Will find first element that is >= element, equivalent to std::lower_bound in c++
// Returns its index and len(arr) if all elements are less
func BinarySearch(arr []int, element int) int {
	l := 0
	r := len(arr)
	for l+1 < r {
		m := (l + r) / 2
		if arr[m] < element {
			l = m
		} else {
			r = m
		}
	}
	return r
}

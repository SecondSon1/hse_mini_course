package hw1

import (
	"strconv"
	"testing"
)

const (
	UINT_MASK    = ^uint(0)
	RANDOM_PRIME = uint(991)
	LARGE_CAP    = 1_000_000
)

func hashFuncOne(s string) uint {
	var result uint
	for _, ch := range s {
		result = (result*RANDOM_PRIME + uint(ch)) & UINT_MASK
	}
	return result
}

func TestBasic(t *testing.T) {
	hash_map := NewHashMap[string, int](hashFuncOne)
	_, found := hash_map.Get("hey")
	if found {
		t.Fatalf("Found something in an empty hashmap, impressive!\n")
	}
	ok := hash_map.Delete("hey")
	if ok {
		t.Fatalf("Deleted something from an empty hashmap, impressive!\n")
	}
	hash_map.Set("hey", 420)
	var val int
	val, found = hash_map.Get("hey")
	if !found {
		t.Fatalf("Not found element\n")
	}
	if val != 420 {
		t.Fatalf("Element found, but value is %d instead of expected %d\n",
			val, 420)
	}
	ok = hash_map.Delete("hey")
	if !ok {
		t.Fatalf("Could not delete element\n")
	}
	_, found = hash_map.Get("hey")
	if found {
		t.Fatalf("Found deleted element\n")
	}
}

func TestLarge(t *testing.T) {
	hash_map := NewHashMap[string, int](hashFuncOne)

	for i := 1; i <= LARGE_CAP; i++ {
		hash_map.Set(strconv.Itoa(i), i)
	}

	if hash_map.size != LARGE_CAP {
		t.Fatalf("Size incorrect, expected %d, got %d\n",
			LARGE_CAP, hash_map.size)
	}

	hash_map.Set("1111", 2222)
	value, found := hash_map.Get("7373")
	if !found {
		t.Fatalf("Could not find element\n")
	}
	if value != 7373 {
		t.Fatalf("Incorrect element: expected 7373, got %d\n",
			hash_map.size)
	}

	hash_map.Delete("6969")
	if _, ok := hash_map.Get("6969"); ok {
		t.Fatalf("Found deleted element\n")
	}

	value, found = hash_map.Get("1111")
	if !found {
		t.Fatalf("Could not element that was re-set\n")
	}
	if value != 2222 {
		t.Fatalf("Incorrect element: expected 2222, got %d\n",
			value)
	}

	if hash_map.size != LARGE_CAP-1 {
		t.Fatalf("Incorrect size: expected %d, got %d\n",
			LARGE_CAP-1, hash_map.size)
	}
}

func TestIteration(t *testing.T) {
	hash_map := NewHashMap[string, int](hashFuncOne)
	for i := 1; i <= 20; i++ {
		hash_map.Set(strconv.Itoa(i), i)
	}

	used := make([]bool, 21)
	elements := hash_map.GetAll()
	if len(elements) != 20 {
		t.Fatalf("Incorrect elements length: expected 20, got %d\n",
			len(elements))
	}
	for _, kv := range elements {
		key := kv.Key
		value := kv.Value
		if !(1 <= value && value <= 20) {
			t.Fatalf("Incorrect key-value pair: (%s, %d)\n", key, value)
		}
		if used[value] {
			t.Fatalf("Value encountered twice: (%s, %d)\n", key, value)
		}
		used[value] = true
		if key != strconv.Itoa(value) {
			t.Fatalf("Value does not match key: (%s, %d)\n", key, value)
		}
	}
}

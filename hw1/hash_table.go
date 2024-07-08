package hw1

const (
	HASH_MAP_LOAD_FACTOR = 2
)

type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

type HashMap[K comparable, V any] struct {
	size      uint
	hash_func func(K) uint
	table     [][]KeyValue[K, V]
}

func NewHashMap[K comparable, V any](hash_func func(K) uint) HashMap[K, V] {
	return HashMap[K, V]{
		size:      0,
		hash_func: hash_func,
		table:     nil,
	}
}

func (h *HashMap[K, V]) GrowIfNeeded() {
	buck := uint(len(h.table))
	if buck == 0 {
		h.table = make([][]KeyValue[K, V], 1)
		return
	}
	if buck*HASH_MAP_LOAD_FACTOR > h.size {
		return
	}
	new_buck := buck * 2
	new_table := make([][]KeyValue[K, V], new_buck)

	var i uint
	for i = 0; i < buck; i++ {
		for _, key_val := range h.table[i] {
			hash := h.hash_func(key_val.Key) % new_buck
			new_table[hash] = append(new_table[hash], key_val)
		}
	}
	h.table = new_table
}

// returns (object, found), object is zero if found is false
func (h *HashMap[K, V]) Get(key K) (V, bool) {
	var zero V
	if h.size == 0 {
		return zero, false
	}
	hash := h.hash_func(key) % uint(len(h.table))
	for _, el := range h.table[hash] {
		if el.Key == key {
			return el.Value, true
		}
	}
	return zero, false
}

func (h *HashMap[K, V]) Set(key K, value V) {
	var hash uint
	if h.size > 0 {
		hash = h.hash_func(key) % uint(len(h.table))
		for ind, el := range h.table[hash] {
			if el.Key == key {
				h.table[hash][ind] = KeyValue[K, V]{
					Key:   key,
					Value: value,
				}
				return
			}
		}
	}
	h.GrowIfNeeded()
	hash = h.hash_func(key) % uint(len(h.table))
	h.table[hash] = append(h.table[hash], KeyValue[K, V]{
		Key:   key,
		Value: value,
	})
	h.size++
}

// true if deleted, false if not found
func (h *HashMap[K, V]) Delete(key K) bool {
	if h.size == 0 {
		return false
	}
	hash := h.hash_func(key) % uint(len(h.table))
	for ind, el := range h.table[hash] {
		if el.Key == key {
			last := len(h.table[hash]) - 1
			h.table[hash][ind], h.table[hash][last] = h.table[hash][last], h.table[hash][ind]
			h.table[hash] = h.table[hash][:last]
			h.size--
			return true
		}
	}
	return false
}

func (h *HashMap[K, V]) GetAll() []KeyValue[K, V] {
	result := make([]KeyValue[K, V], 0, h.size)
	buck := len(h.table)
	for i := 0; i < buck; i++ {
		for _, kv := range h.table[i] {
			result = append(result, kv)
		}
	}
	return result
}

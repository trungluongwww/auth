package slicer

func ChunkBy[T any](items []T, size int) (chunks [][]T) {
	for size < len(items) {
		items, chunks = items[size:], append(chunks, items[0:size:size])
	}
	return append(chunks, items)
}

func Map[T, U any](src []T, fn func(T) U) []U {
	dest := make([]U, len(src))
	for i := range src {
		dest[i] = fn(src[i])
	}
	return dest
}

func Contains[T comparable](items []T, subject T) bool {
	return ContainsAny(items, subject, func(a, b T) bool { return a == b })
}

func ContainsAny[T any](items []T, subject T, eq func(a, b T) bool) bool {
	for _, v := range items {
		if eq(subject, v) {
			return true
		}
	}
	return false
}

func MemorizedContains[T comparable](items []T) func(subject T) bool {
	m := map[T]bool{}
	return func(subject T) bool {
		if v, ok := m[subject]; ok {
			return v
		} else {
			v := Contains(items, subject)
			m[subject] = v
			return v
		}
	}
}

// Produces the set difference of two slices.
func Except[T comparable](a, b []T) []T {
	d := make([]T, 0)
	// Memoization is performed to reduce the amount of computation to some extent.
	containsB := MemorizedContains(b)
	for _, aa := range a {
		if !containsB(aa) {
			d = append(d, aa)
		}
	}
	return d
}

// Produces the set difference of two slices by any equality.
func ExceptAny[T any](a, b []T, eq func(a, b T) bool) []T {
	d := make([]T, 0)
	// I couldn't think of a good way to implement a memoized version of `ContainsAny()`.
	// In small slices, performance will not be a problem.
	for _, aa := range a {
		if !ContainsAny(b, aa, eq) {
			d = append(d, aa)
		}
	}
	return d
}

// Remove duplicates items
func DeDuping[T string | int](src []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range src {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Unique[T1 string | int](intSlice []T1) []T1 {
	keys := make(map[T1]bool)
	list := make([]T1, 0)
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

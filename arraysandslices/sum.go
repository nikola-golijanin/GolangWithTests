package arraysandslices

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(a, b int) int { return a + b }
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	length := len(numbersToSum)
	sums := make([]int, length)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}

		tail := x[1:]
		return append(acc, Sum(tail))
	}

	return Reduce(numbers, sumTail, []int{})
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Find[T any](collection []T, predicate func(T) bool) (T, bool) {
	for _, x := range collection {
		if predicate(x) {
			return x, true
		}
	}
	var zero T
	return zero, false
}

package chapter02

// InsertionSortDesc -
func InsertionSortDesc(input []int) ([]int, error) {
	for i := len(input) - 1; i >= 0; i-- {
		if i != 0 {
			for j := i - 1; j >= 0; j-- {
				if input[i] < input[j] {
					tmp := input[i]
					input[i] = input[j]
					input[j] = tmp
				}
			}
		}
	}
	return input, nil
}

// InsertionSortAsc -
func InsertionSortAsc(input []int) ([]int, error) {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] < input[j] {
				tmp := input[i]
				input[i] = input[j]
				input[j] = tmp
			}
		}
	}
	return input, nil
}

// SelectionSort -
func SelectionSort(input []int) ([]int, error) {
	for i := 0; i < len(input)-1; i++ {
		// find lowest
		n := int(^uint(0) >> 1)
		ij := len(input) - 1
		swap := false
		for j := i; j < len(input); j++ {
			if input[j] < n {
				n = input[j]
				ij = j
				swap = true
			}
		}
		// swap
		if swap {
			tmp := input[i]
			input[i] = input[ij]
			input[ij] = tmp
		}
	}
	return input, nil

}

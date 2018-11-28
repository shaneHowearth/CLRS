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

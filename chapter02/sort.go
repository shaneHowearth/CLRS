package chapter02

// InsertionSort -
func InsertionSort(input []int) ([]int, error) {
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

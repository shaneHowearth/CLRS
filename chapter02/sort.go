package chapter02

// BubbleSort -
func BubbleSort(input []int) ([]int, error) {
	for i := 0; i <= len(input); i++ {
		for j := 0; j < len(input)-(1+i); j++ {
			if input[j+1] < input[j] {
				tmp := input[j+1]
				input[j+1] = input[j]
				input[j] = tmp
			}
		}
	}
	return input, nil
}

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

// Divide and Conquer

// MergeSort -
func MergeSort(A []int) ([]int, error) {
	if len(A) < 1 {
		return A, nil
	}
	left := A[:len(A)/2]
	right := A[len(A)/2:]
	if len(left) != 1 {
		left, _ = MergeSort(left)
	}
	if len(right) != 1 {
		right, _ = MergeSort(right)
	}
	return merge(left, right)
}

func merge(l, r []int) ([]int, error) {
	var result []int
	var i, j int
	for {
		// stop the loop when either l or r are empty
		if i == len(l) || j == len(r) {
			break
		}
		if l[i] < r[j] {
			result = append(result, l[i])
			i++
		} else {
			result = append(result, r[j])
			j++
		}
	}
	if i < len(l) {
		result = append(result, l[i:]...)
	}
	if j < len(r) {
		result = append(result, r[j:]...)
	}
	return result, nil
}

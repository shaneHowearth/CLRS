package chapter02

func LinearSearch(v int, input []int) (int, error) {
	for i := 0; i < len(input); i++ {
		if input[i] == v {
			return i, nil
		}
	}
	return -1, nil
}

package array_slices

func SumSlices(list []int) int {

	sum := 0
	

	for i := range list {
		sum += list[i]
	}

	return sum
}

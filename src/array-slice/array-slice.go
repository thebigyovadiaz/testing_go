package array_slice

func Sum(arr []int) int {
	acc := 0
	for _, v := range arr {
		acc += v
	}

	return acc
}

func SumAll(numSliceAttr ...[]int) []int {
	var sliceResult []int

	for _, arr := range numSliceAttr {
		acc := Sum(arr)
		if acc > 0 {
			sliceResult = append(sliceResult, acc)
		}
	}

	return sliceResult
}

func SumAllTail(numSliceAttr ...[]int) []int {
	var newSlice []int
	for _, arr := range numSliceAttr {
		sumArr := 0
		if len(arr) > 0 {
			sumArr = Sum(arr[1:])
		}

		newSlice = append(newSlice, sumArr)
	}

	return newSlice
}

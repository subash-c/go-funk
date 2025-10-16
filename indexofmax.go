package funk

import "strings"

// IndexOfMaxInt validates the input, compares the elements and returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// This function implements the argMax functionality requested in GitHub issue #139.
// It accepts []int
// It returns int
func IndexOfMaxInt(i []int) int {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var max int
	var maxIndex int
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			max = item
			maxIndex = 0
			continue
		}
		if item > max {
			max = item
			maxIndex = idx
		}
	}
	return maxIndex
}

// IndexOfMaxInt8 validates the input, compares the elements and returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []int8
// It returns int
func IndexOfMaxInt8(i []int8) int {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var max int8
	var maxIndex int
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			max = item
			maxIndex = 0
			continue
		}
		if item > max {
			max = item
			maxIndex = idx
		}
	}
	return maxIndex
}

// IndexOfMaxInt16 validates the input, compares the elements and returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []int16
// It returns int
func IndexOfMaxInt16(i []int16) int {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var max int16
	var maxIndex int
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			max = item
			maxIndex = 0
			continue
		}
		if item > max {
			max = item
			maxIndex = idx
		}
	}
	return maxIndex
}

// IndexOfMaxInt32 validates the input, compares the elements and returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []int32
// It returns int
func IndexOfMaxInt32(i []int32) int {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var max int32
	var maxIndex int
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			max = item
			maxIndex = 0
			continue
		}
		if item > max {
			max = item
			maxIndex = idx
		}
	}
	return maxIndex
}

// IndexOfMaxInt64 validates the input, compares the elements and returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []int64
// It returns int
func IndexOfMaxInt64(i []int64) int {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var max int64
	var maxIndex int
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			max = item
			maxIndex = 0
			continue
		}
		if item > max {
			max = item
			maxIndex = idx
		}
	}
	return maxIndex
}

// IndexOfMaxFloat32 validates the input, compares the elements and returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []float32
// It returns int
func IndexOfMaxFloat32(i []float32) int {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var max float32
	var maxIndex int
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			max = item
			maxIndex = 0
			continue
		}
		if item > max {
			max = item
			maxIndex = idx
		}
	}
	return maxIndex
}

// IndexOfMaxFloat64 validates the input, compares the elements and returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []float64
// It returns int
func IndexOfMaxFloat64(i []float64) int {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var max float64
	var maxIndex int
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			max = item
			maxIndex = 0
			continue
		}
		if item > max {
			max = item
			maxIndex = idx
		}
	}
	return maxIndex
}

// IndexOfMaxString validates the input, compares the elements and returns the index of the maximum element in an array/slice.
// If there are duplicate occurrences of max element, only return the first one.
// It accepts []string
// It returns int
func IndexOfMaxString(i []string) int {
	if len(i) == 0 {
		panic("arg is an empty array/slice")
	}
	var max string
	var maxIndex int
	for idx := 0; idx < len(i); idx++ {
		item := i[idx]
		if idx == 0 {
			max = item
			maxIndex = 0
			continue
		}
		if compareStringsIndexOfMax(max, item) == item {
			max = item
			maxIndex = idx
		}
	}
	return maxIndex
}

// compareStringsIndexOfMax uses the strings.Compare method to compare two strings, and returns the greater one.
func compareStringsIndexOfMax(max, current string) string {
	r := strings.Compare(strings.ToLower(max), strings.ToLower(current))
	if r > 0 {
		return max
	}
	return current
}

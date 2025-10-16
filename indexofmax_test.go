package funk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIndexOfMaxIssue139Example tests the exact example provided in GitHub issue #139
func TestIndexOfMaxIssue139Example(t *testing.T) {
	nums := []int64{8, 3, 4, 44, 0}
	result := IndexOfMaxInt64(nums)
	assert.Equal(t, 3, result, "Should return index 3 for the example in issue #139")
	assert.Equal(t, int64(44), nums[result], "The element at returned index should be 44")
}

func TestIndexOfMaxWithArrayNumericInput(t *testing.T) {
	// Test Data
	d1 := []int{8, 3, 4, 44, 0}
	d1dup := []int{44, 3, 4, 44, 0} // duplicate max at beginning
	n1 := []int{}

	// Calls
	r1 := IndexOfMaxInt(d1)
	r1dup := IndexOfMaxInt(d1dup)

	// Assertions
	assert.Equal(t, 3, r1, "It should return the index of max value in array")
	assert.Equal(t, 0, r1dup, "It should return the first index of duplicate max value")
	assert.Panics(t, func() { IndexOfMaxInt(n1) }, "It should panic")
}

func TestIndexOfMaxWithArrayFloatInput(t *testing.T) {
	// Test Data
	d1 := []float64{2, 38.3, 4, 4.4, 4}
	d1dup := []float64{38.3, 2, 4, 38.3, 4} // duplicate max at beginning
	n1 := []float64{}

	// Calls
	r1 := IndexOfMaxFloat64(d1)
	r1dup := IndexOfMaxFloat64(d1dup)

	// Assertions
	assert.Equal(t, 1, r1, "It should return the index of max value in array")
	assert.Equal(t, 0, r1dup, "It should return the first index of duplicate max value")
	assert.Panics(t, func() { IndexOfMaxFloat64(n1) }, "It should panic")
}

func TestIndexOfMaxSingleElement(t *testing.T) {
	// Test Data with single elements
	d1 := []int{42}
	d2 := []float64{3.14}
	d3 := []string{"hello"}

	// Calls
	r1 := IndexOfMaxInt(d1)
	r2 := IndexOfMaxFloat64(d2)
	r3 := IndexOfMaxString(d3)

	// Assertions
	assert.Equal(t, 0, r1, "Single element should return index 0")
	assert.Equal(t, 0, r2, "Single element should return index 0")
	assert.Equal(t, 0, r3, "Single element should return index 0")
}

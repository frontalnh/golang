package golang_test

import (
	"testing"

	"github.com/tj/assert"
)

func TestFor(t *testing.T) {
	arr := []int{2, 3, 4, 5, 6}

	for i := 0; i < len(arr); i++ {
		if arr[i] < 2 {
			t.Error("Every array should be smaller than 2")
		}
	}

	for i, v := range arr {
		assert.Equal(t, v, arr[i])
	}
}

func TestSlice(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, []int{1, 2, 3}, arr[1:4])
	assert.Equal(t, []int{0, 1, 2, 3, 4}, arr[:5])
}

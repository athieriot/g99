package exercise

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {

	assert.Equal(t, []int{}, Flatten(nil))
	assert.Equal(t, []int{}, Flatten([][]int{{}}))
	assert.Equal(t, []int{5}, Flatten([][]int{{5}}))
	assert.Equal(t, []int{1,9,7, 8, 9}, Flatten([][]int{{1,9},{7}, {8,9}}))
}
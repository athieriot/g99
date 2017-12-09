package exercise

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {

	assert.Nil(t, Reverse(nil))
	assert.Equal(t, []int{}, Reverse([]int{}))
	assert.Equal(t, []int{5}, Reverse([]int{5}))
	assert.Equal(t, []int{7,9,1}, Reverse([]int{1,9,7}))
}
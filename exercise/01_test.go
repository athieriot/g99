package exercise

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLast(t *testing.T) {

	assert.Equal(t, Last(nil), 0)
	assert.Equal(t, Last([]int{}), 0)
	assert.Equal(t, Last([]int{5}), 5)
	assert.Equal(t, Last([]int{1,9,7}), 7)
}
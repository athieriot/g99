package exercise

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {

	assert.Equal(t, Length(nil), 0)
	assert.Equal(t, Length([]int{}), 0)
	assert.Equal(t, Length([]int{5}), 1)
	assert.Equal(t, Length([]int{1,9,7}), 3)
}
package exercise

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPenultimate(t *testing.T) {

	assert.Equal(t, Penultimate(nil), -1)
	assert.Equal(t, Penultimate([]int{}), -1)
	assert.Equal(t, Penultimate([]int{5}), -1)
	assert.Equal(t, Penultimate([]int{1,9,7}), 9)
}
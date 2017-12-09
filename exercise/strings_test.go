package exercise

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {

	assert.Equal(t, len(WordCount("This is great")), 3)
}
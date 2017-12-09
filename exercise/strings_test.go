package exercise

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	assert.Equal(t, len(WordCount("This is great")), 3)
}

func TestTrim(t *testing.T) {
	assert.Equal(t, "Awesome", Trim("   Awesome   "))
}
package design

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSingleton(t *testing.T) {
	s1 := NewSingleton()
	s2 := NewSingleton()
	s3 := NewSingleton()

	assert.Equal(t, s1, s2)
	assert.Equal(t, s2, s3)
}

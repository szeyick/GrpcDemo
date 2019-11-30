package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModel(t *testing.T) {
	m := &Impl{}
	assert.Equal(t, 5, m.Bar(5))
	if m.Bar(5) != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", 5, 5)
	}
}

package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestModel(t *testing.T) {
	m := &Impl{}
	assert.Equal(t, 5, m.Bar(5))
}

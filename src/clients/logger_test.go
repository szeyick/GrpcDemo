package clients

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {
	desc := "validate that a new logger instance is returned"
	logger := NewLogger()
	assert.NotNil(t, logger, desc)
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
	main()
}

func TestSomething(t *testing.T) {
	tests := []struct {
		desc string
	}{
		{
			desc: "this is a test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			total := Sum(5, 5)
			assert.Equal(t, 10, total)
		})
	}
}

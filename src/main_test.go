package src

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	tests := []struct{
		desc string
	}{
		{
			desc: "this is a test",
		},
	}

	for _, tt := range tests {
		t.Parallel()
		t.Run(tt.desc, func(t *testing.T) {
			var blah = &Test{}
			assert.Equal(t, "hello", blah.hello())
		})
	}
}

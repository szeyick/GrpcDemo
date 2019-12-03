package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		desc          string
		valueA        int
		valueB        int
		expectedTotal int
	}{
		{
			desc:          "validate that 1 + 1 = 2",
			valueA:        1,
			valueB:        1,
			expectedTotal: 2,
		},
		{
			desc:          "validate that 2 + 2 = 4",
			valueA:        2,
			valueB:        2,
			expectedTotal: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			actualTotal := Sum(tt.valueA, tt.valueB)
			assert.Equal(t, tt.expectedTotal, actualTotal)
		})
	}
}

func TestDummyMain(t *testing.T) {
	// This test does not test anything other than execute main()
	t.Parallel()
	main()
}

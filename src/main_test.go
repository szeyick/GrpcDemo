package main

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/szeyick/GrpcDemo/src/mocks"
	"github.com/szeyick/GrpcDemo/src/models"
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

			b := &models.Impl{}
			assert.Equal(t, 5, b.Bar(5))

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			model := mocks.NewMockFoo(mockCtrl)
			model.EXPECT().Bar(gomock.Any()).Return(1)

			assert.Equal(t, 1, model.Bar(5))
		})
	}
}

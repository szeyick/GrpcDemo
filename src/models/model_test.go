package models

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/szeyick/GrpcDemo/src/mocks"
	"testing"
)

func TestModel(t *testing.T) {
	m := &Impl{}
	assert.Equal(t, 5, m.Bar(5))
}

func TestMockModel(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	model := mocks.NewMockFoo(mockCtrl)
	model.EXPECT().Bar(gomock.Any()).Return(1)

	assert.Equal(t, 1, model.Bar(5))
}

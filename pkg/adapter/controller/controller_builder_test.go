package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewControllerBuilder(t *testing.T) {
	builder := NewControllerBuilder()

	assert.NotNil(t, builder)
	assert.Empty(t, builder.controllers)
}

func TestAddControllerBuilder(t *testing.T) {
	builder := NewControllerBuilder()

	baseController := &BaseController{}

	builder.Add(baseController)

	assert.Contains(t, builder.controllers, baseController)
}

func TestBuildControllerBuilder(t *testing.T) {
	controller1 := &BaseController{}
	controller2 := &BaseController{}

	finalController := NewControllerBuilder().Add(controller1).Add(controller2).Build()

	assert.Equal(t, controller1, finalController)
	baseController, ok := finalController.(*BaseController)
	assert.True(t, ok)
	assert.Equal(t, controller2, baseController.next)
}

package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBaseController(t *testing.T) {
	controller := NewBaseController()

	assert.NotNil(t, controller)
}

func TestErrorBaseController(t *testing.T) {
	controller := NewBaseController()

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

func TestOneValidPassBaseController(t *testing.T) {
	controller := NewBaseController()

	controller.SetNext(NewBaseController())

	err := controller.Handle(context.Background())

	assert.Error(t, err)
}

package controller

import (
	"testing"

	"github.com/atorrescogollo/terraform-cascade/internal/usecases"
	"github.com/stretchr/testify/assert"
)

func TestHandleCascade(t *testing.T) {
	c := NewCascadeController(
		usecases.NewMockRunRecursiveTerraformUseCase(),
	)
	err := c.HandleCascade()
	assert.EqualError(t, err, "not implemented")
}

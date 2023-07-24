package cmd

import (
	"github.com/atorrescogollo/terraform-cascade/internal/orchestration"
	"github.com/atorrescogollo/terraform-cascade/internal/usecases"
)

// Orchestration
var projectExecutor = orchestration.NewTerraformProjectExecutorWithOS()
var projectOrderResolver = orchestration.NewProjectOrderResolver()

// Use cases
var runRawTerraformUseCase = usecases.NewRunRawTerraformUseCase()
var runRecursiveTerraformUseCase = usecases.NewRunRecursiveTerraformUseCase(
	*projectExecutor,
	*projectOrderResolver,
)

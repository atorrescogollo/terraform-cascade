package cmd

import (
	"github.com/atorrescogollo/terraform-cascade/internal/orchestration"
	"github.com/atorrescogollo/terraform-cascade/internal/usecases"
)

// Orchestration
var projectExecutor = orchestration.NewTerraformProjectExecutorWithOS()
var projectOrderResolver = orchestration.NewProjectOrderResolver()

// Use cases
var runRawTerraformUseCase usecases.RunRawTerraformUseCase = usecases.NewRunRawTerraformUseCaseImpl()
var runRecursiveTerraformUseCase usecases.RunRecursiveTerraformUseCase = usecases.NewRunRecursiveTerraformUseCaseImpl(
	*projectExecutor,
	*projectOrderResolver,
)

package handlers

import (
	"github.com/gaia-pipeline/gaia/providers"
	"github.com/gaia-pipeline/gaia/providers/pipelines"
	"github.com/gaia-pipeline/gaia/providers/workers"
	"github.com/gaia-pipeline/gaia/security"
	"github.com/gaia-pipeline/gaia/security/rbac"
	"github.com/gaia-pipeline/gaia/store"
	"github.com/gaia-pipeline/gaia/workers/pipeline"
	"github.com/gaia-pipeline/gaia/workers/scheduler/service"
)

// Dependencies define dependencies for this service.
type Dependencies struct {
	Scheduler        service.GaiaScheduler
	PipelineService  pipeline.Servicer
	PipelineProvider pipelines.PipelineProviderer
	UserProvider     providers.UserProvider
	RBACProvider     providers.RBACProvider
	WorkerProvider   workers.WorkerProviderer
	Certificate      security.CAAPI
	RBACService      rbac.Service
	Store            store.GaiaStore
}

// GaiaHandler defines handler functions throughout Gaia.
type GaiaHandler struct {
	deps Dependencies
}

// NewGaiaHandler creates a new handler service with the required dependencies.
func NewGaiaHandler(deps Dependencies) *GaiaHandler {
	return &GaiaHandler{deps: deps}
}

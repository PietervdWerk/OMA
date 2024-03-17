package contract

import (
	"context"
	"oma/internal/playgroundlogs"
	"oma/models"
)

type App interface {
	Eval(ctx context.Context, req *models.EvalRequest) (*models.EvalResponse, error)
	Format(ctx context.Context, req *models.FormatRequest) (*models.FormatResponse, error)
	PlaygroundLogs(ctx context.Context) ([]playgroundlogs.PlaygroundLog, error)
}
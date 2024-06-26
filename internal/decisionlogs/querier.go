// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package decisionlogs

import (
	"context"
)

type Querier interface {
	CreateDecisionLog(ctx context.Context, arg CreateDecisionLogParams) (DecisionLog, error)
	GetDecisionLog(ctx context.Context, decisionID string) (DecisionLog, error)
	ListDecisionLogs(ctx context.Context) ([]DecisionLog, error)
}

var _ Querier = (*Queries)(nil)

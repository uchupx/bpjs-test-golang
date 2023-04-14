package data

import (
	"context"

	"github.com/uchupx/bpjs-test-golang/data/model"
)

type TransactionRepository interface {
	Inserts(ctx context.Context, items []model.Transaction) (*int64, error)
}

package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/uchupx/bpjs-test-golang/data/model"
)

type TransactionMySQL struct {
	DB *sqlx.DB
}

const bulkInsert = `
	INSERT INTO transactions(id, customer, quantity, price, timestamp) values %s;
`

func (m TransactionMySQL) Inserts(ctx context.Context, items []model.Transaction) (*int64, error) {
	if len(items) == 0 {
		return nil, nil
	}

	var args []interface{}

	placeholder := strings.TrimRight(strings.Repeat("(?,?,?,?,?),", len(items)), ",")
	query := fmt.Sprintf(bulkInsert, placeholder)

	for _, item := range items {
		args = append(args, item.Id)
		args = append(args, item.Customer)
		args = append(args, item.Quantity)
		args = append(args, item.Price)
		args = append(args, item.Timestamp)
	}

	stmt, err := m.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	affectedRows, err := rows.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &affectedRows, nil
}

func NewTransactionMysql(db *sqlx.DB) TransactionMySQL {
	return TransactionMySQL{
		DB: db,
	}
}

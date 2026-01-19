package db

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/rinnothing/golang-test-task/internal/model"
	"github.com/rinnothing/golang-test-task/internal/usecase/integer"
	"github.com/rinnothing/golang-test-task/pkg/transaction"
)

var _ integer.IntegerRepo = &postgresRepository{}

func (p *postgresRepository) AddInteger(ctx context.Context, num model.Integer) (txErr error) {
	var (
		tx  pgx.Tx
		err error
	)
	if tx, err = transaction.ExtractTx(ctx); err != nil {
		tx, err = p.db.Begin(ctx)
		if err != nil {
			return err
		}

		defer func() {
			if txErr != nil {
				_ = tx.Rollback(ctx)
				return
			}

			_ = tx.Commit(ctx)
		}()
	}

	const queryInteger = `
INSERT INTO integers (integer)
VALUES ($1);
`

	_, err = tx.Exec(ctx, queryInteger, num)
	if err != nil {
		return err
	}

	return nil
}

func (p *postgresRepository) ListIntegers(ctx context.Context) (retIntegers []model.Integer, txErr error) {
	var (
		tx  pgx.Tx
		err error
	)
	if tx, err = transaction.ExtractTx(ctx); err != nil {
		tx, err = p.db.Begin(ctx)
		if err != nil {
			return nil, err
		}

		defer func() {
			if txErr != nil {
				_ = tx.Rollback(ctx)
				return
			}

			_ = tx.Commit(ctx)
		}()
	}

	const queryIntegers = `
SELECT integer FROM integers`

	rows, err := tx.Query(ctx, queryIntegers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var integers []model.Integer
	for rows.Next() {
		var id model.Integer
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		integers = append(integers, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return integers, nil
}

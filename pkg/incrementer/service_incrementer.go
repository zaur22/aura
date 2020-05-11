package incrementer

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type (
	incrementer struct {
		db *sqlx.DB
	}
)

func newIncrementer(ctx context.Context, dto NewIncrementerDTO) (Incrementer, error) {
	var inc = &incrementer{
		db: dto.DBClient,
	}

	if err := inc.init(ctx); err != nil {
		return nil, fmt.Errorf("incrementer service init: %v", err)
	}

	return inc, nil
}

func (i *incrementer) init(ctx context.Context) error {
	tx, err := i.db.BeginTxx(
		ctx,
		&sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		})
	if err != nil {
		return fmt.Errorf("can't start transaction: %v", err)
	}
	defer tx.Rollback()

	var count int
	err = tx.GetContext(ctx, &count,
		"SELECT count(*) FROM increments")
	if err != nil {
		return fmt.Errorf("can't select rows count: %v", err)
	}

	if count == 1 {
		return nil
	} else if count > 1 {
		return fmt.Errorf("bad count of increment row, expected 1, got %v", count)
	}

	_, err = tx.ExecContext(ctx,
		`INSERT INTO increments (current_value, max_value, increment_step) VALUES ($1, $2, $3)`,
		0, 0, 1)
	if err != nil {
		return fmt.Errorf("can't insert init row: %v", err)
	}

	return tx.Commit()
}

func (i *incrementer) GetNumber(ctx context.Context) (uint64, error) {
	var incr = &increment{}

	err := i.db.GetContext(ctx, incr, "SELECT * FROM increments")
	if err != nil {
		return 0, fmt.Errorf("can't select increment row: %v", err)
	}

	return incr.CurrentValue, nil
}

func (i *incrementer) IncrementNumber(ctx context.Context) error {


	tx, err := i.db.BeginTxx(
		ctx,
		&sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
			ReadOnly:  false,
		})
	if err != nil {
		return fmt.Errorf("can't start transaction: %v", err)
	}
	defer tx.Rollback()

	var incr = &increment{}
	err = tx.GetContext(ctx, incr, "SELECT * FROM increments")
	if err != nil {
		return fmt.Errorf("can't select increment row: %v", err)
	}

	incr.CurrentValue = incr.CurrentValue + incr.IncrementStep
	if incr.MaxValue != 0 {
		incr.CurrentValue = incr.CurrentValue % incr.MaxValue
	}

	_, err = tx.ExecContext(ctx,
		`UPDATE increments SET current_value = $1`,
		incr.CurrentValue)
	if err != nil {
		return fmt.Errorf("can't update incremented value: %v", err)
	}

	return tx.Commit()
}

func (i *incrementer) SetSettings(ctx context.Context, dto SetSettingsDTO) error {
	_, err := i.db.ExecContext(ctx,
		`UPDATE increments SET max_value = $1, increment_step = $2`,
		dto.MaxValue,
		dto.IncrementStep)
	if err != nil {
		return fmt.Errorf("can't update config value: %v", err)
	}
	return nil
}

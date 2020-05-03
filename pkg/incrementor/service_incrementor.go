package incrementor

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type(
	entity struct {
		CurrentValue int64 `db:"current_value"`
		MaxValue int64	`db:"max_value"`
		IncrementStep int64 `db:"increment_step"`
	}

	incrementer struct {
		db *sqlx.DB
	}
)

func newIncrementer(dto NewIncrementerDTO) (Incrementer, error) {
	var inc = &incrementer{
		db: dto.DB,
	}

	if err := inc.init(); err != nil{
		return nil, err
	}

	return inc, nil
}

func (inc *incrementer) init() error {
	panic("not implemented")
}

func (inc *incrementer) GetNumber(context.Context) (int64, error){
	panic("not implemented")
}

func (inc *incrementer)  IncrementNumber(context.Context) error{
	panic("not implemented")
}

func (inc *incrementer) SetSettings(context.Context, SetSettingsDTO) error{
	panic("not implemented")
}



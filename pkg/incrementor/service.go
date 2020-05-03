package incrementor

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type (

	//Incrementer сервис для инкрементироания значения числовой синглтон-переменной
	//на заданную величину. Причём значение увеличивается циклически по модулю до
	//заданного максимального значения. Изначально шаг инкерементрования составляет 1,
	//а максимальное значение не задано. Сервис хранит всё своё состояние в БД.
	Incrementer interface {

		//GetNumber возвращает текущее значение переменной
		GetNumber(context.Context) (int64, error)

		//IncrementNumber инкрементирует значение переменной
		IncrementNumber(context.Context) error

		//SetSettings устанавливает параметры инкрементирования
		SetSettings(context.Context, SetSettingsDTO) error
	}

	SetSettingsDTO struct {
		//IncrementStep значение шага инкрементрования
		IncrementStep uint64

		//MaxValue максимальное значение переменной, до которой
		//будет увеличиваться циклически по модулю.
		//Если MaxValue == 0, считается, что максимальное значение не задано
		MaxValue      uint64
	}

	NewIncrementerDTO struct {
		DB *sqlx.DB
	}
)

//NewIncrementer конструктор нового сервиса Incrementor
func NewIncrementer(dto NewIncrementerDTO) (Incrementer, error) {
	panic("not implemented")
}

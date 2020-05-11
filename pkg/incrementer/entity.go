package incrementer

type increment struct {
	ID            uint32 `json:"id" db:"id"`
	CurrentValue  uint64 `json:"current_value" db:"current_value,use_zero"`
	MaxValue      uint64 `json:"max_value" db:"max_value,use_zero"`
	IncrementStep uint64 `json:"increment_step" db:"increment_step,use_zero"`
}

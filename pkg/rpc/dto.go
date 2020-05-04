package rpc

import "github.com/zaur22/aura/pkg/incrementer"

type (
	NewServerDTO struct {
		Incrementer incrementer.Incrementer
	}
)

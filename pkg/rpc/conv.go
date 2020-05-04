package rpc

import (
	"github.com/zaur22/aura/pkg/incrementer"
	"github.com/zaur22/aura/pkg/rpc/api"
)

func convSetSettings(in *api.SetSettingsRequest) incrementer.SetSettingsDTO{
	return incrementer.SetSettingsDTO{
		IncrementStep: (*in).IncrementStep,
		MaxValue:      (*in).MaxValue,
	}
}
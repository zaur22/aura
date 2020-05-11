package rpc

import (
	"context"
	"github.com/zaur22/aura/pkg/incrementer"
	"github.com/zaur22/aura/pkg/rpc/api"
)

type rpc struct {
	incrementer incrementer.Incrementer
}

func newIncrementerServer(incr incrementer.Incrementer) api.IncrementerServer {
	return &rpc{
		incrementer: incr,
	}
}

func (r *rpc) GetNumber(
	ctx context.Context,
	_ *api.GetNumberRequest) (*api.GetNumberResponse, error) {

	v, err := r.incrementer.GetNumber(ctx)
	if err != nil {
		return nil, err
	}

	var resp = &api.GetNumberResponse{
		Value: v,
	}
	return resp, nil
}

func (r *rpc) IncrementNumber(
	ctx context.Context,
	_ *api.IncrementNumberRequest) (*api.IncrementNumberResponse, error) {

	err := r.incrementer.IncrementNumber(ctx)
	if err != nil {
		return nil, err
	}

	return &api.IncrementNumberResponse{}, nil

}

func (r *rpc) SetSettings(
	ctx context.Context,
	in *api.SetSettingsRequest) (*api.SetSettingsResponse, error) {

	err := r.incrementer.SetSettings(
		ctx,
		convSetSettings(in))
	if err != nil {
		return nil, err
	}

	return &api.SetSettingsResponse{}, nil
}

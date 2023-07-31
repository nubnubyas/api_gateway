// Code generated by Kitex v0.6.0. DO NOT EDIT.

package calculatorservice

import (
	"context"
	calculator "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/calculator"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Calculate(ctx context.Context, request *calculator.CalculatorReq, callOptions ...callopt.Option) (r *calculator.CalculatorResp, err error)
	CapCalculate(ctx context.Context, request *calculator.CapCalculatorReq, callOptions ...callopt.Option) (r *calculator.CapCalculatorResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCalculatorServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCalculatorServiceClient struct {
	*kClient
}

func (p *kCalculatorServiceClient) Calculate(ctx context.Context, request *calculator.CalculatorReq, callOptions ...callopt.Option) (r *calculator.CalculatorResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Calculate(ctx, request)
}

func (p *kCalculatorServiceClient) CapCalculate(ctx context.Context, request *calculator.CapCalculatorReq, callOptions ...callopt.Option) (r *calculator.CapCalculatorResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CapCalculate(ctx, request)
}

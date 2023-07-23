package main

import (
	"context"
	"fmt"

	calculator "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/calculator"
)

type CalculatorApiImpl struct{}

type CalculatorResp struct {
	Message string `json:"message"`
}

func (s *CalculatorApiImpl) CapCalculate(ctx context.Context, req *calculator.CapCalculatorReq) (resp *calculator.CapCalculatorResp, err error) {

	// Perform the calculation
	var result float64
	gradesList := req.Num1

	for _, grade := range gradesList {
		result += grade
	}

	result /= float64(len(gradesList))

	return &calculator.CapCalculatorResp{
		Message: result,
	}, nil

}

func (s *CalculatorApiImpl) Calculate(ctx context.Context, req *calculator.CalculatorReq) (resp *calculator.CalculatorResp, err error) {
	// Parse the request into a CalculatorReq struct

	// Perform the calculation
	var result int64
	switch req.Operation {
	case "add":
		result = req.Num1 + req.Num2
	case "subtract":
		result = req.Num1 - req.Num2
	case "multiply":
		result = req.Num1 * req.Num2
	case "divide":
		if req.Num2 == 0 {
			return &calculator.CalculatorResp{
				Message: "division by zero",
			}, nil
		}
		result = req.Num1 / req.Num2
	default:
		return &calculator.CalculatorResp{
			Message: fmt.Sprintf("unknown operation: %s", req.Operation),
		}, nil
	}

	// Create the response
	calculatorResp := fmt.Sprintf("%s %d and %d is %d", req.Operation, req.Num1, req.Num2, result)
	return &calculator.CalculatorResp{
		Message: calculatorResp,
	}, nil
}

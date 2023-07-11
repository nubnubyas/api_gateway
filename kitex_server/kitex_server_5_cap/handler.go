package main

import (
	"context"
	"fmt"

	calculator "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/calculator"
	calculatorservice "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/calculator/calculatorservice"
	grader "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/grader"
)

// UniversityGradesImpl implements the last service interface defined in the IDL.
type UniversityGradesImpl struct{}

// GetGrades implements the UniversityGradesImpl interface.
func (s *UniversityGradesImpl) GetGrades(ctx context.Context, req *grader.GetCapRequest) (resp *grader.GetCapResponse, err error) {
	// TODO: Your code here...

	cap := []int64{}
	// placeholder grades
	grades := []string{"A", "A", "A", "B", "A"}
	for _, grade := range grades {
		switch grade {
		case "A":
			cap = append(cap, 5)
		case "B":
			cap = append(cap, 4)
		case "C":
			cap = append(cap, 3)
		case "D":
			cap = append(cap, 2)
		case "E":
			cap = append(cap, 1)
		case "F":
			cap = append(cap, 0)
		}
	}

	// Perform the calculation
	fmt.Println(cap)
	calReq := new(calculator.CapCalculatorReq)
	calReq.Num1 = cap
	fmt.Println(calReq)
	calcCli, err := calculatorservice.NewClient("calculator")
	// err : service discovery error: internal exception: no resolver available
	if err != nil {
		fmt.Println(err)
		println("error1")
		panic(err)
	}
	calResp, err1 := calcCli.CapCalculate(ctx, calReq)
	if err1 != nil {
		fmt.Println(err1)
		println("error2")
		panic(err1)
	}

	// final is an int 64
	final := calResp.Message

	var response grader.GetCapResponse
	response.Cap = float64(final)
	response.Name = "Testname123"
	response.Major = "Business"
	response.Gender = "Male"
	response.Id = 4

	return &response, nil

	// resp : &grader.GetCapResponse{
	// 	Id:     3,
	// 	Name:   "Testname",
	// 	Major:  "Business",
	// 	Gender: "Male",
	// 	Cap:    final,
	// }, nil
}

// InsertGrades implements the UniversityGradesImpl interface.
func (s *UniversityGradesImpl) InsertGrades(ctx context.Context, req *grader.InsertGradeRequest) (resp *grader.InsertGradeResponse, err error) {
	// TODO: Your code here...
	return &grader.InsertGradeResponse{
		Ok: true,
	}, nil

}

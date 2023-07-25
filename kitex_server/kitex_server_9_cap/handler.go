package main

import (
	"context"
	"fmt"
	"strings"

	database "github.com/cloudwego/api_gateway/kitex_server"
	calculator "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/calculator"
	calculatorservice "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/calculator/calculatorservice"
	grader "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/grader"
	registerCenter "github.com/cloudwego/api_gateway/register_center/shared"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
)

// UniversityGradesImpl implements the last service interface defined in the IDL.
type UniversityGradesImpl struct{}

// GetGrades implements the UniversityGradesImpl interface.
func (s *UniversityGradesImpl) GetGrades(ctx context.Context, req *grader.GetCapRequest) (resp *grader.GetCapResponse, err error) {
	num := fmt.Sprintf("%d", req.StudentId)
	exist, _ := database.NumExists(num)
	if !exist {
		return &grader.GetCapResponse{}, nil
	} else {
		cap := []float64{}
		id := int(req.StudentId)
		response, gradesInString, _ := database.GetGradesDB(id)
		grades := strings.Split(gradesInString, ",")

		for _, grade := range grades {
			switch grade {
			case "A+":
				cap = append(cap, 5)
			case "A":
				cap = append(cap, 5)
			case "A-":
				cap = append(cap, 4.5)
			case "B+":
				cap = append(cap, 4)
			case "B":
				cap = append(cap, 3.5)
			case "B-":
				cap = append(cap, 3)
			case "C+":
				cap = append(cap, 2.5)
			case "C":
				cap = append(cap, 2)
			case "D+":
				cap = append(cap, 1.5)
			case "D":
				cap = append(cap, 1)
			case "F":
				cap = append(cap, 0)
			}
		}

		calReq := new(calculator.CapCalculatorReq)
		calReq.Num1 = cap

		loadbalanceropt := client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer())
		calcCli, err := calculatorservice.NewClient("calculator",
			client.WithResolver(registerCenter.NacosResolver),
			loadbalanceropt)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		calResp, err1 := calcCli.CapCalculate(ctx, calReq)
		if err1 != nil {
			fmt.Println(err1)
			panic(err1)
		}

		response.Cap = calResp.Message
		return response, nil
	}
}

// InsertGrades implements the UniversityGradesImpl interface.
func (s *UniversityGradesImpl) InsertGrades(ctx context.Context, req *grader.InsertGradeRequest) (resp *grader.InsertGradeResponse, err error) {
	num := fmt.Sprintf("%d", req.StudentId)
	exist, _ := database.NumExists(num)
	if !exist {
		return &grader.InsertGradeResponse{
			Ok:  false,
			Msg: "the student id don't exist",
		}, nil
	} else {
		database.InsertGradesDB(req)

		return &grader.InsertGradeResponse{
			Ok: true,
		}, nil
	}
}

package main

import (
	"context"

	api "github.com/cloudwego/api_gateway/kitex_server/kitex_server_3_student/kitex_gen/api"
)

// StudentApiImpl implements the last service interface defined in the IDL.
type StudentApiImpl struct{}

type StudentInfo struct {
	Num    string
	Name   string
	Gender string
}

var StudentData = make(map[string]StudentInfo, 5)

// QueryStudent implements the StudentApiImpl interface.
func (s *StudentApiImpl) QueryStudent(ctx context.Context, req *api.QueryStudentRequest) (resp *api.QueryStudentResponse, err error) {
	// TODO: Your code here...
	stu, exist := StudentData[req.Num]
	if !exist {
		return &api.QueryStudentResponse{
			Msg: "Student does not exist",
		}, nil
	}

	resp = &api.QueryStudentResponse{
		Num:    stu.Num,
		Name:   stu.Name,
		Gender: stu.Gender,
		Msg:    "Student exist",
	}

	println(StudentData)
	return resp, nil
}

// InsertStudent implements the StudentApiImpl interface.
func (s *StudentApiImpl) InsertStudent(ctx context.Context, req *api.InsertStudentRequest) (resp *api.InsertStudentResponse, err error) {
	// TODO: Your code here...
	_, exist := StudentData[req.Num]
	if exist {
		return &api.InsertStudentResponse{
			Ok:  false,
			Msg: "the num has exists",
		}, nil
	}

	StudentData[req.Num] = StudentInfo{
		Num:    req.Num,
		Name:   req.Name,
		Gender: req.Gender,
	}

	return &api.InsertStudentResponse{
		Ok: true,
	}, nil
}

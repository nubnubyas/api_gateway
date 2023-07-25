package main

import (
	"context"

	database "github.com/cloudwego/api_gateway/kitex_server"
	api "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/api"
)

// StudentApiImpl implements the last service interface defined in the IDL.
type StudentApiImpl struct{}

// QueryStudent implements the StudentApiImpl interface.
func (s *StudentApiImpl) QueryStudent(ctx context.Context, req *api.QueryStudentRequest) (resp *api.QueryStudentResponse, err error) {
	exist, _ := database.NumExists(req.Id)
	if !exist {
		return &api.QueryStudentResponse{
			Msg: "Student does not exist in server 2",
		}, nil
	} else {
		resp, _ = database.QueryStudentDB(req.Id)
		resp.Msg = "Student exist in server 2"
		return resp, nil
	}
}

// InsertStudent implements the StudentApiImpl interface.
func (s *StudentApiImpl) InsertStudent(ctx context.Context, req *api.InsertStudentRequest) (resp *api.InsertStudentResponse, err error) {
	exist, _ := database.NumExists(req.Id)
	if exist {
		return &api.InsertStudentResponse{
			Ok:  false,
			Msg: "Student inserted already",
		}, nil
	} else {
		database.InsertStudentDB(req)

		return &api.InsertStudentResponse{
			Ok: true,
		}, nil
	}
}

package main

import (
	"context"
	api "github.com/cloudwego/api_gateway/kitex_gen/api"
)

// StudentApiImpl implements the last service interface defined in the IDL.
type StudentApiImpl struct{}

// QueryStudent implements the StudentApiImpl interface.
func (s *StudentApiImpl) QueryStudent(ctx context.Context, req *api.QueryStudentRequest) (resp *api.QueryStudentResponse, err error) {
	// TODO: Your code here...
	return
}

// InsertStudent implements the StudentApiImpl interface.
func (s *StudentApiImpl) InsertStudent(ctx context.Context, req *api.InsertStudentRequest) (resp *api.InsertStudentResponse, err error) {
	// TODO: Your code here...
	return
}

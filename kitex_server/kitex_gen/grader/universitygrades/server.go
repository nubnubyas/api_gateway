// Code generated by Kitex v0.6.1. DO NOT EDIT.
package universitygrades

import (
	grader "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/grader"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler grader.UniversityGrades, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
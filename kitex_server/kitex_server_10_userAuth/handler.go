package main

import (
	"context"

	userauth "github.com/cloudwego/api_gateway/kitex_server/kitex_gen/userAuth"
)

// UserProfileServiceImpl implements the last service interface defined in the IDL.
type authUserServiceImpl struct{}

// Authenticate implements the authUserServiceImpl interface
func (s *authUserServiceImpl) Authenticate(ctx context.Context, username string, password string) (resp string, err error) {

	return
}

// Authorize implements the authUserServiceImpl interface.
func (s *authUserServiceImpl) Authorize(ctx context.Context, token string) (resp bool, err error) {

	return
}

// GetProfile implements the authUserServiceImpl interface.
func (s *authUserServiceImpl) GetProfile(ctx context.Context, username string) (resp *userauth.UserProfile, err error) {
	// TODO: Your code here...
	return
}

// UpdateProfile implements the authUserServiceImpl interface.
func (s *authUserServiceImpl) UpdateProfile(ctx context.Context, username string, profile *userauth.UserProfile) (resp bool, err error) {
	// TODO: Your code here...
	return
}

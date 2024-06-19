package test10

import (
	"context"
	"fmt"
	test10 "test10/api"
	"test10/internal/transformer"
	"test10/pkg/ent/user"
)

func (s *test10Server) UserCreate(ctx context.Context, request *test10.UserCreateRequest) (*test10.UserCreateRespond, error) {
	// Kiểm tra tính hợp lệ của yêu cầu
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Tìm người dùng trong cơ sở dữ liệu bằng username
	exists, err := s.entClient.User.Query().
		Where(user.UsernameEQ(request.User.Username)).
		Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("username existed")
	}

	// Tạo người dùng mới
	user, err := s.entClient.User.Create().
		SetUsername(request.User.Username).
		SetPassword(request.User.Password).
		SetEmail(request.User.Email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	userinfo := transformer.ToUserInfoPb(user)

	return &test10.UserCreateRespond{
		Message: "Create User Successfully!!",
		User:    userinfo,
	}, nil
}

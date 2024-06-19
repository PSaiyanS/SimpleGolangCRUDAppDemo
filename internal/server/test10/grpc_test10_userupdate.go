package test10

import (
	"context"
	"fmt"

	test10 "test10/api"
	"test10/internal/transformer"
)

func (s *test10Server) UserUpdate(ctx context.Context, request *test10.UserUpdateRequest) (*test10.UserUpdateRespond, error) {
	// Kiểm tra tính hợp lệ của yêu cầu
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Tìm người dùng trong cơ sở dữ liệu bằng ID
	existingUser, err := s.entClient.User.Get(ctx, request.User.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	// Lưu thông tin người dùng đã cập nhật vào cơ sở dữ liệu
	updatedUser, err := existingUser.Update().SetUsername(request.User.Username).SetPassword(request.User.Password).SetEmail(request.User.Email).Save(ctx)
	fmt.Println("existing user:", existingUser)
	fmt.Println("Update user:", updatedUser)

	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Chuyển đổi thông tin người dùng sang định dạng protobuf
	userInfo := transformer.ToUserInfoPb(updatedUser)

	// Trả về phản hồi
	return &test10.UserUpdateRespond{
		Message: "User Update Successfully! ",
		User:    userInfo,
	}, nil
}

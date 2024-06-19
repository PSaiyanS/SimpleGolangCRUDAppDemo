package test10

import (
	"context"

	test10 "test10/api"
)

func (s *test10Server) UserDelete(ctx context.Context, request *test10.UserDeleteRequest) (*test10.UserDeleteRespond, error) {
	// Kiểm tra tính hợp lệ của yêu cầu
	if err := request.Validate(); err != nil {
		return nil, err
	}
	// Lấy ID
	id := request.GetId()
	// Xóa người dùng
	err := s.entClient.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}
	// Trả về phản hồi
	return &test10.UserDeleteRespond{Message: "User Delete Successfullly!! "}, nil
}

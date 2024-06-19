package test10

import (
	"context"

	test10 "test10/api"
	"test10/internal/transformer"
)

func (s *test10Server) UserGetById(ctx context.Context, request *test10.UserGetByIdRequest) (*test10.UserGetByIdRespond, error) {
	// Kiểm tra tính hợp lệ của yêu cầu
	if err := request.Validate(); err != nil {
		return nil, err
	}
	//Lấy id
	id := request.GetId()
	// Truy vấn người dùng có id từ cơ sở dữ liệu
	user, err := s.entClient.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// Chuyển đổi người dùng định dạng protobuf
	userinfo := transformer.ToUserInfoPb(user)
	// Trả về phản hồi
	return &test10.UserGetByIdRespond{User: userinfo}, nil
}

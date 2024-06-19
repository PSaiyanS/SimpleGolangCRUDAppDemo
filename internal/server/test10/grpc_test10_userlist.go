package test10

import (
	"context"

	test10 "test10/api"
	"test10/internal/transformer"
)

func (s *test10Server) UserList(ctx context.Context, request *test10.UserListRequest) (*test10.UserListRespond, error) {
	// Kiểm tra tính hợp lệ của yêu cầu
	if err := request.Validate(); err != nil {
		return nil, err
	}
	offset := request.GetOffset()
	limit := request.GetLimit()

	// Truy vấn tất cả người dùng từ cơ sở dữ liệu
	users, err := s.entClient.User.Query().Offset(int(offset)).Limit(int(limit)).All(ctx)
	if err != nil {
		return nil, err
	}
	//Đếm tổng tất cả user
	totalCount, err := s.entClient.User.Query().Count(ctx)
	if err != nil {
		return nil, err
	}

	// Chuyển đổi danh sách người dùng sang định dạng protobuf
	var userInfos []*test10.User
	for _, user := range users {
		userInfo := transformer.ToUserInfoPb(user)
		userInfos = append(userInfos, userInfo)
	}
	// Trả về phản hồi chứa danh sách người dùng
	return &test10.UserListRespond{
		Users:      userInfos,
		TotalCount: int64(totalCount),
	}, nil
}

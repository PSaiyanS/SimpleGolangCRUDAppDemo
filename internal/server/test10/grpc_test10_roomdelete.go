package test10

import (
	"context"

	test10 "test10/api"
)

func (s *test10Server) RoomDelete(ctx context.Context, request *test10.RoomDeleteRequest) (*test10.RoomDeleteRespond, error) {
	// Kiểm tra tính hợp lệ của yêu cầu
	if err := request.Validate(); err != nil {
		return nil, err
	}
	// Lấy ID
	id := request.GetId()
	// Xóa người dùng
	err := s.entClient.Room.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &test10.RoomDeleteRespond{Message: "Room Delete Successfully!! "}, nil
}

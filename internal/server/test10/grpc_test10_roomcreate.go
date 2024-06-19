package test10

import (
	"context"
	"fmt"

	test10 "test10/api"
	"test10/internal/transformer"
	"test10/pkg/ent/room"
)

func (s *test10Server) RoomCreate(ctx context.Context, request *test10.RoomCreateRequest) (*test10.RoomCreateRespond, error) {
	// Kiểm tra tính hợp lệ của yêu cầu
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Tìm người dùng trong cơ sở dữ liệu bằng name
	exists, err := s.entClient.Room.Query().
		Where(room.NameEQ(request.Room.Name)).
		Exist(ctx)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, fmt.Errorf("username existed")
	}

	// Tạo người dùng mới
	room, err := s.entClient.Room.Create().
		SetName(request.Room.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	roominfo := transformer.ToRoomInfoPb(room)

	return &test10.RoomCreateRespond{
		Message: "Create User Successfully!!",
		Room:    roominfo,
	}, nil
}

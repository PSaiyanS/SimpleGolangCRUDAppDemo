package test10

import (
	"context"
	"fmt"

	test10 "test10/api"
	"test10/internal/transformer"
)

func (s *test10Server) RoomUpdate(ctx context.Context, request *test10.RoomUpdateRequest) (*test10.RoomUpdateRespond, error) {
	// Kiểm tra tính hợp lệ của yêu cầu
	if err := request.Validate(); err != nil {
		return nil, err
	}

	// Tìm người dùng trong cơ sở dữ liệu bằng ID
	existingRoom, err := s.entClient.Room.Get(ctx, request.Room.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	// Lưu thông tin người dùng đã cập nhật vào cơ sở dữ liệu
	updatedRoom, err := existingRoom.Update().SetName(request.Room.Name).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Chuyển đổi thông tin người dùng sang định dạng protobuf
	roomInfo := transformer.ToRoomInfoPb(updatedRoom)

	// Trả về phản hồi
	return &test10.RoomUpdateRespond{
		Message: "User Update Successfully! ",
		Room:    roomInfo,
	}, nil

}

package test10

import (
	"context"

	test10 "test10/api"
	"test10/internal/transformer"
)

func (s *test10Server) RoomGetById(ctx context.Context, request *test10.RoomGetByIdRequest) (*test10.RoomGetByIdRespond, error) {
	//Lấy id
	id := request.GetId()
	// Truy vấn người dùng có id từ cơ sở dữ liệu
	room, err := s.entClient.Room.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// Chuyển đổi người dùng định dạng protobuf
	roominfo := transformer.ToRoomInfoPb(room)
	// Trả về phản hồi
	return &test10.RoomGetByIdRespond{Room: roominfo}, nil
}

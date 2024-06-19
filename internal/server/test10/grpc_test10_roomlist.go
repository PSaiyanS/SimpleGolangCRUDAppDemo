package test10

import (
	"context"

	test10 "test10/api"
	"test10/internal/transformer"
)

func (s *test10Server) RoomList(ctx context.Context, request *test10.RoomListRequest) (*test10.RoomListRespond, error) {
	// Kiểm tra tính hợp lệ của yêu cầu
	if err := request.Validate(); err != nil {
		return nil, err
	}
	offset := request.GetOffset()
	limit := request.GetLimit()

	// Truy vấn tất cả người dùng từ cơ sở dữ liệu
	rooms, err := s.entClient.Room.Query().Offset(int(offset)).Limit(int(limit)).All(ctx)
	if err != nil {
		return nil, err
	}
	//Đếm tổng tất cả room
	totalCount, err := s.entClient.Room.Query().Count(ctx)
	if err != nil {
		return nil, err
	}

	// Chuyển đổi danh sách room sang định dạng protobuf
	var roomInfos []*test10.Room
	for _, room := range rooms {
		roomInfo := transformer.ToRoomInfoPb(room)
		roomInfos = append(roomInfos, roomInfo)
	}
	// Trả về phản hồi chứa danh sách room
	return &test10.RoomListRespond{
		Rooms:      roomInfos,
		TotalCount: int64(totalCount),
	}, nil
}

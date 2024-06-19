package transformer

import (
	test10 "test10/api"
	"test10/pkg/ent"
)

func ToUserInfoPb(info *ent.User) *test10.User {
	if info == nil {
		return nil
	}

	return &test10.User{
		Id:       info.ID,
		Username: info.Username,
		Password: info.Password,
		Email:    info.Email,
	}
}
func ToRoomInfoPb(info *ent.Room) *test10.Room {
	if info == nil {
		return nil
	}

	return &test10.Room{
		Id:   info.ID,
		Name: info.Name,
	}
}
func ToMessageInfoPb(info *ent.Message) *test10.Message {
	if info == nil {
		return nil
	}

	return &test10.Message{
		Id:      info.ID,
		Content: info.Content,
		UserId:  info.Edges.User.Username,
		RoomId:  info.Edges.Room.Name,
	}
}

package test10

import (
	test10 "test10/api"
	"test10/pkg/ent"
)

func NewServer(ent *ent.Client) test10.Test10Server {
	return &test10Server{
		entClient: ent,
	}
}

type test10Server struct {
	test10.UnimplementedTest10Server
	entClient *ent.Client
}

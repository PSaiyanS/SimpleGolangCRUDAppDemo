package server

import (
	"context"

	mykit "gitlab.ugaming.io/marketplace/mykit/pkg/api"
	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"

	//
	dbe "gitlab.ugaming.io/marketplace/database/pkg/ent"

	entcli "test10/pkg/ent"
	//

	pb0 "test10/api"
	"test10/internal/server/test10"
	config "test10/pkg/config"
)

// Serve ...
func Serve(cfg *config.Config) {
	service := newService(cfg, []mykit.Option{}...)
	// ent init
	drv, err := dbe.Open("mysql_test", cfg.GetDatabase())
	if err != nil {
		service.Logger().Fatal("could not connect to db", zap.Error(err))
	}
	ent := entcli.NewClient(entcli.Driver(drv))

	defer func() {
		if err := ent.Close(); err != nil {
			service.Logger().Fatal("could not close ent client", zap.Error(err))
		}
	}()
	if err = ent.Schema.Create(context.Background()); err != nil {
		service.Logger().Fatal("could not init database", zap.Error(err))
	}
	// grpc server
	server := service.Server()
	pb0.RegisterTest10Server(server, test10.NewServer(ent))

	// Register reflection service on gRPC server.
	// Please remove if you it's not necessary for your service
	reflection.Register(server)

	service.Serve()
}

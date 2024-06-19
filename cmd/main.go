package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gitlab.ugaming.io/marketplace/carbon/pkg/config"

	"test10/internal/server"
)

func main() {
	flags := config.ParseFlags()
	server.Run(flags)
}

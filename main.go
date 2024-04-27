package main

import (
	"github.com/Krittin-Khanueng/assessment-tax/config"
	"github.com/Krittin-Khanueng/assessment-tax/databases"
	"github.com/Krittin-Khanueng/assessment-tax/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)
	server.Start()
}

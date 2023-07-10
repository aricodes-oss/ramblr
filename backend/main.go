package main

import (
	"ramblr/db"
	"ramblr/logging"
	"ramblr/models"
	"ramblr/query"
	"ramblr/server"
)

var log = logging.Logger

func init() {
	query.SetDefault(db.Connection)
	db.Connection.AutoMigrate(models.AllModels...)
}

func main() {
	log.Info("Booting!")
	server.Init()
}

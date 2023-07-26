package main

import (
	"github.com/syamsv/apollo-server/common/cmd"
	"github.com/syamsv/apollo-server/common/database"
	"github.com/syamsv/apollo-server/common/migration"
	"github.com/syamsv/apollo-server/common/session"
	"github.com/syamsv/apollo-server/config"
)

func main() {
	config.LoadConfig()
	session.InitCacheManager()
	database.InitDatabaseInstannce()
	if config.MIGRATE {
		migration.Migrate()
	}
	cmd.StartServer()
}

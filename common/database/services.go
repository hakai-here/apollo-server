package database

import (
	"log"

	"github.com/syamsv/apollo-server/pkg/dbfunc/user"
)

var (
	User user.Interface
)

func InitDatabaseInstannce() {
	database, err := GetDbInstance()
	if err != nil {
		log.Fatal("error while connecting to database", err)
	}
	userRepo := user.NewRepository(database)
	User = user.NewService(userRepo)
}

package migration

import (
	"log"

	"github.com/syamsv/apollo-server/common/database"
	"github.com/syamsv/apollo-server/pkg/models"
)

func Migrate() {
	db, err := database.GetDbInstance()
	if err != nil {
		log.Fatalln(err.Error())
	}
	db.AutoMigrate(&models.Users{})
}

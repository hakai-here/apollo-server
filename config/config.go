package config

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/syamsv/apollo-server/common/utils"
)

var (
	SERVER_PORT    = ""
	MIGRATE        = false
	POSTGRES_USER  = ""
	POSTGRES_PASS  = ""
	POSTGRES_DB    = ""
	POSTGRES_HOST  = ""
	POSTGRES_PORT  = ""
	REDIS_HOST     = ""
	REDIS_PORT     = ""
	REDIS_PASSWORD = ""
	SMTP_SERVER    = ""
	SMTP_PORT      = 0
	SMTP_USER      = ""
	SMTP_PASSWORD  = ""
)

func LoadConfig() {
	utils.ImportEnv()
	SERVER_PORT = fmt.Sprintf(":%s", viper.GetString("SERVER_PORT"))
	MIGRATE = viper.GetBool("MIGRATE")
	POSTGRES_USER = viper.GetString("POSTGRES_USER")
	POSTGRES_PASS = viper.GetString("POSTGRES_PASS")
	POSTGRES_DB = viper.GetString("POSTGRES_DB")
	POSTGRES_HOST = viper.GetString("POSTGRES_HOST")
	POSTGRES_PORT = viper.GetString("POSTGRES_PORT")
	REDIS_HOST = viper.GetString("REDIS_HOST")
	REDIS_PORT = viper.GetString("REDIS_PORT")
	REDIS_PASSWORD = viper.GetString("REDIS_PASSWORD")
	SMTP_PASSWORD = viper.GetString("SMTP_PASSWORD")
	SMTP_SERVER = viper.GetString("SMTP_SERVER")
	SMTP_PORT = viper.GetInt("SMTP_PORT")
	SMTP_USER = viper.GetString("SMTP_USER")
}

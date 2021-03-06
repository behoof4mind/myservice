package environment

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type Environment struct {
	DB_URL      string
	DB_USERNAME string
	DB_PASSWORD string
}

var E Environment

func InitEnvVariables() {
	logrus.Info("Reading environment variables")
	var env_exist bool
	E.DB_URL, env_exist = os.LookupEnv("DB_URL")
	if env_exist != true {
		logrus.Warn("'DB_URL' does not exist. Will use `localhost:3306`")
	}
	E.DB_USERNAME, env_exist = os.LookupEnv("DB_USERNAME")
	if env_exist != true {
		logrus.Warn("'DB_USERNAME' does not exist. Will use `root`")
	}
	E.DB_PASSWORD, env_exist = os.LookupEnv("DB_PASSWORD")
	if env_exist != true {
		logrus.Warn("'DB_PASSWORD' does not exist. Will use defailt one")
	}
	fmt.Printf("\n######## Environment ########\nDB_URL:      %s\nDB_USERNAME: %s\n#############################\n\n", E.DB_URL, E.DB_USERNAME)
}

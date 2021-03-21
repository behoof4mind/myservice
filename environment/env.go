package environment

import (
  "fmt"
  "github.com/sirupsen/logrus"
  "os"
)

type Environment struct {
	DB_SERVER string
	DB_PORT   string
	DB_USERNAME string
	DB_PASSWORD string

}

var E Environment

func InitEnvVariables() {
  logrus.Info("Reading environment variables")
  var env_exist bool
  E.DB_SERVER, env_exist = os.LookupEnv("DB_SERVER")
  if env_exist != true {
    logrus.Warn("'DB_SERVER' does not exist. Will use `localhost`")
  }
  E.DB_PORT, env_exist = os.LookupEnv("DB_PORT")
  if env_exist != true {
    logrus.Warn("'DB_PORT' does not exist. Will use `3306`")
  }
  E.DB_USERNAME, env_exist = os.LookupEnv("DB_USERNAME")
  if env_exist != true {
    logrus.Warn("'DB_USERNAME' does not exist. Will use `root`")
  }
  E.DB_PASSWORD, env_exist = os.LookupEnv("DB_PASSWORD")
  if env_exist != true {
    logrus.Warn("'DB_PASSWORD' does not exist. Will use defailt one")
  }
  fmt.Printf("\n######## Environment ########\nDB_SERVER:   %s\nDB_PORT:     %s\nDB_USERNAME: %s\n#############################\n\n", E.DB_SERVER, E.DB_PORT, E.DB_USERNAME)
}

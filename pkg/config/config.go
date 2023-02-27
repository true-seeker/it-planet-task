package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("json")
	config.SetConfigName(env)
	config.AddConfigPath("configs/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
	}
	initFlags()
}

func GetConfig() *viper.Viper {
	return config
}

func initFlags() {
	dbAddr := flag.String("dbAddr", "", "database.address")
	dbUser := flag.String("dbUser", "", "database.user")
	dbPass := flag.String("dbPass", "", "database.password")
	dbName := flag.String("dbName", "", "database.dbname")
	dbPort := flag.String("dbPort", "", "database.port")
	srvAddr := flag.String("srvAddr", "", "server.address")
	srvPort := flag.String("srvPort", "", "server.port")

	flag.Parse()

	if *dbAddr != "" {
		config.Set("database.address", dbAddr)
	}
	if *dbUser != "" {
		config.Set("database.user", dbUser)
	}
	if *dbPass != "" {
		config.Set("database.password", dbPass)
	}
	if *dbName != "" {
		config.Set("database.dbname", dbName)
	}
	if *dbPort != "" {
		config.Set("database.port", dbPort)
	}
	if *srvAddr != "" {
		config.Set("server.address", srvAddr)
	}
	if *srvPort != "" {
		config.Set("server.port", srvPort)
	}
}

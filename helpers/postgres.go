package helpers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"it-planet-task/pkg/config"
	"it-planet-task/pkg/errorHandler"
)

var DB *gorm.DB

func BuildPostgresConnectionString() string {
	var ConnectionString = fmt.Sprintf("host=%s "+
		"user=%s "+
		"password=%s "+
		"dbname=%s "+
		"port=%s "+
		"sslmode=disable TimeZone=Asia/Yekaterinburg",
		config.GetConfig().Get("database.address"),
		config.GetConfig().Get("database.user"),
		config.GetConfig().Get("database.password"),
		config.GetConfig().Get("database.dbname"),
		config.GetConfig().Get("database.port"))
	return ConnectionString

}

func GetConnectionOrCreateAndGet() *gorm.DB {
	if DB != nil {
		return DB
	}

	db, err := gorm.Open(postgres.Open(BuildPostgresConnectionString()), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	errorHandler.FailOnError(err, "Failed to connect to DB")
	DB = db
	return db
}

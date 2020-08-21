package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres dialect
	"github.com/spf13/viper"
)

// InitDatabase for database initialization with specific dialect
func InitDatabase() (db *gorm.DB, err error) {
	dbDriver := viper.GetString("DB_DIALECT")
	var connectionString string

	if dbDriver == "postgres" {
		connectionString = buildPostgresConnectionString()
	}

	db, err = openConnection(dbDriver, connectionString)

	return
}

func openConnection(dbDriver string, connection string) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbDriver, connection)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func buildPostgresConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_POSTGRES_SSL_MODE"))

	return
}

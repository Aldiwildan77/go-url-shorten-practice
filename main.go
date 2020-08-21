package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/database"
	"url-shortener/middleware"

	"github.com/gorilla/handlers"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func onError(err error, message string) {
	if err != nil {
		log.Fatal(message)
		log.Fatal(err)
	}
}

func loadConfig() (err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()

	return
}

func runServer(db *gorm.DB) {
	router := LoadRouter(db)
	corsOptions := middleware.CorsMiddleware()
	addr := fmt.Sprintf("%s:%d", viper.GetString("HOSTNAME"), viper.GetInt("PORT"))

	if err := database.Migrate(db); err != nil {
		onError(err, "Failed to migrate database schema")
	}

	fmt.Printf("Server is running at http://%s\n", addr)
	if err := http.ListenAndServe(addr, handlers.CORS(corsOptions[0], corsOptions[1], corsOptions[2])(router)); err != nil {
		onError(err, "Failed to run Server")
	}
}

func main() {
	err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitDatabase()
	err = db.DB().Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Connected")
	defer db.Close()

	runServer(db)
}

package main

import (
	"crud-golang/src/controller/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading .env file: %v", err)
	}

	viper.AutomaticEnv()

	viper.SetDefault("PORT", 3000)

	port := viper.GetInt("PORT")
	fmt.Println("PORT:", port)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

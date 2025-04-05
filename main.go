package main

import (
	"fmt"
	"log"

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

}

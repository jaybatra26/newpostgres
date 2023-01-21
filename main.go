package main

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


type repository struct{
	DB *gorm.db
}

func main(){

	err := godotenv.Load(".env")

	if(err != nil){
		log.Fatal(err) 
	}
	app := fiber.New()
	r.setuproutes(app)


	app.Listen(8080)

}
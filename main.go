package main

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type book struct{
	Author string `json: "author`
	Title string  `json : "title"`
	Desc  string   `json :"desc"`
}

func (r * repository)createbook(c *fiber.ctx){
	book = book{}

	//body parser isiliye use hua hai kyuki golang struct jo banaaya h vo json form mien create kar rha hai 
	err := c.BodyParser(&book)

	if(err != nil){
		c.Status(http.BAdrequest)
	}

}


type repository struct{
	DB *gorm.db
}


func (r *repository)setuproutes(app *fiber.app){
	api := app.Group("/api")


	app.Post("/createbook",r.createbook)
	app.Get("/getbook",r.getbooks)
	app.Get("/getbook/:id",r.getbookbyid)
	app.Delete("deletebook",r.deletebook)
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
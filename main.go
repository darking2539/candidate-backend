package main

import (
	_"candidate-backend/db"
	"candidate-backend/middleware"
	"candidate-backend/router"
	_"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {

	godotenv.Load()
	port := os.Getenv("PORT")


	//initz gin
	engine := gin.Default()
	engine.Use(middleware.CORSMiddleware())
	router.Setup(engine)

	//initz DB
	// err := db.InitMongoDB()
	// if err != nil{
	// 	log.Fatal(err.Error())
	// }

	engine.Run(":" + port)

}

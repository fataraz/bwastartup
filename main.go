package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// connect to database mysql
	// dsn := "alesha:Alesha2021!@#@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println("Connection to database is good")

	// // get table users
	// var users []user.User
	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.ID)
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("============================================================")

	// }

	//router := gin.Default()
	//router.GET("/handler", Handler)
	//router.Run()

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// connect to database mysql
	dsn := "alesha:Alesha2021!@#@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser) // handler->user.go
	router.Run()

	//userInput := user.RegisterUserInput{}
	//userInput.Name = "Test simpan dari service"
	//userInput.Email = "contoh@gmail.com"
	//userInput.Occupation = "Anak band"
	//userInput.Password = "password"
	//userService.RegisterUser(userInput)

	//user := user.User{
	//	Name: "Test Simpan",
	//}
	//userRepository.Save(user)

	// input dari user
	// handler mapping input dari user -> struct input
	// service : melakukan mapping dari struct input ke struct User
	// repository save struct User ke db
	// db

}

func Handler(c *gin.Context) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// connect to database mysql
	dsn := "alesha:Alesha2021!@#@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)

	// input dari user
	// handler mapping input ke struct
	// service : melakukan mapping dari struct input ke struct User
	// repository save struct User ke db
	// db

}

package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// connect to database mysql
	dsn := "alesha:Alesha2021!@#@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// Repository
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	// Service
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	campaingService := campaign.NewService(campaignRepository)

	router := gin.Default()
	api := router.Group("/api/v1")

	userHandler := handler.NewUserHandler(userService, authService)
	api.POST("/users", userHandler.RegisterUser) // handler->user.go
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailable)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	campaignHandler := handler.NewCampaignHandler(campaingService)
	api.GET("/campaigns", campaignHandler.GetCampaigns)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // karena middleware berada ditengah jadi menggunakan abort with status JSON. Hentikan (abort) tidak dilanjutkan
			return
		}

		// Bearer tokentokentoken
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // karena middleware berada ditengah jadi menggunakan abort with status JSON. Hentikan (abort) tidak dilanjutkan
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // karena middleware berada ditengah jadi menggunakan abort with status JSON. Hentikan (abort) tidak dilanjutkan
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response) // karena middleware berada ditengah jadi menggunakan abort with status JSON. Hentikan (abort) tidak dilanjutkan
			return
		}

		c.Set("currentUser", user)

	}
}

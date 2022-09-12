package main

import (
	itemImpl "dealljobs/domain/user/impl"
	"dealljobs/handler"
	"dealljobs/pkg"
	"dealljobs/util"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	pkg.InitDatabase()

	// init service & repo
	itemRepo := itemImpl.NewRepo(pkg.Database)
	itemSvc := itemImpl.NewService(itemRepo)

	// init handler
	apiHandler := handler.NewApiHandler(itemSvc)

	v1 := r.Group("/api/v1", Logger())
	{
		v1.POST("/login", apiHandler.Login)

		authorized := v1.Group("", util.AuthenticateMiddleware())

		authorized.GET("/check-auth", apiHandler.CheckAuth)
		authorized.GET("/user/:id", apiHandler.GetUser)

		admin := authorized.Use(util.AuthenticateAdmin())

		admin.GET("/users", apiHandler.GetAllUsers)
		admin.POST("/user", apiHandler.CreateUser)
		admin.PUT("/user", apiHandler.UpdateUser)
		admin.DELETE("/user", apiHandler.DeleteUser)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

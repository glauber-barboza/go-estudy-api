package routes

import (
	"estudar.com.br/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(route *gin.Engine) *gin.Engine {
	route.Group("/api/v1")
	{
		books := route.Group("/book")
		{
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/:id", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return route
}

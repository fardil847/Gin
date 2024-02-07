package routers

import (
	"Gin/controllers"

	_ "Gin/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Car API
// @version 1.0
// @description This is a sample service for managing cars
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("cars/:id", controllers.GetOneCar)
	router.POST("/cars", controllers.CreateCars)
	router.GET("/cars", controllers.GetAllCars)
	router.PATCH("/cars/:id", controllers.UpdateCar)
	router.DELETE("/cars/:id", controllers.DeleteCars)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router

}

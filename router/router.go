package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Postgres driver
	"github.com/louislaugier/go-bitmex/login"
)

// Start the router
func Start() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	ExposeHeaders:    []string{"Content-Type", "Date"},
	// 	AllowCredentials: true,
	// 	// AllowOriginFunc: func(origin string) bool {
	// 	// 	return origin == "http://localhost:3000"
	// 	// },
	// 	AllowAllOrigins: true,
	// }))

	r.GET("/api/v1/login", login.GET())

	return r
}

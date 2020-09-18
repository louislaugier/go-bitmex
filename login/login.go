package login

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qct/bitmex-go/swagger"
)

// GET login response
func GET() func(c *gin.Context) {
	return func(c *gin.Context) {

		var (
			APIkey    = c.Request.URL.Query()["key"][0]
			secretKey = c.Request.URL.Query()["secret"][0]
			Auth      = context.WithValue(context.TODO(), swagger.ContextAPIKey, swagger.APIKey{
				Key:    APIkey,
				Secret: secretKey,
			})
		)
		userAPI := swagger.NewAPIClient(swagger.NewConfiguration()).UserApi
		_, response, err := userAPI.UserGetWallet(Auth, map[string]interface{}{
			"currency": "",
		})
		if err != nil {
			log.Println("error: ", err)
		}
		log.Println(response.Status)
		if response.Status == "200 OK" {
			c.JSON(200, &gin.H{
				"statusCode": "200",
				"message":    "OK",
				"error":      nil,
				"meta": gin.H{
					"query": c.Request.URL.Query(),
				},
				"data": "OK",
			})
		} else if response.Status == "401 Unauthorized" {
			c.JSON(401, &gin.H{
				"statusCode": "401",
				"message":    "Wrong credentials.",
				"error":      nil,
				"meta": gin.H{
					"query": c.Request.URL.Query(),
				},
				"data": nil,
			})
		} else {
			c.JSON(401, &gin.H{
				"statusCode": "500",
				"message":    "Unexpected server error. Please try again.",
				"error":      response.Status,
				"meta": gin.H{
					"query": c.Request.URL.Query(),
				},
				"data": nil,
			})
		}
	}
}

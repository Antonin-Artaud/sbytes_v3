package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sbytes_v3/controllers"
	"sbytes_v3/services"
)

const uri = "mongodb://localhost:27017"

func main() {

	ginServer := gin.Default()

	err := ginServer.SetTrustedProxies([]string{"192.168.1.15"})
	if err != nil {
		log.Println(err)
	}

	mongoDb := services.NewMongoService(uri)

	ticketGroup := ginServer.Group("/api/v1/tickets")
	{
		controller := controllers.NewTicketController(mongoDb)
		ticketGroup.POST("/", controller.CreateTicket)
		ticketGroup.GET("/:id", controller.GetTicket)
		ticketGroup.PUT("/:id", controller.UpdateTicket)
	}

	if err := ginServer.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sbytes_v3/controllers"
	"sbytes_v3/services"
)

func main() {

	properties := loadPropertiesFromYamlFile()

	ginServer := gin.Default()

	if err := ginServer.SetTrustedProxies(properties.Server.TrustedProxies); err != nil {
		log.Println(err)
	}

	mongoDb := services.NewMongoService(properties.Database)

	ticketGroup := ginServer.Group(properties.Server.ContextTicket)
	{
		controller := controllers.NewTicketController(mongoDb)
		ticketGroup.POST("/", controller.CreateTicket)
		ticketGroup.GET("/:id", controller.GetTicket)
		ticketGroup.PUT("/:id", controller.UpdateTicket)
	}

	if err := ginServer.Run(properties.Server.Port); err != nil {
		log.Fatal(err)
	}
}

func loadPropertiesFromYamlFile() Properties {
	var properties Properties
	if err := cleanenv.ReadConfig("properties.yaml", &properties); err != nil {
		return Properties{}
	}
	return properties
}

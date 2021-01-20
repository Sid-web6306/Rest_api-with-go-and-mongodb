package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/Sid-web6306/restapi/config"
	"github.com/Sid-web6306/restapi/routes"
)


func main(){
	config.Connect()
	//SETTING UP ROUTES
	router:=gin.Default()
	
	// Simple group: v2
	routes.Routes(router)
	PORT:="400"
	
	// SERVER LISTENING ON PORT
	log.Fatal(router.Run(":"+PORT))

}
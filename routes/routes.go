package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Sid-web6306/restapi/controllers"
)
//CREATING DIFFERENT ROUTES INFORMATION
func Routes(router *gin.Engine){
	
	v1:=router.Group("/api/v1")
	{
		v1.GET("/",welcome)
		v1.GET("/persons",controllers.GetAllPersons)
		v1.POST("/persons/create",controllers.CreatePerson)
		v1.GET("/persons/searchName",controllers.FindPerson)
	}
	router.NoRoute(errorMessage)
	
		
	
	

}

// WELCOME TO PERSON API
func welcome(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"status": 200,
		"Message": "Welcome TO Persons API",
	})
	return
}


// ROUTE NOT FOUND FUNCTION
func errorMessage(c *gin.Context){
	c.JSON(http.StatusNotFound,gin.H{
		"status": 404,
		"Message": "Route Not Found",
	})
	return
}


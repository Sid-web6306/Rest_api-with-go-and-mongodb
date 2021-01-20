package controllers

import(
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	
	"github.com/Sid-web6306/restapi/models"
	
)



var collection *mongo.Collection

func PersonCollection(c *mongo.Database){
	// PERSONS TABLE CREATE
	collection = c.Collection("Persons")
}

// FECTHING ALL DATA INSERTED INTO DATABASE

func GetAllPersons(c *gin.Context){

	persons:=[]Person.Person{}
	cursor,err:= collection.Find(context.TODO(),bson.M{})
	if err!=nil{
		log.Printf("Error While getting all Perons: %v\n",err)
		c.JSON(http.StatusInternalServerError,gin.H{
			"status": http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var person Person.Person
		cursor.Decode(&person)
		persons = append(persons, person)
	}

	c.JSON(http.StatusOK,gin.H{
		"status":http.StatusOK,
		"Message":"All Persons Information",
		"data": persons,
	})
	return

}

// INSERTING A PERSON INFORMATION
func CreatePerson(c *gin.Context){
	
	var person Person.Person
	c.BindJSON(&person)
	newPerson:=Person.Person{
		ID:primitive.NewObjectID(),
		Name:person.Name,
		Address:person.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_,err:= collection.InsertOne(context.TODO(),newPerson)
	if err!= nil{
		log.Printf("Error While Creating Mulitple Persons info, %v\n",err)
		c.JSON(http.StatusInternalServerError,gin.H{
			"Status":http.StatusInternalServerError,
			"message":"Something went wrong while creating",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"Status":http.StatusOK,
		"Message":"Persons Created successfully",
		"ID":newPerson.ID,
	})
	return
}
// FINDING PERSON INFORMATION BY NAME

func FindPerson(c *gin.Context){
	personName:=c.Query("name")

	cursor,err1:= collection.Find(context.TODO(),bson.M{"name": personName})
	var personFiltered []bson.M
	if err := cursor.All(context.TODO(),&personFiltered); err != nil || err1!=nil {
		log.Printf("Error while searching a Person Info, %v\n",err)
		c.JSON(http.StatusInternalServerError,gin.H{
			"status":http.StatusInternalServerError,
			"Message":"Person Info Not Found",

		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"status":http.StatusOK,
		"Message":"Person Found",
		"data":personFiltered,
	})
	return

}


package todo

import (
	"cicd/internal/database"
	"cicd/internal/database/model"
	"cicd/templates"
	"cicd/templates/components"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ListTODO(c *gin.Context) {
	var todoList []model.TODOModel

	// Get collection
	collection := database.GetCollection(database.TODO)
	// Get datas
	cursor, _ := collection.Find(context.TODO(), gin.H{})
	if err := cursor.All(context.TODO(), &todoList); err != nil {
		c.HTML(http.StatusOK, "", templates.Error(err))
		return
	}
	c.HTML(http.StatusOK, "", templates.ListTodo(todoList))
}

func GetTODO(c *gin.Context) {}

func CreateTODO(c *gin.Context) {
	var task model.TODOModel
	todoCollection := database.GetCollection(database.TODO)

	if err := faker.FakeData(&task); err != nil {
		log.Printf("Fail to generate data: %v", err.Error())
		return
	}
	task.IsComplete = false
	if _, err := todoCollection.InsertOne(context.TODO(), task); err != nil {
		log.Printf("Fail to insert data: %v", err.Error())
	}
	c.HTML(http.StatusOK, "", components.Card(task))
}

func UpdateTOOD(c *gin.Context) {}

func DeleteTODO(c *gin.Context) {
	var tid string

	// Get collection
	collection := database.GetCollection(database.TODO)

	if tid = c.Param("tid"); tid == "" {
		c.HTML(http.StatusBadRequest, "", templates.Error(errors.New("Task ID not found")))
		return
	}

	var filter bson.M
	fmt.Println(tid)
	if objId, err := primitive.ObjectIDFromHex(tid); err != nil {
		c.HTML(http.StatusBadRequest, "", templates.Error(err))
		return
	} else {
		filter = bson.M{"_id": objId}
	}

	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {

		c.HTML(http.StatusBadRequest, "", templates.Error(err))
		return
	}
	fmt.Println("Here")
	ListTODO(c)
}

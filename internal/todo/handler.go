package todo

import (
	"cicd/internal/database"
	"cicd/internal/database/model"
	"cicd/templates"
	"cicd/templates/layout"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTODO(c *gin.Context) {
	var todoList []model.TODOModel

	// Get collection
	collection := database.GetCollection(database.TODO)
	// Get datas
	cursor, _ := collection.Find(context.TODO(), gin.H{})
	if err := cursor.All(context.TODO(), &todoList); err != nil {
		c.HTML(http.StatusOK, "", templates.Error(err))
	}
	c.HTML(http.StatusOK, "", layout.Base("Hello"))
}

func GetTODO(c *gin.Context) {}

func CreateTODO(c *gin.Context) {}

func UpdateTOOD(c *gin.Context) {}

func DeleteTODO(c *gin.Context) {}

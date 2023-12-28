package model

import (
	"cicd/internal/database/model/common"
)

type TODOModel struct {
	common.ObjectId `bson:"inline"`
	Title           string `json:"title" bson:"title" faker:"oneof: Grocery Shopping, Work Presentation, Fitness Routine, Book Club Reading, Home Cleaning, Career Development, Financial Planning, Quality Time with Family, Gardening, Project Deadline"`
	Description     string `json:"description" bson:"description" faker:"sentence"`
	IsComplete	    bool `json:"isComplete" bson:"isComplete"`
	common.Datetime `bson:"inline"`
}

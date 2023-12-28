package model

import (
	"cicd/internal/database/model/common"
)

type TODOModel struct {
	common.ObjectId `bson:"inline"`
	Title           string `json:"title" bson:"title"`
	Description     string `json:"description" bson:"description"`
	common.Datetime `bson:"inline"`
}

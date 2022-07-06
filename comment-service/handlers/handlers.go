package handlers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoHandler struct {
	CommentCollection *mongo.Collection
	ReportCollection  *mongo.Collection
}

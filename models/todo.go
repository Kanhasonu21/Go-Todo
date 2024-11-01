package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TODO struct {
	Task      string `bson:"task"`
	Completed bool   `bson:"completed"`
}

var collection *mongo.Collection

func SetCollection(c *mongo.Collection) {
	collection = c
}

func CreateTODO(todo TODO) (*TODO, error) {
	_, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetTODO() ([]TODO, error) {
	var todos []TODO
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var todo TODO
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil

}

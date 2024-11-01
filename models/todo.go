package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TODO struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Task      string             `bson:"task"`
	Completed bool               `bson:"completed"`
	CreatedAt time.Time          `bson:"created_at"`
	updatedAt time.Time          `bson:"updated_at"`
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

func UpdateTask(id string, update TODO) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err // Handle invalid ObjectID format
	}
	filter := bson.M{"_id": objID}
	updateData := bson.M{"$set": bson.M{"completed": update.Completed}}
	result, err := collection.UpdateOne(context.Background(), filter, updateData)
	fmt.Println(result)
	//panic(result)
	if err != nil {
		println(err)
		return false, err
	}
	return true, nil

}

func DeleteTask(id string) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objID}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount > 0, nil
}

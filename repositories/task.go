package repositories

import (
	"candidate-backend/db"
	"candidate-backend/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateTaskRepo(taskModel models.TaskModel) (response string, err error) {

	mongoDB, err := db.GetMongoDB()
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongoDB.Collection("tasks")

	result, err := collection.InsertOne(ctx, taskModel)
	if err != nil {
		return
	}

	response = result.InsertedID.(primitive.ObjectID).Hex()
	return
}

func GetTaskListRepo(page int64, perPage int64, keyword string) (response []models.TaskModel, err error) {

	mongoDB, err := db.GetMongoDB()
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongoDB.Collection("tasks")

	filter := bson.M{}
	filter["keepStatus"] = false

	if keyword != "" {
		filter["keyword"] = ""
	}

	//pagination
	limit := perPage
	skip := (page - 1) * limit

	findOptions := options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  bson.M{"createdDate": -1},
	}

	cursor, err := collection.Find(ctx, filter, &findOptions)
	if err != nil {
		return
	}

	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &response); err != nil {
		return
	}

	return
}

func GetTaskDetailRepo(taskId string) (response models.TaskModel, err error) {

	mongoDB, err := db.GetMongoDB()
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		err = errors.New("id format is wrong")
		return
	}

	collection := mongoDB.Collection("tasks")

	filter := bson.M{}
	filter["_id"] = objID

	if err = collection.FindOne(ctx, filter).Decode(&response); err != nil {
		return
	}

	return
}

func AddCommentRepo(taskId string, comment models.Comment) (response models.TaskModel, err error) {

	mongoDB, err := db.GetMongoDB()
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		err = errors.New("id format is wrong")
		return
	}

	collection := mongoDB.Collection("tasks")

	filter := bson.M{}
	filter["_id"] = objID

	update := bson.M{}
	update["$push"] = bson.M{
		"comments": comment,
	}

	result := collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return
	}

	err = result.Decode(&response)
	if err != nil {
		return
	}

	return
}

func KeepTaskRepo(taskId string) (response models.TaskModel, err error) {

	mongoDB, err := db.GetMongoDB()
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		err = errors.New("id format is wrong")
		return
	}

	collection := mongoDB.Collection("tasks")

	filter := bson.M{}
	filter["_id"] = objID

	update := bson.M{}
	update["$set"] = bson.M{
		"keepStatus": true,
	}

	result := collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return
	}

	err = result.Decode(&response)
	if err != nil {
		return
	}

	return
}

func SubmitStatusRepo(taskId string, status string) (response models.TaskModel, err error) {

	mongoDB, err := db.GetMongoDB()
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		err = errors.New("id format is wrong")
		return
	}

	collection := mongoDB.Collection("tasks")

	filter := bson.M{}
	filter["_id"] = objID

	update := bson.M{}
	update["$set"] = bson.M{
		"status": status,
	}

	result := collection.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return
	}

	err = result.Decode(&response)
	if err != nil {
		return
	}

	return
}
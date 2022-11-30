package repository

import (
	"context"
	"crowstream_reproduction_ms/app/domain"
	"crowstream_reproduction_ms/database"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserMetadata interface {
	Create(*domain.UserVideoMetadata) (primitive.ObjectID, error)
	GetById(*string, *int) (domain.UserVideoMetadata, error)
	GetAll() ([]domain.UserVideoMetadata, error)
}

type UserMetadataMongo struct{}

var collection = database.OpenCollection(database.Client, "UserVideoMetadata")

func (u UserMetadataMongo) Create(uvm *domain.UserVideoMetadata) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	var current domain.UserVideoMetadata

	err := collection.FindOne(ctx, bson.M{"$and": []bson.M{{"user_id": uvm.UserId}, {"video_id": uvm.VideoId}}}).Decode(&current)
	if err == nil {

		_, err := collection.UpdateOne(ctx, bson.M{"_id": current.ID}, bson.D{
			{"$set", bson.D{{"video_progress", uvm.VideoProgress}, {"video_progress_time", uvm.VideoProgressTime}}},
		})
		if err != nil {
			log.Printf("Could not update userdata: %v", err)
			return primitive.NilObjectID, err
		}

		return current.ID, err
	} else {
		uvm.ID = primitive.NewObjectID()

		result, err := collection.InsertOne(ctx, uvm)
		if err != nil {
			log.Printf("Could not create userdata: %v", err)
			return primitive.NilObjectID, err
		}

		oid := result.InsertedID.(primitive.ObjectID)

		return oid, err
	}

}

func (u UserMetadataMongo) GetAll() ([]domain.UserVideoMetadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []domain.UserVideoMetadata{}, err
	}

	var results []domain.UserVideoMetadata

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var click domain.UserVideoMetadata
		if err = cursor.Decode(&click); err != nil {
			log.Fatal(err)
		}
		results = append(results, click)
	}

	return results, err
}

func (u UserMetadataMongo) GetById(userId *string, videoId *int) (domain.UserVideoMetadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	var result domain.UserVideoMetadata

	err := collection.FindOne(ctx, bson.M{"$and": []bson.M{{"user_id": userId}, {"video_id": videoId}}}).Decode(&result)
	if err != nil {
		return domain.UserVideoMetadata{}, err
	}

	return result, err
}

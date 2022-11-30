package repository

import (
	"context"
	"crowstream_reproduction_ms/app/domain"
	"crowstream_reproduction_ms/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type ClickMetadata interface {
	Create(metadata *domain.ClickCountMetadata) (primitive.ObjectID, error)
	GetById(*string, *int) (domain.ClickCountMetadata, error)
	UpdateClickVideo(metadata *domain.ClickCountMetadata) (domain.ClickCountMetadata, error)
	GetAll()([]domain.ClickCountMetadata, error)
}

type ClickMetadataMongo struct{}

var clickCollection = database.OpenCollection(database.Client, "ClickCountMetadata")

func (u ClickMetadataMongo) Create(uvm *domain.ClickCountMetadata) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	uvm.ID = primitive.NewObjectID()

	result, err := clickCollection.InsertOne(ctx, uvm)
	if err != nil {
		log.Printf("Could not create click userdata: %v", err)
		return primitive.NilObjectID, err
	}

	oid := result.InsertedID.(primitive.ObjectID)

	return oid, err

}

func (u ClickMetadataMongo) GetAll() ([]domain.ClickCountMetadata, error){
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	cursor, err := clickCollection.Find(ctx,bson.M{})
	if err != nil {
		return []domain.ClickCountMetadata{}, err
	}

	var results []domain.ClickCountMetadata

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var click domain.ClickCountMetadata
		if err = cursor.Decode(&click); err != nil {
			log.Fatal(err)
		}
		results = append(results, click)
	}

	return results,err
}

func (u ClickMetadataMongo) GetById(userId *string,videoId *int) (domain.ClickCountMetadata, error){
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	var result domain.ClickCountMetadata

	err := clickCollection.FindOne(ctx, bson.M{"$and": []bson.M{{"user_id":userId}, {"video_id":videoId} }}).Decode(&result)
	if err != nil {
		return domain.ClickCountMetadata{}, err
	}

	return result,err
}

func (u ClickMetadataMongo) UpdateClickVideo(uvm *domain.ClickCountMetadata) (domain.ClickCountMetadata, error){
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	var result domain.ClickCountMetadata

	var err error

	result, err = u.GetById(&uvm.UserId, &uvm.VideoId)
	if err != nil {
		return domain.ClickCountMetadata{}, err
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cmu := clickCollection.FindOneAndUpdate(ctx, bson.M{"_id": result.ID}, bson.D{{"$set", bson.D{{"click_video", uvm.ClickVideo}}}}, &opt)

	var doc domain.ClickCountMetadata
	err = cmu.Decode(&doc)

	return doc,err
}

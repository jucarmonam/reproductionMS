package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserVideoMetadata struct {
	ID                primitive.ObjectID `bson:"_id"`
	UserId            string             `json:"user_id" bson:"user_id"`
	VideoId           int                `json:"video_id" bson:"video_id"`
	VideoProgress     float32            `json:"video_progress" bson:"video_progress"`
	VideoProgressTime string             `json:"video_progress_time" bson:"video_progress_time"`
}

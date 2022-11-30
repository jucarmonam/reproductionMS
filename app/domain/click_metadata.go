package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClickCountMetadata struct {
	ID                primitive.ObjectID `bson:"_id"`
	UserId            string             `json:"user_id" bson:"user_id"`
	VideoId           int                `json:"video_id" bson:"video_id"`
	ClickDescription  bool            	 `json:"click_description" bson:"click_description"`
	ClickVideo 		  bool             	 `json:"click_video" bson:"click_video"`
}

package app

import (
	"crowstream_reproduction_ms/app/controller"
	"crowstream_reproduction_ms/app/repository"

	"github.com/gin-gonic/gin"
)

var umc controller.UserMetadata
var cmc controller.ClickMetadata

// TODO: Replace this with dependency injection.
func InitDepenedencies() {
	//video dependencies
	umr := repository.UserMetadataMongo{}
	umc = controller.UserMetadata{Umr: umr}

	//click dependencies
	cmr := repository.ClickMetadataMongo{}
	cmc = controller.ClickMetadata{Cmr: cmr}
}

func InitRoutes(router *gin.Engine) error {
	//video routes
	router.GET("/user-video-metadata", umc.FindAll)
	router.GET("/user-video-metadata/:userId/:videoId", umc.FindVideoMetadata)
	router.POST("/user-video-metadata", umc.PostUserVideoMetadata)
	//click routes
	router.GET("/click-count-metadata", cmc.FindAll)
	router.GET("/click-count-metadata/:userId/:videoId", cmc.FindClickMetadata)
	router.POST("/click-count-metadata", cmc.PostClickCountMetadata)
	router.PUT("/click-count-metadata", cmc.PutClickCountMetadata)

	return nil
}

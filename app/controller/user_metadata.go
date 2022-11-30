package controller

import (
	"crowstream_reproduction_ms/app/domain"
	"crowstream_reproduction_ms/app/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: Replace with proper dependency injection.
type UserMetadata struct {
	Umr repository.UserMetadata
}

func (uc UserMetadata) FindAll(c *gin.Context) {
	ccm, err := uc.Umr.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ccm)
}

func (uc UserMetadata) FindVideoMetadata(c *gin.Context) {

	userId := c.Param("userId")
	videoId, _ := strconv.Atoi(c.Param("videoId"))

	uvm, err := uc.Umr.GetById(&userId, &videoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, uvm)
}

func (uc UserMetadata) PostUserVideoMetadata(c *gin.Context) {
	var json domain.UserVideoMetadata

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Validate input data (maybe on a valid format, but not with the rigth info).
	/*

		type domain.UserVideoMetadataInput struct {
			// ...
		}

		validator.ValidVideoId(input.VideoId)
		validator.ValidUserId(input.UserId)

	*/

	log.Printf("%v\n", json)
	oid, err := uc.Umr.Create(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"ID": oid})
}

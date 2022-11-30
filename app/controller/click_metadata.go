package controller

import (
	"crowstream_reproduction_ms/app/domain"
	"crowstream_reproduction_ms/app/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type ClickMetadata struct {
	Cmr repository.ClickMetadata
}

func (uc ClickMetadata) FindAll(c *gin.Context){
	ccm, err := uc.Cmr.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ccm)
}

func (uc ClickMetadata) FindClickMetadata(c *gin.Context){

	userId := c.Param("userId")
	videoId, _ := strconv.Atoi(c.Param("videoId"))

	ccm, err := uc.Cmr.GetById(&userId, &videoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ccm)
}

func (uc ClickMetadata) PostClickCountMetadata(c *gin.Context) {
	var json domain.ClickCountMetadata

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("%v\n", json)
	oid, err := uc.Cmr.Create(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"ID": oid})
}

func (uc ClickMetadata) PutClickCountMetadata(c *gin.Context) {
	var json domain.ClickCountMetadata

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("%v\n", json)
	ccm, err := uc.Cmr.UpdateClickVideo(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, ccm)
}

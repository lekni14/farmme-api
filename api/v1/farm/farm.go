package farm

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
	// "os"

	// guuid "github.com/google/uuid"

	// handle_user "farmme-api/api/v1/user"
	"farmme-api/repository"
	"farmme-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"farmme-api/model"
	"farmme-api/pkg/app"
	"farmme-api/pkg/e"
	"github.com/gin-gonic/gin"
)

// FarmAPI is a representation of a FarmAPI
type FarmAPI struct {
	FarmRepository repository.FarmRepository
}

func (api FarmAPI) AddFarm(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	userID, _, _ := utils.GetTokenValue(c)
	ownerObjectID, _ := primitive.ObjectIDFromHex(userID)

	var json model.Farm

	json.OwnerID = ownerObjectID


	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err2 := api.FarmRepository.ExistByName(json.Name)

	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err2.Error()})
		return
	}

	if exists {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_FARM, nil)
		return
	}
	json.Address = []model.Farmaddress{}
	farmID, err := api.FarmRepository.AddFarm(json)
	if err != nil {
		log.Println("error AddFarm", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, farmID)
	//c.JSON(http.StatusCreated, gin.H{"status": "susess"})
	//c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

// func (api EventAPI) MyEvent(c *gin.Context) {
func (api FarmAPI) MyFarm(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	userID, _, _ := utils.GetTokenValue(c)
	log.Println("Print userID", userID)
	farm, err := api.FarmRepository.GetFarmByUser(userID)
	if err != nil {
		log.Println("error AddFarm", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, farm)

}
package cow

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

// CowAPI is a representation of a CowAPI
type CowAPI struct {
	CowRepository repository.CowRepository
}

func (api CowAPI) AddCow(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	userID, _, _ := utils.GetTokenValue(c)
	ownerObjectID, _ := primitive.ObjectIDFromHex(userID)

	var json model.Cow

	json.OwnerID = ownerObjectID


	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cowID, err := api.CowRepository.AddCow(json)
	if err != nil {
		log.Println("error AddCow", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, cowID)
}
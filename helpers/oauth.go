package helpers

import (
	"TTT/mod/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleCallback(c *gin.Context) {
	var gtrm models.GoogleTokenResponseModel
	err := c.BindJSON(&gtrm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error getting body")
		return
	}
	fmt.Printf("Body %s", gtrm.UserData.ID)
	// Save token to DB.
	// Validate token
	c.JSON(http.StatusOK, gtrm)
}

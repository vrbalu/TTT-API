package controllers

import (
	"TTT/mod/models"
	"TTT/mod/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FriendshipsController struct{}

func (*FriendshipsController) CreateFriendship(c *gin.Context) {
	var newFriendship models.FriendshipCreate
	err := c.BindJSON(&newFriendship)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error parsing body")
		return
	}
	err = services.FriendshipService.CreateFriendship(&newFriendship)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error creating friendship"+err.Error())
		return
	}
	c.JSON(http.StatusCreated, "")
}

func (*FriendshipsController) DeleteFriendship(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error deleting friendship"+err.Error())
		return
	}
	err = services.FriendshipService.DeleteFriendship(intId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error deleting friendship"+err.Error())
		return
	}
	c.JSON(http.StatusNoContent, "")
}

func (*FriendshipsController) GetFriendshipById(c *gin.Context) {
	id := c.Param("id")
	res, err := services.FriendshipService.GetFriendshipById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error getting friendship"+err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (*FriendshipsController) GetFriendships(c *gin.Context) {
	user := c.Query("user")
	isPending := c.Query("isPending")
	forRequest := c.Query("forRequest")
	res, err := services.FriendshipService.GetFriendships(user, isPending, forRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error getting friendship"+err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (*FriendshipsController) UpdateFriendship(c *gin.Context) {
	id := c.Param("id")
	isPending := c.Query("isPending")
	isPendingBool, err := strconv.ParseBool(isPending)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Error parsing pending status"+err.Error())
		return
	}
	err = services.FriendshipService.UpdateFriendshipPendingStatus(id, isPendingBool)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Error updating friendship"+err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

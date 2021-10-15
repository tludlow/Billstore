package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"test1/user_api/repo"
)

func getUserDetails(c *gin.Context){
	userId, _ := uuid.FromString(c.GetString("userId"))
	details := repo.GetUserDetailsById(userId)
	c.JSON(http.StatusOK, gin.H{"user": details})
}
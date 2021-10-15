package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"test1/auth_api/repo"
	"test1/auth_api/commands"
	"test1/auth_api/hashing"
	"test1/auth_api/requests"
	"test1/core/jwt"
	"time"
)

func register(c *gin.Context) {
	var request requests.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := uuid.NewV4()
	userIdentityId, _ := uuid.NewV4()
	refreshTokenId, _ := uuid.NewV4()
	go commands.CreateUser(userId, request.Username, request.Email)
	go commands.CreateEmailUserIdentity(userIdentityId, request.Email, request.Password, userId)
	go commands.CreateRefreshToken(refreshTokenId, userIdentityId, userId)
	setRefreshTokenCookie(refreshTokenId.String(), c)
	setAccessTokenCookie(userId.String(), c)
	c.JSON(http.StatusOK, gin.H{"username": request.Username})
}

func login(c *gin.Context){
	var request requests.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	identity := repo.GetEmailUserIdentityByEmail(request.Email)
	validPassword, _ := hashing.ComparePasswordAndHash(request.Password, identity.Password)
	if !validPassword{
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password incorrect"})
		return
	}
	refreshTokenId, _ := uuid.NewV4()
	go commands.CreateRefreshToken(refreshTokenId, identity.Id, identity.UserId)
	setRefreshTokenCookie(refreshTokenId.String(), c)
	setAccessTokenCookie(identity.UserId.String(), c)
}

func logout(c *gin.Context){
	defer c.SetCookie(
		"access_token",
		"",
		-1,
		"/",
		"localhost:5001",
		false,
		true,
	)
	defer c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"localhost:5001",
		false,
		true,
	)
	refreshTokenCookie, err := c.Request.Cookie("refresh_token")
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token already revoked"})
		return
	}
	tokenId, err := jwt.Validate(refreshTokenCookie.Value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh_token already revoked"})
		return
	}
	go commands.DeleteRefreshToken(tokenId)
}

func accessToken(c *gin.Context){
	refreshTokenCookie, err := c.Request.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh_token"})
		return
	}
	tokenId, err := jwt.Validate(refreshTokenCookie.Value)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	refreshToken, err := repo.GetRefreshTokenById(tokenId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Revoked refresh_token"})
		return
	}
	setAccessTokenCookie(refreshToken.UserId.String(), c)
	c.JSON(http.StatusOK, gin.H{"id": tokenId.String()})
}

func clearRefreshTokens(c *gin.Context){
	userId, _ := uuid.FromString(c.GetString("userId"))
	count := commands.DeleteAllRefreshTokens(userId)
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func test(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"id": c.GetString("userId")})
}

func setRefreshTokenCookie(id string, c *gin.Context){
	c.SetCookie(
		"refresh_token",
		jwt.Generate(id, time.Now().Add(24 * time.Hour * 365).Unix()),
		3600 * 24 * 365,
		"/",
		"localhost:5001",
		false,
		true,
	)
}

func setAccessTokenCookie(userId string, c *gin.Context){
	c.SetCookie(
		"access_token",
		jwt.Generate(userId, time.Now().Add(20 * time.Hour).Unix()),
		60 * 20,
		"/",
		"localhost:5001",
		false,
		true,
	)
}
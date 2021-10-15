package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"test1/auth_api/consumers"
	"test1/auth_api/repo"
	"test1/auth_api/validators"
	"test1/core/middleware"
	"test1/core/mongo_repo"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	if os.Getenv("MONGODB_CONNECTION_STRING") == ""{
		mongo_repo.ConnectionString = "mongodb://root:password@127.0.0.1:27017/"
	}
	if os.Getenv("MONGODB_DB") == ""{
		mongo_repo.Database = "_auth"
	}

	var cancelCtx context.CancelFunc
	repo.Ctx, cancelCtx = context.WithCancel(context.Background())
	defer cancelCtx()
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://root:password@127.0.0.1:27017/").SetRegistry(mongo_repo.MongoRegistry))
	if err != nil {log.Fatalf("Error creating mongo client: %+v", err)}
	defer client.Disconnect(repo.Ctx) // add disconnect call here
	if err := client.Connect(repo.Ctx); err != nil {log.Fatalf("Failed to connect to MongoDB: %+v", err)}
	repo.Repo = &mongo_repo.ConnectionWrapper{Client: client, Ctx: repo.Ctx, CancelCtx: cancelCtx}
	repo.UserCollection = client.Database(mongo_repo.Database).Collection("emailUserIdentities")
	repo.ChannelCollection = client.Database(mongo_repo.Database).Collection("refreshTokens")

	r := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("isEmail", validators.IsEmail)
		_ = v.RegisterValidation("emailUnique", validators.EmailUnique)
	}
	auth := r.Group("", middleware.Auth)
	auth.GET("/api/test", test)
	auth.POST("/api/clear_refresh_tokens", clearRefreshTokens)
	r.POST("/api/register", register)
	r.POST("/api/login", login)
	r.POST("/api/logout", logout)
	r.GET("/api/access_token", accessToken)
	go consumers.Start()
	err = r.Run(":5001")
	if err != nil {return}
}
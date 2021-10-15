package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"test1/core/middleware"
	"test1/core/mongo_repo"
	"test1/user_api/consumers"
	"test1/user_api/repo"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	if os.Getenv("MONGODB_CONNECTION_STRING") == ""{
		mongo_repo.ConnectionString = "mongodb://root:password@127.0.0.1:27017/"
	}
	if os.Getenv("MONGODB_DB") == ""{
		mongo_repo.Database = "_user"
	}

	ctx, cfunc := context.WithCancel(context.Background())
	defer cfunc()
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://root:password@127.0.0.1:27017/").SetRegistry(mongo_repo.MongoRegistry))
	if err != nil {log.Fatalf("Error creating mongo client: %+v", err)}
	defer client.Disconnect(ctx) // add disconnect call here
	if err := client.Connect(ctx); err != nil {log.Fatalf("Failed to connect to MongoDB: %+v", err)}
	repo.Repo = &mongo_repo.ConnectionWrapper{Client: client, Ctx: ctx, CancelCtx: cfunc}

	r := gin.Default()
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	_ = v.RegisterValidation("isEmail", validators.IsEmail)
	//	_ = v.RegisterValidation("emailUnique", validators.EmailUnique)
	//}

	auth := r.Group("", middleware.Auth)
	auth.GET("/api/user_details", getUserDetails)
	go consumers.Start()
	err = r.Run(":5005")
	if err != nil {return}
}
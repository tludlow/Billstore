package repo

import (
	"context"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"test1/core/domain"
	"test1/core/mongo_repo"
	"time"
)

var Repo *mongo_repo.ConnectionWrapper

var(
	Ctx context.Context
	UserCollection *mongo.Collection
	ChannelCollection *mongo.Collection
	ChannelSettingsCollection *mongo.Collection
)

func AddEmailUserIdentity(identity domain.EmailUserIdentity){
	doc := bson.M{
		"id": identity.Id,
		"email": identity.Email,
		"password": identity.Password,
		"userId": identity.UserId,
	}
	Repo.AddSingle("emailUserIdentities", doc)
}
func GetEmailUserIdentityByEmail(email string) domain.EmailUserIdentity {
	ctx, cancelCtx := context.WithTimeout(Repo.Ctx, time.Second * 3)
	defer cancelCtx()
	collection := Repo.Client.Database(mongo_repo.Database).Collection("emailUserIdentities")
	filter := bson.D{{"email", email}}
	res := new(domain.EmailUserIdentity)
	err := collection.FindOne(ctx, filter).Decode(res)
	if err != nil {log.Println(err)}
	return *res
}
func GetEmailUserIdentityById(id uuid.UUID) domain.EmailUserIdentity {
	ctx, cancelCtx := context.WithTimeout(Repo.Ctx, time.Second * 3)
	defer cancelCtx()
	collection := Repo.Client.Database(mongo_repo.Database).Collection("emailUserIdentities")
	filter := bson.D{{"id", id}}
	res := new(domain.EmailUserIdentity)
	err := collection.FindOne(ctx, filter).Decode(res)
	if err != nil {log.Println(err)}
	return *res
}
func EmailUserIdentityExists(email string) bool{
	ctx, cancelCtx := context.WithTimeout(Repo.Ctx, time.Second * 3)
	defer cancelCtx()
	collection := Repo.Client.Database(mongo_repo.Database).Collection("emailUserIdentities")
	filter := bson.D{{"email", email}}
	count, _ := collection.CountDocuments(ctx, filter)
	return count > 0
}
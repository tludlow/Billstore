package repo

import (
	"context"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"test1/core/mongo_repo"
	"test1/user_api/view_models"
	"time"
)

var(
	Repo *mongo_repo.ConnectionWrapper
)

func AddUserDetails(details view_models.UserDetails){
	doc := bson.M{
		"id":       details.Id,
		"username": details.Username,
		"contactEmail": details.ContactEmail,
		"ageVerified": details.AgeVerified,
		"matureContentFilter": details.MatureContentFilter,
	}
	Repo.AddSingle("userDetails", doc)
}

func GetUserDetailsById(id uuid.UUID) view_models.UserDetails{
	ctx, cancelCtx := context.WithTimeout(Repo.Ctx, time.Second * 3)
	defer cancelCtx()
	collection := Repo.Client.Database(mongo_repo.Database).Collection("userDetails")
	filter := bson.D{{"id", id}}
	res := new(view_models.UserDetails)
	err := collection.FindOne(ctx, filter).Decode(res)
	if err != nil {log.Println(err)}
	return *res
}
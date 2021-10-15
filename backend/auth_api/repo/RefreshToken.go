package repo

import (
	"context"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"test1/core/domain"
	"test1/core/mongo_repo"
	"time"
)

func GetRefreshTokenById(id uuid.UUID) (domain.RefreshToken, error){
	ctx, cancelCtx := context.WithTimeout(Repo.Ctx, time.Second * 3)
	defer cancelCtx()
	collection := Repo.Client.Database(mongo_repo.Database).Collection("refreshTokens")
	filter := bson.D{{"id", id}}
	res := new(domain.RefreshToken)
	err := collection.FindOne(ctx, filter).Decode(res)
	if err != nil {return domain.RefreshToken{}, err}
	return *res, nil
}

func AddRefreshToken(token domain.RefreshToken){
	doc := bson.M{
		"id": token.Id,
		"userId": token.UserId,
		"userIdentityId": token.UserIdentityId,
	}
	Repo.AddSingle("refreshTokens", doc)
}

func DeleteRefreshTokenById(id uuid.UUID){
	doc := bson.M{
		"id": id,
	}
	Repo.DeleteSingle("refreshTokens", doc)
}

func GetAllRefreshTokensForUser(userId uuid.UUID) []domain.RefreshToken{
	ctx, cancelCtx := context.WithTimeout(Repo.Ctx, time.Second * 3)
	defer cancelCtx()
	collection := Repo.Client.Database(mongo_repo.Database).Collection("refreshTokens")
	filter := bson.D{{"userId", userId}}
	cur, err := collection.Find(ctx, filter)
	if err != nil { log.Fatal(err) }
	defer cur.Close(ctx)
	var results []domain.RefreshToken
	if err = cur.All(ctx, &results); err != nil {log.Fatal(err)}
	if err != nil {log.Fatal(err)}
	return results
}

func DeleteAllRefreshTokensForUser(userId uuid.UUID) int64{
	ctx, cancelCtx := context.WithTimeout(Repo.Ctx, time.Second * 3)
	defer cancelCtx()
	collection := Repo.Client.Database(mongo_repo.Database).Collection("refreshTokens")
	filter := bson.D{{"userId", userId}}
	res, _ := collection.DeleteMany(ctx, filter)
	return res.DeletedCount
}
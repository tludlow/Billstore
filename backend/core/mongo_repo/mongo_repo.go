package mongo_repo

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"reflect"
	"time"
)

var (
	tUUID   = reflect.TypeOf(uuid.UUID{})
	uuidSubtype = byte(0x04)

	MongoRegistry = bson.NewRegistryBuilder().
		RegisterTypeEncoder(tUUID, bsoncodec.ValueEncoderFunc(uuidEncodeValue)).
		RegisterTypeDecoder(tUUID, bsoncodec.ValueDecoderFunc(uuidDecodeValue)).
		Build()
	ConnectionString = os.Getenv("MONGODB_CONNECTION_STRING")
	Database = os.Getenv("MONGODB_DB")
)

type Repo struct{
	Client *mongo.Client
}
//func NewRepo() Repo {
//	client, err := mongo.NewClient(options.Client().ApplyURI(ConnectionString).SetRegistry(mongoRegistry))
//	if err != nil {log.Println(err)}
//	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
//	repo := Repo{
//		Client:    client,
//	}
//	err = client.Connect(ctx)
//	if err != nil {log.Println(err)}
//	return repo
//}

func uuidEncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	if !val.IsValid() || val.Type() != tUUID {
		return bsoncodec.ValueEncoderError{Name: "uuidEncodeValue", Types: []reflect.Type{tUUID}, Received: val}
	}
	b := val.Interface().(uuid.UUID)
	return vw.WriteBinaryWithSubtype(b[:], uuidSubtype)
}

func uuidDecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Type() != tUUID {
		return bsoncodec.ValueDecoderError{Name: "uuidDecodeValue", Types: []reflect.Type{tUUID}, Received: val}
	}

	var data []byte
	var subtype byte
	var err error
	switch vrType := vr.Type(); vrType {
	case bsontype.Binary:
		data, subtype, err = vr.ReadBinary()
		if subtype != uuidSubtype {
			return fmt.Errorf("unsupported binary subtype %v for UUID", subtype)
		}
	case bsontype.Null:
		err = vr.ReadNull()
	case bsontype.Undefined:
		err = vr.ReadUndefined()
	default:
		return fmt.Errorf("cannot decode %v into a UUID", vrType)
	}

	if err != nil {
		return err
	}
	uuid2, err := uuid.FromBytes(data)
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(uuid2))
	return nil
}

type ConnectionWrapper struct{
	Client *mongo.Client
	Ctx    context.Context
	CancelCtx context.CancelFunc
}
func (c *ConnectionWrapper) AddSingle(tableName string, doc bson.M){
	ctx, cfunc := context.WithTimeout(c.Ctx, time.Second * 3)
	defer cfunc() // good form to add this
	collection := c.Client.Database(Database).Collection(tableName)
	res, err := collection.InsertOne(ctx, doc)
	if err != nil {log.Fatal(err)}
	fmt.Println(res)
}

func (c *ConnectionWrapper) DeleteSingle(tableName string, doc bson.M){
	ctx, cfunc := context.WithTimeout(c.Ctx, time.Second * 3)
	defer cfunc() // good form to add this
	collection := c.Client.Database(Database).Collection(tableName)
	collection.DeleteOne(ctx, doc)
}
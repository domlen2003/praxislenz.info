package handlers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/url"
	"os"
	"time"
)

var mongoClient *mongo.Client

type InfoNode struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Type      Infotypes          `bson:"type,omitempty"`
	Content   string             `bson:"content,omitempty"`
	Timestamp string             `bson:"timestamp,omitempty"`
}

func TestConnection() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://praxislenz:"+url.QueryEscape(os.Getenv("MONGO_PASSWORD"))+"@202.61.250.84:42069/?authSource=praxislenz"))
	if err != nil {
		log.Fatal("ClientConnect: ", err)
	}
	defer client.Disconnect(ctx)
}

func UpdateInfo(info InfoNode) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://praxislenz:"+url.QueryEscape(os.Getenv("MONGO_PASSWORD"))+"@202.61.250.84:42069/?authSource=praxislenz"))
	if err != nil {
		log.Fatal("ClientConnect: ", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("praxislenz")
	infos := db.Collection("website-infos")

	result, err := infos.InsertOne(ctx, info)
	if err != nil {
		log.Fatal("UpdateInfo: ", err)
	}
	fmt.Println(result.InsertedID)
}

func GetInfo() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://praxislenz:"+url.QueryEscape(os.Getenv("MONGO_PASSWORD"))+"@202.61.250.84:42069/?authSource=praxislenz"))
	if err != nil {
		log.Fatal("ClientConnect: ", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("praxislenz")
	infos := db.Collection("website-infos")

	var infoNodes []InfoNode
	infoCursor, err := infos.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = infoCursor.All(ctx, &infoNodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(infoNodes)
}

//Idiomatic enum implementation for infotypes
type infotype string

const (
	CoronaInfo  infotype = "cinfo"
	GeneralInfo infotype = "ginfo"
)

type Infotypes interface {
	Infotype() infotype
}

func (b infotype) Infotype() infotype {
	return b
}

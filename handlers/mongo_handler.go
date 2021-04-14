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
var cacheValid = make(map[Infotype]bool)
var cache = make(map[Infotype][]InfoNode)

type InfoNode struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Type      Infotype           `bson:"type,omitempty"`
	Content   string             `bson:"content,omitempty"`
	Timestamp string             `bson:"timestamp,omitempty"`
}
type Infotype string

const (
	CoronaInfo   Infotype = "Corona Info"
	GeneralInfo  Infotype = "Generelle Info"
	OpeningHours Infotype = "Öffnungszeiten"
)

func StartMongoHandler() {
	//Cache Invalidator invalidiert den cache alle 10 Minuten,
	//falls durch äußeren Eingriff der Cache nicht mehr gültig sein sollte
	go func() {
		for true {
			for key, _ := range cacheValid {
				cacheValid[key] = false
			}
			fmt.Println("Cache invalidated")
			time.Sleep(10 * time.Minute)
		}
	}()
}

func UpdateInfo(info InfoNode) {
	//Delete all occurrences of this specific type
	toDelete := GetInfo(info.Type)
	for _, element := range toDelete {
		deleteInfo(element.ID)
	}
	//Get the Website-Info Collection
	infos, ctx := getInfoCollection()
	//Insert the new Value
	_, err := infos.InsertOne(ctx, info)
	if err != nil {
		log.Fatal("UpdateInfo: ", err)
	}
	//Invalidate cache for the updated type
	delete(cache, info.Type)
	cacheValid[info.Type] = false
}

func deleteInfo(id primitive.ObjectID) {
	//Get the Website-Info Collection
	infos, ctx := getInfoCollection()
	//Delete the given ID
	_, err := infos.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
}

func GetInfo(itype Infotype) []InfoNode {
	//Check if the cache has been invalidated by an Update-Info
	if cacheValid[itype] {
		//return the cache for the given type if cache is valid
		return cache[itype]
	} else {
		//Get the Website-Info Collection
		infos, ctx := getInfoCollection()
		//Search for the Infotype and return all occurrences
		var infoNodes []InfoNode
		infoCursor, err := infos.Find(ctx, bson.M{"type": itype})
		if err != nil {
			log.Fatal(err)
		}
		if err = infoCursor.All(ctx, &infoNodes); err != nil {
			log.Fatal(err)
		}
		//In case of not instantiated values for this Infotype
		//a new Value of this Infotype gets stored
		//and the Info request gets repeated
		if len(infoNodes) < 1 {
			dummyInfo := InfoNode{
				ID:        primitive.ObjectID{},
				Type:      itype,
				Content:   " ",
				Timestamp: time.Now().Format("2.1.2006 15:04"),
			}
			_, err := infos.InsertOne(ctx, infos)
			if err != nil {
				log.Fatal("UpdateInfo: ", err)
			}
			return []InfoNode{dummyInfo}
		} else {
			//Write the Values to the Cache and revalidate the cache for the specific Infotype
			cacheValid[itype] = true
			cache[itype] = infoNodes
			return infoNodes
		}
	}
}

//returns a short lived connection to the Website-Infos Collection
func getInfoCollection() (*mongo.Collection, context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://praxislenz:"+url.QueryEscape(os.Getenv("MONGO_PASSWORD"))+"@202.61.250.84:42069/?authSource=praxislenz"))
	if err != nil {
		log.Fatal("ClientConnect: ", err)
	}

	db := client.Database("praxislenz")
	return db.Collection("website-infos"), ctx
}

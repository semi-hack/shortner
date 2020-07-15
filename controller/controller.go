package controller

import (
	"context"
	"fmt"
	"log"
	"shortner/db"
	"shortner/models"

	gonanoid "github.com/matoous/go-nanoid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Generateshort ...
func Generateshort(link *models.Link) (primitive.ObjectID, error) {
	id, err := gonanoid.Nanoid(5)
	link.Short = id
	result, err := db.GetConnection().Collection(models.Linkcollection).InsertOne(context.TODO(), link)
	if err != nil {
		log.Printf("could not generate id", err)
	}
	zut := result.InsertedID.(primitive.ObjectID)
	return zut, nil

}

// Getslug ...
func Getslug(slug string) (*models.Link, error) {
	var link *models.Link
	query := bson.M{"short": slug}
	err := db.GetConnection().Collection(models.Linkcollection).FindOne(context.TODO(), query).Decode(&link)
	if err != nil {
		log.Printf("failure", err)
		return nil, err
	}
	fmt.Println(link)
	log.Printf("Todos: %v")
	return link, nil

}

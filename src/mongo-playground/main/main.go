package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Prof struct {
	Name    string
	Age     int
	Country string
}

func main() {
	// Create client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Define Structures
	khan := Prof{"Adil Khan", 40, "Pakistan"}
	merkin := Prof{"Leonid Merkin", 48, "Russia"}
	succi := Prof{"Giancarlo Succi", 55, "Italy"}

	//Get collection
	collection := client.Database("test").Collection("profs")

	// Insert one document
	insertResult, err := collection.InsertOne(context.TODO(), succi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert many documents
	profs := []interface{}{khan, merkin}
	insertManyResult, err := collection.InsertMany(context.TODO(), profs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// Creating filter for update
	filter := bson.D{{"name", "Giancarlo Succi"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	// Updating Age
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Search of document
	var result Prof
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	// Using cursor
	fmt.Printf("\nUsing cursor\n")
	options := options.Find()
	options.SetLimit(100)
	filterCur := bson.M{}

	var results []*Prof

	cur, err := collection.Find(context.TODO(), filterCur, options)
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for cur.Next(context.TODO()) {
		var elem Prof
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
		fmt.Printf("Found document %d in cursor: %+v\n", i, elem)
		i++
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

	//// Deletion
	//filterDel := bson.M{}
	//deleteResult, err := collection.DeleteMany(context.TODO(), filterDel)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Deleted %v documents in the profs collection\n", deleteResult.DeletedCount)

	// Disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

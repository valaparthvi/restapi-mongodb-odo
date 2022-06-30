package repository

import (
	"context"
	"fmt"
	"os"

	"go-rest-mongodb/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlacesRepository struct{}

var collection = new(mongo.Collection)

const PlacesCollection = "Places"

func init() {
	// Connect to DB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	username := getBinding("username")
	password := getBinding("password")
	host := getBinding("host")
	defaultPort := "27017"
	port := getBindingWithDefault("port", &defaultPort)
	defaultDatabase := "go-rest-mongodb"
	database := getBindingWithDefault("database", &defaultDatabase)

	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%s/admin?ssl=false", username, password, host, port)
	fmt.Printf("DEBUG: MongoDB connection string: %s\n", mongoUri)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}

	r := client.Database("admin").RunCommand(
		context.Background(),
		bson.D{{"grantRolesToUser", username},
			{"roles", bson.A{bson.D{{"db", database}, {"role", "readWrite"}}}}},
	)
	if r.Err() != nil {
		panic(r.Err())
	}

	collection = client.Database(database).Collection(PlacesCollection)
}

func getBinding(name string) string {
	return getBindingWithDefault(name, nil)
}

func getBindingWithDefault(name string, defaultValue *string) string {
	u := os.Getenv(name)
	if u == "" {
		if defaultValue != nil {
			return *defaultValue
		}
		log.Fatalf("No binding %s found", name)
	}
	return u
}

// Get all Places
func (p *PlacesRepository) FindAll() ([]models.Place, error) {
	var places []models.Place

	findOptions := options.Find()
	findOptions.SetLimit(100)

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	// Finding multiple documents returns a cursor
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result models.Place
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		places = append(places, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return places, err
}

// Create a new Place
func (p *PlacesRepository) Insert(place models.Place) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, &place)
	fmt.Println("Inserted a single document: ", result.InsertedID)
	return result.InsertedID, err
}

// Delete an existing Place
func (p *PlacesRepository) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.DeleteOne(ctx, filter)
	fmt.Println("Deleted a single document: ", result.DeletedCount)
	return err
}

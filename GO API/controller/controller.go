package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Pushpajit/go-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection credentials (put it into ENV_VAR)
var URi = "mongodb+srv://nexus:264516@cluster0.jbwxcev.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
var DBNAME = "Netflix"
var COLNAME = "watchlist"

// Create mongoDB collection pointer *MOST IMP
var collection *mongo.Collection

// Special method that runs one time only for database conenction
func init() {
	// Client options
	clientoption := options.Client().ApplyURI(URi)

	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientoption)

	// Check for connection err
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected!!")

	// Setting up the collection
	collection = client.Database(DBNAME).Collection(COLNAME)

	fmt.Println("Collection is ready to use!!")

}

// Insert a data into database
func insertMovie(movie *model.Netflix) {
	insertRes, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v row inserted", insertRes.InsertedID)
}

// Get all movies from the database
func getAllMovies() []primitive.M {
	// Find() method returns a cursor to work with.
	curr, err := collection.Find(context.Background(), bson.D{{}}, nil)

	// deffer the close connection.
	defer curr.Close(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for curr.Next(context.Background()) {
		var movie bson.M

		// Check for error and decode.
		if err := curr.Decode(&movie); err != nil {
			log.Fatal(err)
		}

		// Return all the movies.
		movies = append(movies, movie)
	}

	return movies
}

// Home Route
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the darkest side of the world 🌎</h1>"))
}

// Actual controller for insert one movie
func InsertOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decoding JSON data from request body
	var data model.Netflix
	json.NewDecoder(r.Body).Decode(&data)

	// Calling the helper fucn
	insertMovie(&data)

	json.NewEncoder(w).Encode(data)

}

// Actual controller for getAllMovies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Retrive all the movies from the database helper function.
	res := getAllMovies()

	json.NewEncoder(w).Encode(res)
}

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Pallav46/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://pallavkumar6200:Pallav%401709@cluster0.dfjfcrg.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
// const connectionString = "mongodb+srv://pallavkumar6200:Pallav@1709@cluster0.dfjfcrg.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const collName = "watchlist"

// Moat important
var collection *mongo.Collection

// connect with mongodb
func init() {
	// client option
	clientOptions := options.Client().ApplyURI(connectionString)
	// client, err := mongo.Connect(context.Background(), clientOptions)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// err = client.Ping(context.Background(), nil)
	// err = client.Ping(context.TODO(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)
	fmt.Println("Collection  instance is ready")
}

// Mongodb helpers - file

// insert 1 movie
func InsertOneMovie(movie model.Netflix) {
	// create a context
	ctx := context.TODO()
	// insert 1 document
	newMovie, err := collection.InsertOne(ctx, movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie with id: ", newMovie.InsertedID)
}

// update 1 record
func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	// update one document
	updatedMovie, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated 1 movie with id: %s %d times", updatedMovie.UpsertedID, updatedMovie.ModifiedCount)
}

// delete 1 record
func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	// delete one document
	deletedMovie, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted 1 movie with count: %d", deletedMovie.DeletedCount)
}

// delete all record
func DeleteAllMovie() int64 {
	// delete all documents
	deletedMovies, err := collection.DeleteMany(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted all movies with count: %d", deletedMovies.DeletedCount)

	return deletedMovies.DeletedCount
}

// find all movies
func FindAllMovies() []primitive.M {
	// find all documents
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cursor.Next(context.TODO()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.TODO())

	return movies
}

// Actual controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	movies := FindAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := DeleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
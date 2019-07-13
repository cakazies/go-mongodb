package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	bson "github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/local/go-mongo/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Student struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

func main() {

	log.Println("Go-mongo is running ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn := options.Client().ApplyURI("mongodb://localhost:27017")
	db, err := mongo.Connect(ctx, conn)
	utils.FindErrors(err, "DB is not connect ")
	client = db
	router := mux.NewRouter()
	router.HandleFunc("/student", CreateStudent).Methods(http.MethodPost)
	router.HandleFunc("/students", GetStudents).Methods(http.MethodGet)
	router.HandleFunc("/student/{id}", GetStudent).Methods(http.MethodGet)
	router.HandleFunc("/student/{id}/update", UpdateStudent).Methods(http.MethodPost)
	router.HandleFunc("/student/{id}", DeleteStudent).Methods(http.MethodDelete)

	http.ListenAndServe("127.0.0.1:8000", router)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var stud Student
	err := json.NewDecoder(r.Body).Decode(&stud)
	utils.FindErrors(err, "Error Decode")
	collection := client.Database("go-mongo").Collection("student")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println(ctx.Err())
	result, err := collection.InsertOne(ctx, stud)
	utils.FindErrors(err, "Insert Student Error")
	json.NewEncoder(w).Encode(result)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var students []Student
	collection := client.Database("go-mongo").Collection("student")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	// utils.FindErrors(err, "Error Decode")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message" : "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var stud Student
		cursor.Decode(&stud)
		students = append(students, stud)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message" : "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	utils.FindErrors(err, "Error Decode")

	var student Student
	collection := client.Database("go-mongo").Collection("student")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, Student{ID: id}).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message" : "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(student)
}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	utils.FindErrors(err, "Error Decode")
	var student Student
	err = json.NewDecoder(r.Body).Decode(&student)
	update := bson.M{"$set": student}
	collection := client.Database("go-mongo").Collection("student")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, Student{ID: id}, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message" : "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(result)

}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	utils.FindErrors(err, "Error Read Params")

	collection := client.Database("go-mongo").Collection("student")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, Student{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message" : "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(w).Encode(result)
}

package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/local/go-mongo/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const table = "student"

func CreateStudent(c echo.Context) error {
	var stud utils.Student
	err := json.NewDecoder(c.Request().Body).Decode(&stud)
	utils.FindErrors(err, "Error Decode")
	collection := DB.Database("go-mongo").Collection(table)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, stud)
	utils.FindErrors(err, "Insert Student Error")
	i := result.InsertedID
	return c.JSON(http.StatusOK, utils.Resp(nil, fmt.Sprintf("Insert Success with %s", i)))
}

func GetStudents(c echo.Context) error {
	var students []utils.Student
	collection := DB.Database("go-mongo").Collection(table)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Resp(nil, fmt.Sprintf("Error get Data : %s", err)))
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var stud utils.Student
		cursor.Decode(&stud)
		students = append(students, stud)
	}
	if err := cursor.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Resp(nil, fmt.Sprintf("Error get Data: %s", err)))
	}
	return c.JSON(http.StatusOK, utils.Resp(students, "Request Success"))
}

func GetStudent(c echo.Context) error {
	params := c.Param("id")
	id, err := primitive.ObjectIDFromHex(params)
	utils.FindErrors(err, "Error Decode")

	var student []utils.Student
	var stud utils.Student

	collection := DB.Database("go-mongo").Collection(table)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, utils.Student{ID: id}).Decode(&stud)
	student = append(student, stud)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return c.JSON(http.StatusInternalServerError, utils.Resp(nil, fmt.Sprintf("ID : %s doesn't exist", params)))
		}
		return c.JSON(http.StatusInternalServerError, utils.Resp(nil, fmt.Sprintf("Error get Per Data : %s", err)))
	}
	return c.JSON(http.StatusOK, utils.Resp(student, "Request Success"))
}

func UpdateStudent(c echo.Context) error {
	params := c.Param("id")
	id, err := primitive.ObjectIDFromHex(params)
	utils.FindErrors(err, "Error Decode")
	var student utils.Student

	err = json.NewDecoder(c.Request().Body).Decode(&student)
	update := bson.M{"$set": student}
	collection := DB.Database("go-mongo").Collection(table)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(ctx, utils.Student{ID: id}, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Resp(nil, fmt.Sprintf("Error get Per Data : %s", err)))
	}
	i := result.UpsertedID
	return c.JSON(http.StatusOK, utils.Resp(nil, fmt.Sprintf("Update Success with : %s", i)))
}

func DeleteStudent(c echo.Context) error {
	params := c.Param("id")
	id, err := primitive.ObjectIDFromHex(params)
	utils.FindErrors(err, "Error Read Params")

	collection := DB.Database("go-mongo").Collection(table)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, utils.Student{ID: id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.Resp(nil, fmt.Sprintf("Error get Per Data : %s", err)))
	}
	i := result.DeletedCount
	return c.JSON(http.StatusOK, utils.Resp(nil, fmt.Sprintf("Update Success with : %d", i)))
}

package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/labstack/echo"
	"github.com/cakazies/go-mongodb/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const table = "student"

// CreateStudent function for insert data in table student
func CreateStudent(c echo.Context) error {
	var stud utils.Student
	err := json.NewDecoder(c.Request().Body).Decode(&stud)
	utils.LogError(err, "Error Decode data json")
	collection := DB.Database("go-mongo").Collection(table)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, stud)
	utils.LogError(err, "Insert Student Error")
	i := result.InsertedID
	return c.JSON(http.StatusOK, utils.Resp(nil, fmt.Sprintf("Insert Success with %s", i)))
}

// GetStudents for get data student
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

// GetStudent function for get sata single student
func GetStudent(c echo.Context) error {
	params := c.Param("id")
	id, err := primitive.ObjectIDFromHex(params)
	utils.LogError(err, "Error Get ID Data")

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

// UpdateStudent function for update data student
func UpdateStudent(c echo.Context) error {
	params := c.Param("id")
	id, err := primitive.ObjectIDFromHex(params)
	utils.LogError(err, "Error Get id Data")
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

// DeleteStudent function for delete permanent data student
func DeleteStudent(c echo.Context) error {
	params := c.Param("id")
	id, err := primitive.ObjectIDFromHex(params)
	utils.LogError(err, "Error Read Params")

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

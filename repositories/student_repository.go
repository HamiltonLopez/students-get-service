package repositories

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "example.com/students-get-service/models"
    "log"
    "os"  
)

type StudentRepository struct {
    collection *mongo.Collection
}

func NewStudentRepository() *StudentRepository {
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatal("MONGO_URI not set in environment")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("school").Collection("students")
    return &StudentRepository{collection}
}

func (repo *StudentRepository) GetAllStudents() ([]models.Student, error) {
    var students []models.Student
    cursor, err := repo.collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var student models.Student
        if err := cursor.Decode(&student); err != nil {
            return nil, err
        }
        students = append(students, student)
    }

    return students, nil
}


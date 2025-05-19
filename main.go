package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "example.com/students-get-service/controllers"
    "example.com/students-get-service/services"
    "example.com/students-get-service/repositories"
)

func init() {
    godotenv.Load()
}

func main() {
    repo := repositories.NewStudentRepository()
    service := services.NewStudentService(repo)
    controller := controllers.NewStudentController(service)
    
    r := mux.NewRouter()

    r.HandleFunc("/students", controller.GetStudents).Methods("GET")

    fmt.Println("Servicio GET Students escuchando en el puerto 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}


package controllers

import (
    "encoding/json"
    "net/http"
    "example.com/students-get-service/services"
)

type StudentController struct {
    Service *services.StudentService
}

func NewStudentController(service *services.StudentService) *StudentController {
    return &StudentController{
        Service: service,
    }
}

func (c *StudentController) GetStudents(w http.ResponseWriter, r *http.Request) {
    students, err := c.Service.GetStudents()
    if err != nil {
        http.Error(w, "Error al obtener estudiantes", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message":  "Estudiantes obtenidos correctamente",
        "students": students,
    })
}


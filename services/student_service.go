package services

import (
    "example.com/students-get-service/models"
    "example.com/students-get-service/repositories"
)

type StudentServiceInterface interface {
    GetStudents() ([]models.Student, error)
}

type StudentService struct {
    repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
    return &StudentService{repo}
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
    return s.repo.GetAllStudents()
}


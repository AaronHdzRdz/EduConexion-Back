package services

import (
    "errors"

    "gorm.io/gorm"
    "gorm/models"
    "github.com/lib/pq"
    "strconv"
    "strings"
)

type StudentService struct {
    db *gorm.DB
}

func NewStudentService(db *gorm.DB) *StudentService {
    return &StudentService{db: db}
}

func (s *StudentService) Create(st *models.Student) error {
    err := s.db.Create(st).Error
    if err != nil {
        if pqErr, ok := err.(*pq.Error); ok {
            if pqErr.Code.Name() == "unique_violation" {
                return errors.New("el correo electrónico ya está en uso") // Mensaje de error personalizado
            }
        }
        return err // Retorna el error original si no es una violación de unicidad
    }
    return nil
}

func (s *StudentService) GetAll() ([]models.Student, error) {
    var list []models.Student
    err := s.db.Find(&list).Error
    return list, err
}

func (s *StudentService) GetByID(id uint) (*models.Student, error) {
    var st models.Student
    if err := s.db.First(&st, id).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &st, nil
}

func (s *StudentService) Update(st *models.Student) error {
    return s.db.Save(st).Error
}

func (s *StudentService) Delete(id uint) error {
    return s.db.Delete(&models.Student{}, id).Error
}

func (s *StudentService) Search(query string) ([]models.Student, error) {
	var students []models.Student
	if id, err := strconv.ParseUint(query, 10, 32); err == nil {
		// Buscar por ID
		err := s.db.Where("id = ?", id).Find(&students).Error
		return students, err
	} else {
		// Buscar por nombre (insensible a mayúsculas)
		err := s.db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(query)+"%").Find(&students).Error
		return students, err
	}
}

func (s *StudentService) GetAllByUser(userID uint) ([]models.Student, error) {
  var list []models.Student
  err := s.db.Where("user_id = ?", userID).Find(&list).Error
  return list, err
}
package services

import (
	"gorm.io/gorm"
	"gorm/models"
	"errors"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Create(user *models.User) error {
	// Verificar si el correo electrónico ya existe
	var existingUser models.User
	result := s.db.Where("email = ?", user.Email).First(&existingUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Si no se encuentra el registro, no hay duplicado, podemos continuar
		} else {
			// Si es otro error, lo retornamos
			return result.Error
		}
	} else {
		// Si no hubo error, significa que se encontró un registro con ese correo
		return errors.New("el correo electrónico ya está registrado")
	}

	// Si el correo no existe, crear el usuario
	result = s.db.Create(user)
	return result.Error
}

//GetAll
func (s *UserService) GetAll() ([]models.User, error) {
	var users []models.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

//GetByID
func (s *UserService) GetByID(id uint) (*models.User, error) {
	var user models.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound){
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

//Update
func (s *UserService) Update(user *models.User) error {
	result := s.db.Save(user)
	return result.Error
}

//Delete
func (s *UserService) Delete(id uint) error {
	result := s.db.Delete(&models.User{}, id)
	return result.Error
}


func (s *UserService) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // Retorna nil, nil para indicar que no se encontró el usuario
		}
		return nil, result.Error
	}
	return &user, nil
}

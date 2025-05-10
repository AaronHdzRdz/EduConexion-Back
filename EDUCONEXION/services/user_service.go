package services

import (
	"errors"

	"gorm.io/gorm"
	"gorm/models"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Create(u *models.User) error {
	return s.db.Create(u).Error
}

func (s *UserService) GetAll() ([]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	return users, err
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	var u models.User
	if err := s.db.First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
    var u models.User
    if err := s.db.Where("email = ?", email).First(&u).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        return nil, err
    }
    return &u, nil
}

func (s *UserService) Update(u *models.User) error {
	return s.db.Save(u).Error
}
func (s *UserService) Delete(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}

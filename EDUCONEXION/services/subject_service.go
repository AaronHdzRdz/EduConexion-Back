package services

import (
	"gorm.io/gorm"
	"gorm/models"
)

type SubjectService struct {
	db *gorm.DB
}

func NewSubjectService(db *gorm.DB) *SubjectService {
	return &SubjectService{db: db}
}

func (s *SubjectService) Create(sub *models.Subject) error {
	return s.db.Create(sub).Error
}

func (s *SubjectService) GetAllByUser(userID uint) ([]models.Subject, error) {
	var list []models.Subject
	err := s.db.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (s *SubjectService) GetByID(id, userID uint) (*models.Subject, error) {
	var sub models.Subject
	err := s.db.Where("id = ? AND user_id = ?", id, userID).First(&sub).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &sub, nil
}

func (s *SubjectService) Update(sub *models.Subject) error {
	return s.db.Save(sub).Error
}

func (s *SubjectService) Delete(id, userID uint) error {
	return s.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Subject{}).Error
}

package store

import "trello-api-client/pkg/models"

func (s *storage) GetListsByUserID(userID int) ([]models.List, error){
	var lists []models.List

	err := s.gorm.Table("list").Where("user_id=?", userID).Scan(&lists).Error

	return lists, err
}

func (s *storage) GetListByID(listID string) (*models.List, error){
	list := models.List{
		ID: listID,
	}

	err := s.gorm.First(&list).Error

	return &list, err
}


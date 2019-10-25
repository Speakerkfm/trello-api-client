package store

import (
	"trello-api-client/pkg/models"
)

func (s *storage) CreateCard(userID int, name, description string, list *models.List) error{
	card := models.Card{
		Name: name,
		Description: description,
		ListID: list.ID,
		UserID: userID,
	}

	return s.gorm.Create(&card).Error
}

func (s *storage) DeleteCardByID(cardID string) error{
	card := models.Card{
		ID: cardID,
	}

	return s.gorm.Delete(&card).Error
}

func (s *storage) UpdateCardStatusByID(userID int, cardID, listID string) error{
	card := models.Card{
		ID: cardID,
	}

	if err := s.gorm.First(&card).Error; err != nil {
		return err
	}

	card.ListID = listID

	return s.gorm.Save(&card).Error
}

func (s *storage) GetCardsByUserID(userID int) ([]models.Card, error){
	var cards []models.Card

	err := s.gorm.Table("card").Scan(&cards).Where("user_id=?", userID).Error
	if err != nil {
		return nil, err
	}

	for idx := range cards{
		list, err := s.GetListByID(cards[idx].ListID)
		if err != nil {
			return nil, err
		}

		cards[idx].Status = list.Name
	}

	return cards, err
}

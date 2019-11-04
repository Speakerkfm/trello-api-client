package store

import (
	"github.com/jinzhu/gorm"
	"trello-api-client/pkg/models"
)

func (s *storage) FindOrCreateUser(trelloID string) (*models.User, error){
	var user models.User
	var err error
	if err = s.gorm.Table("user").Where("trello_id=?", trelloID).Scan(&user).Error; err != nil && err != gorm.ErrRecordNotFound{
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		tx := s.gorm.Begin()

		user.TrelloID = trelloID

		if err := tx.Create(&user).Error; err != nil{
			tx.Rollback()

			return nil, err
		}

		list := models.List{
			Name: "to do",
			UserID: user.ID,
		}

		if err := tx.Create(&list).Error; err != nil {
			tx.Rollback()

			return nil, err
		}

		list = models.List{
			Name: "done",
			UserID: user.ID,
		}

		if err := tx.Create(&list).Error; err != nil {
			tx.Rollback()

			return nil, err
		}

		tx.Commit()
	}

	return &user, nil
}

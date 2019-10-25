package own

import (
	"gitlab.loc/xsolla-login/go-xsolla-login/pkg/log"
	"trello-api-client/pkg/models"
	"trello-api-client/pkg/store"
)

func GetUserCards(userID int) ([]models.Card, error){
	cards, err := store.Storage.GetCardsByUserID(userID)
	if err != nil {
		log.Error().Err(err).Msg("Error while getting user cards")

		return nil, err
	}

	return cards, nil
}

func GetUserLists(userID int) ([]models.List, error){
	lists, err := store.Storage.GetListsByUserID(userID)
	if err != nil {
		log.Error().Err(err).Msg("Error while getting user lists")

		return nil, err
	}

	return lists, nil
}

func CreateCard(userID int, cardName, cardDescription string) error {
	lists, err := store.Storage.GetListsByUserID(userID)
	if err != nil || len(lists) == 0{
		log.Error().Err(err).Msg("Error while getting user's list")

		return err
	}

	if err := store.Storage.CreateCard(userID, cardName, cardDescription, &lists[0]); err != nil {
		log.Error().Err(err).Msg("Error while creating user card")

		return err
	}

	return nil
}

func DeleteUserCard(cardID string) error{
	if err := store.Storage.DeleteCardByID(cardID); err != nil {
		log.Error().Err(err).Msg("Error while deleting user card")

		return err
	}

	return nil
}

func UpdateCardStatus(userID int, cardID, listID string) error{
	if err := store.Storage.UpdateCardStatusByID(userID, cardID, listID); err != nil {
		log.Error().Err(err).Msg("Error while deleting user card")

		return err
	}

	return nil
}
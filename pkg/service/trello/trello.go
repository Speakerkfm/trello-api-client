package trello

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/pkg/errors"
	"gitlab.loc/xsolla-login/go-xsolla-login/pkg/log"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"trello-api-client/pkg/models"
	"trello-api-client/pkg/store"
)

const (
	callbackUrl = "%s%s/v1/trello/callback"
	userProfileUrl = "https://api.trello.com/1/members/me"
	boardsUrl   = "https://api.trello.com/1/members/me/boards"
	boardByIdUrl = "https://api.trello.com/1/boards/%s?lists=open"
	cardByIdUrl = "https://api.trello.com/1/cards/%s"
	cardsUrl = "https://api.trello.com/1/cards"
	boardCardsByIdUrl = "https://api.trello.com/1/boards/%s/cards"
)

var authorizeEndpoint = oauth1.Endpoint{
	RequestTokenURL: "https://trello.com/1/OAuthGetRequestToken",
	AuthorizeURL:    "https://trello.com/1/OAuthAuthorizeToken?scope=read,write,account&name=Trello API",
	AccessTokenURL:  "https://trello.com/1/OAuthGetAccessToken",
}

var AuthConfig oauth1.Config

func Config(appHost, appPort, consumerKey, consumerSecret string) {
	AuthConfig = oauth1.Config{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		CallbackURL:    fmt.Sprintf(callbackUrl, appHost, appPort),
		Endpoint:       authorizeEndpoint,
	}
}

func GetUserBoards(token string) ([]models.Board, error) {
	bits, err := doTrelloRequest(boardsUrl, http.MethodGet, token, http.NoBody)
	if err != nil {
		return nil, err
	}

	var boards []models.Board
	err = json.Unmarshal(bits, &boards)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal failed")

		return nil, err
	}

	return boards, nil
}

func GetUserBoard(boardID, token string) (*models.Board, error) {
	bits, err := doTrelloRequest(fmt.Sprintf(boardByIdUrl, boardID), http.MethodGet, token, http.NoBody)
	if err != nil {
		return nil, err
	}

	var board models.Board
	err = json.Unmarshal(bits, &board)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal failed")

		return nil, err
	}

	return &board, nil
}

func DeleteBoardById(boardID, token string) error {
	_, err := doTrelloRequest(fmt.Sprintf(boardByIdUrl, boardID), http.MethodDelete, token, http.NoBody)

	return err
}

func DeleteCardById(cardID, token string) error {
	_, err := doTrelloRequest(fmt.Sprintf(cardByIdUrl, cardID), http.MethodDelete, token, http.NoBody)

	return err
}

func UpdateCardStatusById(cardID, listID, token string) error {
	u, err := url.Parse(fmt.Sprintf(cardByIdUrl, cardID))
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add("idList", listID)
	u.RawQuery = q.Encode()

	_, err = doTrelloRequest(u.String(), http.MethodPut, token, http.NoBody)

	return err
}

func CreateCard(listID, cardName, cardDescription, token string) error {
	u, err := url.Parse(cardsUrl)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add("idList", listID)
	q.Add("name", cardName)
	q.Add("desc", cardDescription)
	u.RawQuery = q.Encode()

	_, err = doTrelloRequest(u.String(), http.MethodPost, token, http.NoBody)

	return err
}

func GetBoardCards(board *models.Board, token string) ([]models.Card, error) {
	bits, err := doTrelloRequest(fmt.Sprintf(boardCardsByIdUrl, board.ID), http.MethodGet, token, http.NoBody)
	if err != nil {
		return nil, err
	}

	var cards []models.Card
	err = json.Unmarshal(bits, &cards)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal failed")

		return nil, err
	}

	for key := range cards {
		cards[key].Status = board.GetStatusByListID(cards[key].ListID)
	}

	return cards, nil
}

func doTrelloRequest(requestUrl, method, token string, body io.Reader) ([]byte, error){
	u, err := url.Parse(requestUrl)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Add("key", AuthConfig.ConsumerKey)
	q.Add("token", token)
	u.RawQuery = q.Encode()

	log.Info().Msg(u.String())

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		log.Error().Err(err).Msg("Request not created")
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("Request failed")
		return nil, err
	}
	defer resp.Body.Close()

	bits, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Bad body")
		return nil, err
	}

	if resp.StatusCode >= 400 {
		log.Error().Msg(fmt.Sprintf("Trello answered with status code %v. Reason: %s", resp.StatusCode, string(bits)))

		return nil, errors.New("Dependency service is unavailable")
	}

	return bits, nil
}


func GetUser(token string) (*models.User, error){
	bits, err := doTrelloRequest(userProfileUrl, http.MethodGet, token, http.NoBody)
	if err != nil {
		return nil, err
	}

	var trelloUser models.TrelloUser
	err = json.Unmarshal(bits, &trelloUser)
	if err != nil {
		log.Error().Err(err).Msg("Unmarshal failed")

		return nil, err
	}

	user, err := store.Storage.FindOrCreateUser(trelloUser.ID)
	if err != nil {
		log.Error().Err(err).Msg("User creating failed")

		return nil, err
	}

	return user, nil
}
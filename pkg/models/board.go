package models

const statusToDo = "ToDo"

type Board struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Lists []List `json:"lists"`
}

func (b *Board) GetStatusByListID(listID string) string {
	for _, list := range b.Lists {
		if listID == list.ID {
			return list.Name
		}
	}

	return statusToDo
}
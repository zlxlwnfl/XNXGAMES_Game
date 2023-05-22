package dto

type GetAllGamesRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

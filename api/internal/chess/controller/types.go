package controller

type UserGames struct {
	Id          string `json:"id"`
	Username    string	`json:"username"`
	White       string `json:"White"`
	Black       string `json:"Black"`
	Date        string `json:"Date"`
	Result      string `json:"Result"`
}
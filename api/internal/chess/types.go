package chess

type User struct {
	Username 	string		`json:"username"`
	Email		string		`json:"email"`
}


type Game struct {
	Id          string `bson:"_id", json:"id"`
	Username    string
	Site        string `json:"Site"`
	BlackElo    string `json:"BlackElo"`
	BlackTitle  string `json:"BlackTitle"`
	ECO         string `json:"ECO"`
	White       string `json:"White"`
	Black       string `json:"Black"`
	EventDate   string `json:"EventDate"`
	Variation   string `json:"Variation"`
	Date        string `json:"Date"`
	BlackTeam   string `json:"BlackTeam"`
	WhiteTeam   string `json:"WhiteTeam"`
	Event       string `json:"Event"`
	Round       string `json:"Round"`
	Result      string `json:"Result"`
	BlackFideID string `json:"BlackFideId"`
	Opening     string `json:"Opening"`
	WhiteElo    string `json:"WhiteElo"`
	WhiteFideID string `json:"WhiteFideId"`
	WhiteTitle  string `json:"WhiteTitle"`
	Moves       []Move `json:"Moves"`
}

type Move struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Promote byte   `json:"promote"`
	San     string `json:"san"`
	Fen     string `json:"fen"`
}

func (receiver *Game) AddMove(m Move) {
	receiver.Moves = append(receiver.Moves, m)
}

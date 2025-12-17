package responses

type TeamInformationResponse struct {
	Team    TeamInformationPublic `json:"team"`
	Players []Player              `json:"players"`
}

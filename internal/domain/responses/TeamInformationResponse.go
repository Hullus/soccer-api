package responses

type TeamInformationResponse struct {
	Team    TeamInformation `json:"team"`
	Players []Player        `json:"players"`
}

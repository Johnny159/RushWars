package RushWars

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	Client        = &http.Client{}
	BaseURL       = "https://api.rushstats.com/v1/"
	GetPlayerURL  = "player/"
	GetTeamURL    = "team/"
	SearchTeamURL = "search/team/"
	TopPlayerURL  = "leaderboard/players"
)

type rushWars struct {
	token string
}

func NewRushWars(token string) (rushWars, error) {
	return rushWars{token,}, nil
}

func (rw rushWars) MakeRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("Authorization", rw.token)
	resp, err := Client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (rw rushWars) GetPlayer(tag string) (Player, error) {
	body, err := rw.MakeRequest(BaseURL + GetPlayerURL + tag)
	if err != nil {
		return Player{}, err
	}

	var player Player
	if err := json.Unmarshal(body, &player); err != nil {
		return Player{}, err
	}

	return player, nil
}

func (rw rushWars) GetTeam(tag string) (Team, error) {
	body, err := rw.MakeRequest(BaseURL + GetTeamURL + tag)
	if err != nil {
		return Team{}, err
	}

	var team Team
	if err := json.Unmarshal(body, &team); err != nil {
		return Team{}, err
	}

	return team, nil
}

func (rw rushWars) SearchTeam(keyword string) ([]TeamSearch, error) {
	body, err := rw.MakeRequest(BaseURL + SearchTeamURL + keyword)
	if err != nil {
		return []TeamSearch{}, err
	}

	var teams []TeamSearch
	if err := json.Unmarshal(body, &teams); err != nil {
		return []TeamSearch{}, err
	}

	return teams, nil
}

func (rw rushWars) TopPlayers() ([]TopPlayer, error) {
	body, err := rw.MakeRequest(BaseURL + TopPlayerURL)
	if err != nil {
		return []TopPlayer{}, err
	}

	var topPlayers []TopPlayer
	if err := json.Unmarshal(body, &topPlayers); err != nil {
		return []TopPlayer{}, err
	}

	return topPlayers, nil
}

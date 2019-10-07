package RushWars

type Member struct {
	ID                 ID     `json:"id"`
	Tag                string `json:"tag"`
	Name               string `json:"name"`
	Role               string `json:"role"`
	ExpLevel           int    `json:"expLevel"`
	Stars              int    `json:"stars"`
	ResearchInput      int    `json:"researchInput"`
	LastLogin          int    `json:"lastLogin"`
	TimeSinceLastLogin string `json:"timeSinceLastLogin"`
}

type Team struct {
	ID             ID       `json:"id"`
	Tag            string   `json:"tag"`
	Name           string   `json:"name"`
	BadgeID        int      `json:"badgeId"`
	BadgeIDUrl     string   `json:"badgeIdUrl"`
	MembersCount   int      `json:"membersCount"`
	Score          int      `json:"score"`
	RequiredScore  int      `json:"requiredScore"`
	DominationsWon int      `json:"dominationsWon"`
	Description    string   `json:"description"`
	Members        []Member `json:"members"`
}

type TeamSearch struct {
	Tag           string `json:"tag"`
	Name          string `json:"name"`
	Type          int    `json:"type"`
	MembersCount  int    `json:"membersCount"`
	Score         int    `json:"score"`
	RequiredStars int    `json:"requiredStars"`
}

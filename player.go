package RushWars

type ID struct {
	High							int `json:"high"`
	Low								int `json:"low"`
}

type Base struct {
	Name							string	`json:"name"`
	SCID							int		`json:"scid"`
	URL								string	`json:"url"`
}

type Defense struct {
	Base
	Level							int		`json:"level"`
	Cards							int		`json:"cards"`
}

type Commander struct {
	Defense
	TimeUntilFullyHealed			int	`json:"timeUntilFullyHealed"`
	SecondsUntilFullyHealed 		int `json:"secondsUntilFullyHealed"`
}

type PlayerTeam struct {
	ID								ID		`json:"id"`
	Tag								string	`json:"tag"`
	Name							string	`json:"name"`
	BadgeID							int		`json:"badgeId"`
	Role							string	`json:"role"`
}

type Variable struct {
	ChopperLevel					int			`json:"chopperLevel"`
	TotalAttacks					int			`json:"totalAttacks"`
	TotalAttacksWon					int			`json:"totalAttacksWon"`
	TotalAttacksLost				int 		`json:"totalAttacksLost"`
	TotalDefensesWon				int 		`json:"totalDefensesWon"`
	TotalDefensesLost				int			`json:"totalDefensesLost"`
	TotalDefenses					int			`json:"totalDefenses"`
	PendingDefenseStars 			int			`json:"pendingDefenseStars"`
	CurrentDefenseBoxStars			int			`json:"currentDefenseBoxStars"`
	CurrentStars					int 		`json:"currentStars"`
	CycleOffset						int			`json:"cycleOffset"`
	GoldMineTimerTicks				int			`json:"goldMineTimerTicks"`
	TotalFreeChestsOpened			int			`json:"totalFreeChestsOpened"`
	FavoriteCard					Base		`json:"favoriteCard"`
	FavoriteTroop					Base		`json:"favoriteTroop"`
	FavoriteAirdrop					Base		`json:"favoriteAirdrop"`
	AttackStars						int			`json:"attackStars"`
	DefenseStars					int			`json:"defenseStars"`
	TotalDominationAttacks			int			`json:"totalDomination_attacks"`
	DominationAttacksWon			int			`json:"dominationAttacksWon"`
	DominationAttacksLost			int			`json:"dominationAttacksLost"`
	PendingDominationChest			int			`json:"pendingDominationChest"`
	PendingDominationGems			int			`json:"pendingDominationGems"`
	PendingDominationGolds			int			`json:"pendingDominationGolds"`
	PendingDominationOwnStats		int			`json:"pendingDominationOwnStats"`
	PendingDominationOpponentStats	int			`json:"pendingDominationOpponentStats"`
	PendingDominationResult			int			`json:"pendingDominationResult"`
	ChestsSinceNewCard				int			`json:"chestsSinceNewCard"`
	DefenseReady					int			`json:"defenseReady"`
	DominationTeamWinCount			int			`json:"dominationTeamWinCount"`
	DominationTeamLoseCount			int			`json:"dominationTeamLoseCount"`
	DominationGoldsLooted			int			`json:"dominationGoldsLooted"`
	DominationStars					int			`json:"dominationStars"`
	TotalGoldDonated				int			`json:"totalGoldDonated"`
	TotalGoldLooted					int			`json:"totalGoldLooted"`
	SeasonPoints					int			`json:"seasonPoints"`
	ObstaclesRemoved				int			`json:"obstaclesRemoved"`
	HQLevel							int			`json:"hqLevel"`
}

type Player struct {
	ID                       		ID          `json:"id"`
	Tag                      		string      `json:"tag"`
	Exp                      		int         `json:"exp"`
	Name                    		string      `json:"name"`
	Team                    		PlayerTeam  `json:"team"`
	Stars                   		int         `json:"stars"`
	ExpLevel               		  	int         `json:"expLevel"`
	SecondsSinceLastActivity 		int         `json:"secondsSinceLastActivity"`
	TimeSinceLastActivity    		string      `json:"timeSinceLastActivity"`
	Defenses                		[]Defense   `json:"defenses"`
	Variables                		Variable    `json:"variables"`
	Airdrops                 		[]Defense   `json:"airdrops"`
	Commanders              		[]Commander `json:"commanders"`
	Troops                   		[]Defense   `json:"troops"`
}


type TopPlayerTeam struct {
	ID								ID			`json:"id"`
	Tag								string		`json:"tag"`
	Name							string		`json:"name"`
}

type TopPlayer struct {
	ID       						ID            `json:"id"`
	Tag      						string        `json:"tag"`
	Name    						string        `json:"name"`
	Position 						int           `json:"position"`
	Stars    						int           `json:"stars"`
	Team     						TopPlayerTeam `json:"team"`
}

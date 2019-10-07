# RushWars API Wrapper

A simple, yet effective and powerful wrapper written in GoLang with zero dependencies. Make it concurrent, or run it as it is. Your wish. **This wrapper is highly extensible**.

## Usage

Add this into your project using dep

```bash
dep ensure -add github.com/smartclash/RushWars
```

And extend this project into your `main` function like this. Replace `{token}` with your token that you obtained from the rushstats dashboard

```go
package main

import (
	...
	rusher "github.com/smartclash/RushWars"
)


func main() {
	rush, err := rusher.NewRushWars("{token}")
	if err != nil {
		fmt.Println("Woah, we got an error", err.Error())
	}
	
	...
}
```

### Get Player Data from Tag

The function `GetPlayer("{token}")` returns an instance of `Player` type structure and an error if exists

```go
player, err := rush.GetPlayer("2PP")
```

### Get Team Data from Tag

The function `GetTeam("{token}")` returns an instance of `Team` type structure and an error if exists

```go
team, err := rush.GetTeam("2PP")
```

### Search Teams from Query

The function `SearchTeam("{Query}")` returns an instance of `[]TeamSearch` type structure and an error if exists

```go
teams, err := rush.SearchTeam("Hello")
```

### Get top players in leaderboard

The function `TopPlayers()` returns an instance of `[]TopPlayer` type structure and an error if exists.

```go
players, err := rush.TopPlayers()
```

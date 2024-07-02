package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/cmd"
)

type ListPlayers struct {
}

func (c ListPlayers) Run(source cmd.Source, output *cmd.Output) {
	var result string
	players := system.Server.Players()
	playerCount := len(players)
	output.Printf("There are %v/%v players online:", playerCount, system.Server.MaxPlayerCount())

	// Avoid output overflow
	if playerCount <= 10{
		for _, p := range players{
			result += p.Data().Username + " "
		}
	}else{
		result = ">Too many players to display<"
	}

	output.Printf(result)
}

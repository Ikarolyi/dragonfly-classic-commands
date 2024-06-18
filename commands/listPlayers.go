package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/utils"
	"github.com/df-mc/dragonfly/server/cmd"
)

type ListPlayers struct {
}


func (c ListPlayers) Run(source cmd.Source, output *cmd.Output) {
	var result string
	for _, p := range utils.PlayerList {
		result += p.Data().Username + " "
	}

	output.Printf(result)
}


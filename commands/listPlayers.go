package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/cmd"
)

type ListPlayers struct {
}

func (c ListPlayers) Run(source cmd.Source, output *cmd.Output) {
	var result string
	for _, p := range system.Server.Players() {
		result += p.Data().Username + " "
	}

	output.Printf(result)
}

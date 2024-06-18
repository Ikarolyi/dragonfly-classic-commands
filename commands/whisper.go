package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Whisper struct {
	Target []cmd.Target
	Message string
}


func (c Whisper) Run(source cmd.Source, output *cmd.Output) {
	for _, t := range c.Target{
		p, _ := t.(*player.Player)
		p.SendCommandOutput(output)
	}

	output.Printf(c.Message)
}
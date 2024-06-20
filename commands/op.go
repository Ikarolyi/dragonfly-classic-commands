package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Op struct {
	Player []cmd.Target
}

func (c Op) Run(source cmd.Source, output *cmd.Output) {
	// if !permissions.AuthSource(source, permissions.LEVEL_HOST, output){
	// 	return
	// }
	for _, t := range c.Player {
		p := t.(*player.Player)
		permissions.SetLevel(p.XUID(), permissions.LEVEL_OPERATOR)
	}
}

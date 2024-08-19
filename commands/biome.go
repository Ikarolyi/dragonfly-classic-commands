package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
)

type Biome struct{}


func (c Biome) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}

	sender := system.GetSender(source)

	pos := sender.Position()
	w := sender.World()
	
	output.Print(w.Biome(cube.PosFromVec3(pos)).String())
}


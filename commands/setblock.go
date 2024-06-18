package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/utils"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type Setblock struct {
	Position mgl64.Vec3
	TileName string
}

func (c Setblock) Run(source cmd.Source, output *cmd.Output) {
	sender := utils.GetSender(source)
	if sender == nil{
		return
	}

	w := sender.World()

	b, err := world.BlockByName(c.TileName, nil)
	if err{
		return
	}

	pos := cube.Pos{int(c.Position[0]), int(c.Position[1]), int(c.Position[2])}

	w.SetBlock(pos, b, &world.SetOpts{DisableBlockUpdates: true, DisableLiquidDisplacement: true})
}
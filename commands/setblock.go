package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type SetblockNormal struct {
	Position mgl64.Vec3
	TileName system.BlockEnum
	SetBlockMode cmd.Optional[SetBlockModeEnum]
}

type SetBlockStates struct{
	Position mgl64.Vec3
	TileName system.BlockEnum
	BlockStates int16
	SetBlockMode cmd.Optional[SetBlockModeEnum]
}

type SetBlockModeEnum string

func (sbme SetBlockModeEnum) Type() string{
	return "SetBlockMode"
}

func (sbme SetBlockModeEnum) Options(source cmd.Source) []string{
	return []string{"destroy", "keep", "replace"}
}

func (c SetblockNormal) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}

	sender := system.GetSender(source)
	if sender == nil {
		return
	}


	w := sender.World()

	b := c.TileName.GetBlock()

	pos := cube.Pos{int(c.Position[0]), int(c.Position[1]), int(c.Position[2])}

	w.SetBlock(pos, b, &world.SetOpts{DisableBlockUpdates: true, DisableLiquidDisplacement: true})
}

func (c SetBlockStates) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}

	sender := system.GetSender(source)
	if sender == nil {
		return
	}


	w := sender.World()

	b := c.TileName.GetBlockWithMeta(c.BlockStates)

	if b == nil{
		output.Errorf("No meta state >%v< found for %s blocks", c.BlockStates, c.TileName)
	}

	pos := cube.Pos{int(c.Position[0]), int(c.Position[1]), int(c.Position[2])}

	w.SetBlock(pos, b, &world.SetOpts{DisableBlockUpdates: true, DisableLiquidDisplacement: true})
}

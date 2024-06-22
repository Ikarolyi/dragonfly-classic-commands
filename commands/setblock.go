package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type SetblockNormal struct {
	Position mgl64.Vec3
	TileName system.BlockEnum
	SetBlockMode cmd.Optional[SetBlockModeEnum] `cmd:""`
}

type SetBlockStates struct{
	Position mgl64.Vec3
	TileName system.BlockEnum
	BlockStates int16
	SetBlockMode cmd.Optional[SetBlockModeEnum] `cmd:""`
}

type SetBlockModeEnum string

func (sbme SetBlockModeEnum) Type() string{
	return "SetBlockMode"
}

func (sbme SetBlockModeEnum) Options(source cmd.Source) []string{
	return []string{"destroy", "keep", "replace"}
}

func RunSetblock(Position mgl64.Vec3, TileName system.BlockEnum, BlockStates int16, BlockStatesUsed bool, SetBlockMode cmd.Optional[SetBlockModeEnum], source cmd.Source, output *cmd.Output){
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}
	
	sender := system.GetSender(source)
	if sender == nil {
		return
	}


	w := sender.World()

	var b world.Block

	if BlockStatesUsed{
		b = TileName.GetBlockWithMeta(BlockStates)

		if b == nil{
			output.Errorf("No meta state >%v< found for %s blocks", BlockStates, TileName)
			return
		}
	}else{
		b = TileName.GetBlock()
	}

	pos := system.Vec2Cube(Position)

	setBlockModeVal := SetBlockMode.LoadOr("replace")

	switch setBlockModeVal{
		case "replace":
			w.SetBlock(pos, b, &world.SetOpts{DisableBlockUpdates: true, DisableLiquidDisplacement: true})
		case "destroy":
			return
		case "keep":
			blockOnPos, success := system.GetBlockOnPos(Position, w)
			if !success{
				output.Error("Can't check for air: chunk not loaded")
				return
			}

			airBlock := block.Air{}

			if blockOnPos == airBlock{
				w.SetBlock(pos, b, &world.SetOpts{DisableBlockUpdates: true, DisableLiquidDisplacement: true})
			}else{
				output.Errorf("Block at [%v, %v, %v] is not air", pos[0], pos[1], pos[2])
			}
	}
}

func (c SetblockNormal) Run(source cmd.Source, output *cmd.Output) {
	RunSetblock(c.Position, c.TileName, 0, false, c.SetBlockMode, source, output)
}

func (c SetBlockStates) Run(source cmd.Source, output *cmd.Output) {
	RunSetblock(c.Position, c.TileName, c.BlockStates, false, c.SetBlockMode, source, output)
}

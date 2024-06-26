package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Clear struct {
	Player cmd.Optional[[]cmd.Target]
}


func (c Clear) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}

	sender := system.GetSender(source)

	target, targetSet := c.Player.Load()
	var targetPlayer *player.Player

	if !targetSet{
		if sender == nil{
			// Abbort if called from console and no player specified
			return
		}else{
			targetPlayer = sender
		}
	}else{
		targetPlayer, _ = target[0].(*player.Player)
	}

	targetPlayer.Inventory().Clear()
	targetPlayer.Armour().Clear()
}


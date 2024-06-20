package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
)

type Teleport struct {
	Victim      cmd.Optional[[]cmd.Target]
	Target      cmd.Optional[[]cmd.Target]
	Destination cmd.Optional[mgl64.Vec3]
}

func (c Teleport) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}
	
	sender := system.GetSender(source)

	victimTarget, victimSet := c.Victim.Load()
	var victim *player.Player
	if !victimSet {
		victim = sender
	} else {
		victim, _ = victimTarget[0].(*player.Player)
	}

	var destination mgl64.Vec3
	destination, desdestinationSet := c.Destination.Load()

	target, targetSet := c.Target.Load()
	if !targetSet && !desdestinationSet {
		destination = sender.Position()
	} else if !desdestinationSet {
		targetP, _ := target[0].(*player.Player)
		destination = targetP.Position()
	}

	victim.Teleport(destination)
}

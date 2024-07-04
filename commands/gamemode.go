package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type GamemodeByEnum struct {
	Gamemode GamemodeEnum
	Player cmd.Optional[[]cmd.Target]
}

type GamemodeEnum string

func (ge GamemodeEnum) Type() string{
	return "SetBlockMode"
}

func (ge GamemodeEnum) Options(source cmd.Source) []string{
	return []string{"survival", "creative", "adventure", "default", "spectator", "s", "c", "d", "0", "1", "2", "5"}
}

func (c GamemodeByEnum) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}

	var targetGamemode world.GameMode
	var defaultGamemode = false
	switch c.Gamemode{
		case "survival", "s", "0":
			targetGamemode = world.GameModeSurvival
		case "creative", "c", "1":
			targetGamemode = world.GameModeCreative
		case "adventure", "2":
			targetGamemode = world.GameModeAdventure
		case "default", "d", "5":
			defaultGamemode = true
		case "spectator":
			targetGamemode = world.GameModeSpectator
	}
	
	victim, victimSpec := c.Player.Load()
	if !victimSpec{
		sender := system.GetSender(source)
		if defaultGamemode{
			sender.SetGameMode(sender.World().DefaultGameMode())
		}else{
			sender.SetGameMode(targetGamemode)
		}
	}else{
		for _, v := range victim{
			player, isPlayer := v.(*player.Player)
			if isPlayer{
				if defaultGamemode{
					player.SetGameMode(player.World().DefaultGameMode())
				}else{
					player.SetGameMode(targetGamemode)
				}
			}
		}
	}
}
package handlers

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/player"
)

func HandleJoin(p *player.Player) {
	system.PlayerList = append(system.PlayerList, p)
}

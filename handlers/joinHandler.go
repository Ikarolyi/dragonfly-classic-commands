package handlers

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/utils"
	"github.com/df-mc/dragonfly/server/player"
)


func HandleJoin(p *player.Player) {
	utils.PlayerList = append(utils.PlayerList, p)
}
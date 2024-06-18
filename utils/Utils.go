package utils

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

var PlayerList []*player.Player

func GetSender(s cmd.Source) *player.Player {
	if p, isPlayer := s.(*player.Player); isPlayer{
		return p
	}

	return nil
}
package classicCommands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/commands"
	"github.com/Ikarolyi/dragonfly-classic-commands/handlers"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

func Init() {
	cmd.Register(cmd.New("clear", "Clears the full inventory of a player", []string{}, commands.Clear{}))
	cmd.Register(cmd.New("setblock", "...", []string{}, commands.Setblock{}))
	cmd.Register(cmd.New("list", "Lists the players on the server", []string{}, commands.ListPlayers{}))
	cmd.Register(cmd.New("tell", "Sends a private message to one or more players", []string{"msg", "w"}, commands.Whisper{}))
}

func PassAccept(p *player.Player){
	handlers.HandleJoin(p)
}
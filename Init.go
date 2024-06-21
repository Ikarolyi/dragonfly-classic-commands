package classicCommands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/commands"
	"github.com/Ikarolyi/dragonfly-classic-commands/handlers"
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

func Init() {	
	permissions.Init()

	cmd.Register(cmd.New("clear", "Clears the full inventory of a player", []string{}, commands.Clear{}))
	cmd.Register(cmd.New("setblock", "Changes a block to an other block", []string{}, commands.SetblockNormal{}, commands.SetBlockStates{}))
	cmd.Register(cmd.New("list", "Lists the players on the server", []string{}, commands.ListPlayers{}))
	cmd.Register(cmd.New("tell", "Sends a private message to one or more players", []string{"msg", "w"}, commands.Whisper{}))
	cmd.Register(cmd.New("teleport", "Teleports players", []string{"tp"}, commands.Teleport{}))
	cmd.Register(cmd.New("op", "Promotes a player's permission level", []string{}, commands.Op{}))
	cmd.Register(cmd.New("permission", "Reloads permisson.json without buffering", []string{"ops"}, commands.Permission{}))
	cmd.Register(cmd.New("time", "Changes or queries the world's game time", []string{}, commands.AddTime{}, commands.QueryTime{}, commands.SetTimeToTime{}, commands.SetTimeToAmount{}))
}

func Save() {
	permissions.Buffer()
}

func PassAccept(p *player.Player){
	handlers.HandleJoin(p)
}
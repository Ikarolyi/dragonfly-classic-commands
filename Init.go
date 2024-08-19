package classicCommands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/commands"
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

func Init(srv *server.Server) {
	permissions.Init()
	system.Server = srv

	cmd.Register(cmd.New("clear", "Clears the full inventory of a player", []string{}, commands.Clear{}))
	cmd.Register(cmd.New("setblock", "Changes a block to another block", []string{}, commands.SetblockNormal{}, commands.SetBlockStates{}))
	cmd.Register(cmd.New("list", "Lists the players on the server", []string{}, commands.ListPlayers{}))
	cmd.Register(cmd.New("tell", "Sends a private message to one or more players", []string{"msg", "w"}, commands.Whisper{}))
	cmd.Register(cmd.New("teleport", "Teleports players", []string{"tp"}, commands.TeleportSenderToTarget{}, commands.TeleportVictimToTarget{}, commands.TeleportSenderToPos{}, commands.TeleportVictimToPos{}))
	cmd.Register(cmd.New("op", "Promotes a player's permission level", []string{}, commands.Op{}))
	cmd.Register(cmd.New("permission", "Reloads permisson.json without buffering", []string{"ops"}, commands.Permission{}))
	cmd.Register(cmd.New("time", "Changes or queries the world's game time", []string{}, commands.AddTime{}, commands.QueryTime{}, commands.SetTimeToTime{}, commands.SetTimeToAmount{}))
	cmd.Register(cmd.New("gamemode", "Sets a player's game mode", []string{}, commands.GamemodeByEnum{}))
	cmd.Register(cmd.New("say", "Sends a message to the chat", []string{}, commands.Say{}))
	cmd.Register(cmd.New("weather", "Sets the weather", []string{}, commands.SetWeather{}))
	cmd.Register(cmd.New("biome", "Queries the biome", []string{}, commands.Biome{}))
}

func Save() {
	permissions.Buffer()
}

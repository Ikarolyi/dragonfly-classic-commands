package classicCommands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/Ikarolyi/dragonfly-classic-commands/commands"
)

func RegisterClassicCommands() {
	cmd.Register(cmd.New("clear", "Clears the full inventory of a player", []string{}, commands.Clear{}))
	cmd.Register(cmd.New("setblock", "...", []string{}, commands.Setblock{}))

	println("command registered")
}
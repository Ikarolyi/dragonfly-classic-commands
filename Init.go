package classicCommands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/Ikarolyi/dragonfly-classic-commands/commands"
)

func RegisterClassicCommands() {
	cmd.Register(cmd.New("clear", "An example of using commands", []string{"eg"}, commands.Clear{}))

	println("command registered")
}
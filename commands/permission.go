package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/df-mc/dragonfly/server/cmd"
)

type Permission struct {
	Reload cmd.SubCommand `cmd:"reload"`
}

func (c Permission) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_HOST, output){
		return
	}
	
	permissions.Init()
	output.Print("Permissions have been reloaded.")
}
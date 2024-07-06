package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
)

type Say struct {
	Message cmd.Varargs
}

func (c Say) Run(source cmd.Source, output *cmd.Output) {
	sender := system.GetSender(source)

	if sender == nil{
		// Command executed from the console
		chat.Global.WriteString("<§c§lServer§r> " + string(c.Message))
	}else{
		sender.Chat(c.Message)
	}
}

package permissions

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/cmd"
)

type PermissionLevel struct{
	uint8
}

var(
	LEVEL_NORMAL = PermissionLevel{0}
	LEVEL_OPERATOR = PermissionLevel{1}
	LEVEL_HOST = PermissionLevel{2}
	LEVEL_AUTOMATION = PermissionLevel{3}
	LEVEL_ADMIN = PermissionLevel{4}
)

func GetLevel(xuid string) PermissionLevel{
	value, found :=	cache[xuid]
	if found{
		return value
	}else{
		value = Get(xuid)
		cache[xuid] = value
		return value
	}
}

func SetLevel(xuid string, level PermissionLevel) {
	cache[xuid] = level
}

func AuthSource(source cmd.Source, min PermissionLevel, output *cmd.Output) bool{
	p := system.GetSender(source)

	if p == nil{
		// Command sent by the owner
		return true
	}

	level := GetLevel(p.XUID())
	if level.uint8 < min.uint8{
		output.Errorf("Your permission level %v is lower than the command's requirement %v", level, min)

		return false
	}else{
		return true
	}
}

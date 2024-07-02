package commands

import (
	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"

	//    "github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	//    "math"
)

type TeleportSenderToTarget struct {
	Destination    []cmd.Target
	CheckForBlocks cmd.Optional[bool] `cmd:"checkForBlocks"`
}

func (c TeleportSenderToTarget) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
		return
	}
	sender := system.GetSender(source)
	destPlayer, _ := c.Destination[0].(*player.Player)
	destPos := destPlayer.Position()

	if !teleportCheckForBlocksSuccess(destPos, destPlayer.World(), c.CheckForBlocks, output) {
		return
	}

	sender.Teleport(destPos)
}

type TeleportVictimToTarget struct {
	Victim         []cmd.Target
	Destination    []cmd.Target
	CheckForBlocks cmd.Optional[bool] `cmd:"checkForBlocks"`
}

func (c TeleportVictimToTarget) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
		return
	}

	destPlayer, _ := c.Destination[0].(*player.Player)
	destPos := destPlayer.Position()

	if !teleportCheckForBlocksSuccess(destPos, destPlayer.World(), c.CheckForBlocks, output) {
		return
	}

	for _, v := range c.Victim {
		victimPlayer, _ := v.(*player.Player)
		victimPlayer.Teleport(destPos)
	}
}

type TeleportSenderToPos struct {
	Destination    mgl64.Vec3
	CheckForBlocks cmd.Optional[bool] `cmd:"checkForBlocks"`
}

func (c TeleportSenderToPos) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
		return
	}
	sender := system.GetSender(source)
	destPos := c.Destination

	if !teleportCheckForBlocksSuccess(destPos, sender.World(), c.CheckForBlocks, output) {
		return
	}

	sender.Teleport(destPos)
}

type TeleportVictimToPos struct {
	Victim         []cmd.Target
	Destination    mgl64.Vec3
	CheckForBlocks cmd.Optional[bool] `cmd:"checkForBlocks"`
}

func (c TeleportVictimToPos) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
		return
	}
	sender := system.GetSender(source)
	destPos := c.Destination

	if !teleportCheckForBlocksSuccess(destPos, sender.World(), c.CheckForBlocks, output) {
		return
	}

	for _, v := range c.Victim {
		victimPlayer, _ := v.(*player.Player)
		victimPlayer.Teleport(destPos)
	}
}

//
// Specifying rotation is not possible with the current api

//type TeleportSenderRotSpec struct {
//    Destination    mgl64.Vec3
//    YRot           cmd.Optional[float64]
//    XRot           cmd.Optional[float64]
//    CheckForBlocks cmd.Optional[bool]
//}
//
//func (c TeleportSenderRotSpec) Run(source cmd.Source, output *cmd.Output) {
//    if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
//        return
//    }
//
//    sender := system.GetSender(source)
//    destPos := c.Destination
//
//    if !teleportCheckForBlocksSuccess(destPos, sender.World(), c.CheckForBlocks, output) {
//        return
//    }
//
//    sender.Teleport(destPos)
//    playerRot := sender.Rotation()
//    setPlayerRot(sender, c.YRot.LoadOr(playerRot.Yaw()), c.XRot.LoadOr(playerRot.Pitch()))
//}
//
//type TeleportSenderFacingPos struct {
//    Destination    mgl64.Vec3
//    Facing         mgl64.Vec3
//    CheckForBlocks cmd.Optional[bool]
//}
//
//func (c TeleportSenderFacingPos) Run(source cmd.Source, output *cmd.Output) {
//    if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
//        return
//    }
//
//    sender := system.GetSender(source)
//    destPos := c.Destination
//
//    if !teleportCheckForBlocksSuccess(destPos, sender.World(), c.CheckForBlocks, output) {
//        return
//    }
//
//    sender.Teleport(destPos)
//    playerEyePos := entity.EyePosition(sender)
//    yaw, pitch := LookAtPos(playerEyePos, c.Facing)
//
//    setPlayerRot(sender, yaw, pitch)
//}
//
//type TeleportSenderFacingTarget struct {
//    Destination    mgl64.Vec3
//    Facing         []cmd.Target
//    CheckForBlocks cmd.Optional[bool]
//}
//
//func (c TeleportSenderFacingTarget) Run(source cmd.Source, output *cmd.Output) {
//    if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
//        return
//    }
//
//    sender := system.GetSender(source)
//    destPos := c.Destination
//
//    if !teleportCheckForBlocksSuccess(destPos, sender.World(), c.CheckForBlocks, output) {
//        return
//    }
//
//    sender.Teleport(destPos)
//    playerEyePos := entity.EyePosition(sender)
//    yaw, pitch := LookAtPos(playerEyePos, c.Facing[0].Position())
//
//    setPlayerRot(sender, yaw, pitch)
//}
//
//type TeleportVictimRotSpec struct {
//    Victim         []cmd.Target
//    Destination    mgl64.Vec3
//    YRot           cmd.Optional[float64]
//    XRot           cmd.Optional[float64]
//    CheckForBlocks cmd.Optional[bool]
//}
//
//func (c TeleportVictimRotSpec) Run(source cmd.Source, output *cmd.Output) {
//    if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
//        return
//    }
//
//    sender := system.GetSender(source)
//    destPos := c.Destination
//
//    if !teleportCheckForBlocksSuccess(destPos, sender.World(), c.CheckForBlocks, output) {
//        return
//    }
//
//    for _, v := range c.Victim {
//        victimPlayer, _ := v.(*player.Player)
//        victimPlayer.Teleport(destPos)
//
//        playerRot := victimPlayer.Rotation()
//        setPlayerRot(sender, c.YRot.LoadOr(playerRot.Yaw()), c.XRot.LoadOr(playerRot.Pitch()))
//    }
//}
//
//type TeleportVictimFacingPos struct {
//    Victim         []cmd.Target
//    Destination    mgl64.Vec3
//    Facing         mgl64.Vec3
//    CheckForBlocks cmd.Optional[bool]
//}
//
//func (c TeleportVictimFacingPos) Run(source cmd.Source, output *cmd.Output) {
//    if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
//        return
//    }
//
//    sender := system.GetSender(source)
//    destPos := c.Destination
//
//    if !teleportCheckForBlocksSuccess(destPos, sender.World(), c.CheckForBlocks, output) {
//        return
//    }
//
//    for _, v := range c.Victim {
//        victimPlayer, _ := v.(*player.Player)
//        victimPlayer.Teleport(destPos)
//
//        playerEyePos := entity.EyePosition(victimPlayer)
//        yaw, pitch := LookAtPos(playerEyePos, c.Facing)
//
//        setPlayerRot(victimPlayer, yaw, pitch)
//    }
//}
//
//type TeleportVictimFacingTarget struct {
//    Victim         []cmd.Target
//    Destination    mgl64.Vec3
//    Facing         []cmd.Target
//    CheckForBlocks cmd.Optional[bool]
//}
//
//func (c TeleportVictimFacingTarget) Run(source cmd.Source, output *cmd.Output) {
//    if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output) {
//        return
//    }
//
//    sender := system.GetSender(source)
//    destPos := c.Destination
//
//    if !teleportCheckForBlocksSuccess(destPos, sender.World(), c.CheckForBlocks, output) {
//        return
//    }
//
//    for _, v := range c.Victim {
//        victimPlayer, _ := v.(*player.Player)
//        victimPlayer.Teleport(destPos)
//
//        playerEyePos := entity.EyePosition(victimPlayer)
//        yaw, pitch := LookAtPos(playerEyePos, c.Facing[0].Position())
//
//        setPlayerRot(victimPlayer, yaw, pitch)
//    }
//}

func teleportCheckForBlocksSuccess(pos mgl64.Vec3, w *world.World, CheckForBlocks cmd.Optional[bool], output *cmd.Output) bool {
	if CheckForBlocks.LoadOr(false) {
		blockOnPos := w.Block(system.Vec2Cube(pos))
		posOverBlock := cube.Pos{int(pos[0]), int(pos[1] + 1), int(pos[2])}
		blockOverPos := w.Block(posOverBlock)
		air := block.Air{}
		if blockOnPos != air || blockOverPos != air {
			output.Printf("Unable to teleport to [%v, %v, %v]; the area is not clear", pos[0], pos[1], pos[2])
			return false
		}
	}

	output.Printf("Teleported to [%v, %v, %v]", pos[0], pos[1], pos[2])
	return true
}

//
//func setPlayerRot(p *player.Player, yaw float64, pitch float64) {
//    playerRot := p.Rotation()
//    deltaYaw := yaw - playerRot.Yaw()
//    deltaPitch := pitch - playerRot.Pitch()
//    _, _ = deltaYaw, deltaPitch
//
//    println("Before", int(p.Rotation()[0]), int(p.Rotation()[1]))
//    p.Move(mgl64.Vec3{0, 0, 100}, deltaYaw, deltaPitch)
//    p.SendTitle(title.New("Bruh"))
//    //    p.Move(mgl64.Vec3{0, 0, 0}, 180, 180)
//    println("After", int(p.Rotation()[0]), int(p.Rotation()[1]))
//}
//
//func LookAtPos(playerEyes mgl64.Vec3, facing mgl64.Vec3) (float64, float64) {
//    //	Trigonometry hell
//    difference := playerEyes.Sub(facing)
//    var yaw = math.Atan(difference.Y()/difference.Z()) * (180 / math.Pi)
//    var pitch = math.Atan(difference.X()/difference.Z()) * (180 / math.Pi)
//    return yaw, pitch
//}

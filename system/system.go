package system

import (
	"math/rand"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

var Server *server.Server

func GetSender(s cmd.Source) *player.Player {
	if p, isPlayer := s.(*player.Player); isPlayer{
		return p
	}

	return nil
}

type ItemEnum string

func (ie ItemEnum) Type() string{
	return "Item"
}

func (ie ItemEnum) Options(source cmd.Source) []string{
	var result []string
	for _, it := range world.Items(){
		name, _ := it.EncodeItem()
		result = append(result, name)
	}
	return result
}

func (ie ItemEnum) GetBlock() world.Item{
	for _, it := range world.Items(){
		name, _ := it.EncodeItem()
		if ie == ItemEnum(name){
			return it
		}
	}
	return nil
}

type BlockEnum string

func (be BlockEnum) Type() string{
	return "Block"
}

func (be BlockEnum) Options(source cmd.Source) []string{
	var result []string
	for _, it := range world.Items(){
		if _, isBlock := it.(world.Block); isBlock{
			// result = append(result, item.DisplayName(it, language.English))
			name, _ := it.EncodeItem()
			result = append(result, name)
		}
	}
	return result
}

func (be BlockEnum) GetBlock() world.Block{
	for _, it := range world.Items(){
		if b, isBlock := it.(world.Block); isBlock{
			name, _ := it.EncodeItem()
			if be == BlockEnum(name){
				return b
			}
		}
	}
	return nil
}

func (be BlockEnum) GetBlockWithMeta(targetMeta int16) world.Block{
	for _, it := range world.Items(){
		if b, isBlock := it.(world.Block); isBlock{
			name, meta := it.EncodeItem()
			if be == BlockEnum(name) && meta == targetMeta{
				return b
			}
		}
	}
	return nil
}

func ChunkPosFromBlockPos(pos cube.Pos) world.ChunkPos {
	return world.ChunkPos{int32(pos[0] >> 4), int32(pos[2] >> 4)}
}

func GetChunkOfBlock(pos mgl64.Vec3, w *world.World) (*world.Column, bool){
	chunkPos := ChunkPosFromBlockPos(Vec2Cube(pos))
	v := world.NopViewer{}

	ld := world.NewLoader(1, w, v)
	ld.Move(pos)
	ld.Load(1)
	return ld.Chunk(chunkPos)
}

func GetBlockOnPos(pos mgl64.Vec3, w *world.World) (world.Block, bool){
	column, success := GetChunkOfBlock(pos, w)
	if !success{
		return nil, success
	}
	cubePos := Vec2Cube(pos)
	x := uint8(cubePos[0] % 16)
	y := int16(cubePos[1])
	z := uint8(cubePos[2] % 16)

	block, _ := world.BlockByRuntimeID(column.Chunk.Block(x, y, z, 0))

	return block, true
}

func Vec2Cube(vec mgl64.Vec3) cube.Pos{
	return cube.Pos{int(vec[0]), int(vec[1]), int(vec[2])}
}

func BlockBreakDrops(pos cube.Pos, w *world.World, p *player.Player) {
	b := w.Block(pos)

	breakable, isBreakable := b.(block.Breakable)
	if !isBreakable{
		return
	}

	breakInfo := breakable.BreakInfo()
	perfectTools := []item.Tool{
		item.Axe{Tier: item.ToolTierNetherite},
		item.Hoe{Tier: item.ToolTierNetherite},
		item.Shovel{Tier: item.ToolTierNetherite},
		item.Shears{},
		item.Pickaxe{Tier: item.ToolTierNetherite},
	}

	for _, tool := range perfectTools{
		drops := breakInfo.Drops(tool, nil)
		if len(drops) != 0{
			if breakInfo.BreakHandler != nil {
				breakInfo.BreakHandler(pos, w, p)
			}
			
			if container, ok := b.(block.Container); ok {
				// If the block is a container, it should drop its inventory contents regardless whether the
				// player is in creative mode or not.
				drops = container.Inventory().Items()
				if breakable, ok := b.(block.Breakable); ok {
					if breakable.BreakInfo().Harvestable(tool) {
						drops = append(drops, breakable.BreakInfo().Drops(tool, nil)...)
					}
				}
				container.Inventory().Clear()
			} else if breakable, ok := b.(block.Breakable); ok {
				if breakable.BreakInfo().Harvestable(tool) {
					drops = breakable.BreakInfo().Drops(tool, nil)
				}
			} else if it, ok := b.(world.Item); ok {
				drops = []item.Stack{item.NewStack(it, 1)}
			}

			for _, drop := range drops {
				ent := entity.NewItem(drop, pos.Vec3Centre())
				ent.SetVelocity(mgl64.Vec3{rand.Float64()*0.2 - 0.1, 0.2, rand.Float64()*0.2 - 0.1})
				w.AddEntity(ent)
			}
			return
		}
	}

}
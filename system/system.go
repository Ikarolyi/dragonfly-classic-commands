package system

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)


var PlayerList []*player.Player

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
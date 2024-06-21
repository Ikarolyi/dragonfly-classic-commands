package system

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
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
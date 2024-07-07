package commands

import (
	"math/rand"
	"time"

	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/df-mc/dragonfly/server/cmd"
)

// It's impossible to query the weather (world.conf is private)

// type QueryWeather struct {
// 	Query cmd.SubCommand `cmd:"query"`
// }

type SetWeather struct {
	Target WeatherTypeEnum
	Duriation cmd.Optional[int]
}

type WeatherTypeEnum string

func (wte WeatherTypeEnum) Type() string{
	return "Weather"
}

func (wte WeatherTypeEnum) Options(source cmd.Source) []string{
	return []string{"clear", "rain", "thunder"}
}

func (c SetWeather) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}

	w := source.World()

	duriation := time.Duration(c.Duriation.LoadOr(rand.Intn(600) + 300)) * time.Second

	switch c.Target {
		case "clear":
			// It's impossible to stop weather for a specific duriation
			w.StopRaining()
			w.StopThundering()
		case "rain":
			w.StartRaining(duriation)
		case "thunder":
			w.StartThundering(duriation)
	}
}
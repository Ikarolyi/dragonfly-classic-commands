package commands

import (
	"fmt"

	"github.com/Ikarolyi/dragonfly-classic-commands/permissions"
	"github.com/Ikarolyi/dragonfly-classic-commands/system"
	"github.com/df-mc/dragonfly/server/cmd"
)

type AddTime struct{
	Add cmd.SubCommand `cmd:"add"`
	Amount int
}

type QueryTime struct{
	Query cmd.SubCommand `cmd:"query"`
	QueryMode QueryTimeEnum `cmd:""`
}

type QueryTimeEnum string

func (qtm QueryTimeEnum) Options(source cmd.Source) []string{
	return []string{"daytime", "gametime", "day"}
}

func (qtm QueryTimeEnum) Type() string{
	return ""
}

type SetTimeToTime struct{
	Set cmd.SubCommand `cmd:"set"`
	TimeSpec TimeSpecEnum `cmd:"TimeSpec"`
}

type TimeSpecEnum string

func (tse TimeSpecEnum) Options(source cmd.Source) []string{
	return []string{"day","night","noon","midnight","sunrise","sunset"}
}

func (tse TimeSpecEnum) Type() string{
	return "TimeSpec"
}



type SetTimeToAmount struct{
	Set cmd.SubCommand `cmd:"set"`
	TimeAmount int `cmd:"amout"`
}

func (c AddTime) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}

	w := system.GetSender(source).World()
	w.SetTime(w.Time() + c.Amount)


	output.Printf("Added %v to the time", c.Amount)
}

func (c QueryTime) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}
	w := system.GetSender(source).World()

	var result string
	var queryModeName string
	switch c.QueryMode{
		case "daytime":
			result = fmt.Sprint(w.Time() % 24000)
			queryModeName = "Daytime"
		case "gametime":
			result = fmt.Sprint(w.Time() % 2147483647)
			queryModeName = "Gametime"
		case "day":
			result = fmt.Sprint((w.Time() / 24000) % 2147483647)
			queryModeName = "Day"
	}

	output.Printf("%s is %s", queryModeName, result)
}

func (c SetTimeToAmount) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}

	w := system.GetSender(source).World()
	time := w.Time()
	lastMidnight := time - (time % 24000)

	targetTime := lastMidnight + c.TimeAmount

	w.SetTime(targetTime)

	output.Printf("Set the time to %v", targetTime)
}

func (c SetTimeToTime) Run(source cmd.Source, output *cmd.Output) {
	if !permissions.AuthSource(source, permissions.LEVEL_OPERATOR, output){
		return
	}

	var amount int
	switch c.TimeSpec{
		case "day":
			amount = 1000
		case "night":
			amount = 13000
		case "noon":
			amount = 6000
		case "midnight":
			amount = 18000
		case "sunrise":
			amount = 23000
		case "sunset":
			amount = 12000
	}

	w := system.GetSender(source).World()
	time := w.Time()
	lastMidnight := time - (time % 24000)
	targetTime := lastMidnight + amount

	w.SetTime(targetTime)

	output.Printf("Set the time to %v", targetTime)
}
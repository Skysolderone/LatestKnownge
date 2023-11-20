//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// func Initialize() DefaultTest {
// 	wire.Build(DefaultTest)
// 	return DefaultTest
// }

func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}

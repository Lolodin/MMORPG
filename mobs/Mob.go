package mobs

import "Test/chunk"

type Mob struct {
	chunk.Coordinate
	Name string
	health int
	speed int
	strength int
	Parent *MobGenerator
}


func NewMob(name string, X,Y int, g *MobGenerator) Mob {
	var mob Mob
	mob.X = X
	mob.Y = Y
	mob.Name = name
	mob.health = 100
	mob.strength = 1
	mob.speed = 15
	mob.Parent = g
	return mob
}
func(m *Mob) Die() {
	m.Parent.CurrentMob = nil
}
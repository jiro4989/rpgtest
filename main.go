package main

import (
	"fmt"
)

type Attacker interface {
	Attack(p *Param)
}

type Equiper interface {
	Equip(p *Param)
}

type Actor struct {
	Param
}

func (a *Actor) Attack(p *Param) {
	dmg := a.ATK*4 - p.DEF*2
	p.HP -= dmg
}

type Enemy struct {
	Param
}

func (e *Enemy) Attack(p *Param) {
	dmg := e.ATK*4 - p.DEF*2
	p.HP -= dmg
}

// Param は能力値
type Param struct {
	MHP int64
	MMP int64
	HP  int64
	MP  int64
	ATK int64
	DEF int64
	MAT int64
	MDF int64
	SPD int64
	LUK int64
}

type Weapon struct {
	ID   int
	Name string
	Param
}

type Weapons []Weapon

type Armor struct {
	ID   int
	Name string
	Param
}

type Armors []Armor

var weapons = Weapons{
	Weapon{ID: 0, Name: "なし", Param: Param{HP: 0, MP: 0, ATK: 0, DEF: 0, MAT: 0, MDF: 0, SPD: 5, LUK: 0}},
	Weapon{ID: 1, Name: "ひのきのぼう", Param: Param{HP: 0, MP: 0, ATK: 5, DEF: 0, MAT: 0, MDF: 0, SPD: 0, LUK: 0}},
}

func main() {
	actor := Actor{
		Param: Param{
			MHP: 100,
			MMP: 100,
			HP:  100,
			MP:  100,
			ATK: 25,
			DEF: 10,
			MAT: 100,
			MDF: 100,
			SPD: 100,
			LUK: 100,
		},
	}
	enemy := Enemy{
		Param: Param{
			MHP: 100,
			MMP: 100,
			HP:  100,
			MP:  100,
			ATK: 10,
			DEF: 4,
			MAT: 100,
			MDF: 100,
			SPD: 100,
			LUK: 100,
		},
	}
	fmt.Println("hello")

	fmt.Println("actor:", actor)
	fmt.Println("enemy:", enemy)

	actor.Attack(&enemy.Param)
	enemy.Attack(&actor.Param)

	fmt.Println("actor:", actor)
	fmt.Println("enemy:", enemy)
}

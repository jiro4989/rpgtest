package main

import "fmt"

type Attacker interface {
	Attack(p *Param)
}

type Actor struct {
	Param
}

func (a *Actor) Attack(p *Param) {
	dmg := a.Param.ATK*4 - p.DEF*2
	p.HP -= dmg
}

type Enemy struct {
	Param
}

func (e *Enemy) Attack(p *Param) {
	dmg := e.Param.ATK*4 - p.DEF*2
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
	Param
}
type Armor struct {
	Param
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

package main

import (
	"fmt"
)

type Attacker interface {
	Attack(p *Param)
}

type Equiper interface {
	EquipWeapon(w Weapon) Param
	EquipArmor(a Armor) Param
	RemoveWeapon() Weapon
	RemoveArmor() Armor
}

type Actor struct {
	MHP int64
	MMP int64
	Param
	Weapon
	Armor
}

func (a *Actor) Attack(p *Param) {
	dmg := a.ATK*4 - p.DEF*2
	p.HP -= dmg
}

func (a *Actor) EquipWeapon(w Weapon) Param {
	a.RemoveWeapon()
	diff := a.Param.Sub(&w.Param)
	a.Param = a.Param.Add(&w.Param)
	a.Weapon = w
	return diff
}

func (a *Actor) EquipArmor(am Armor) Param {
	a.RemoveArmor()
	diff := a.Param.Sub(&am.Param)
	a.Param = a.Param.Add(&am.Param)
	a.Armor = am
	return diff
}

func (a *Actor) RemoveWeapon() Weapon {
	w := a.Weapon
	a.Param = a.Param.Sub(&w.Param)
	a.Weapon = weapons[0]
	return w
}

func (a *Actor) RemoveArmor() Armor {
	am := a.Armor
	a.Param = a.Param.Sub(&am.Param)
	a.Armor = armors[0]
	return am
}

type Enemy struct {
	MHP int64
	MMP int64
	Param
}

func (e *Enemy) Attack(p *Param) {
	dmg := e.ATK*4 - p.DEF*2
	p.HP -= dmg
}

// Param は能力値
type Param struct {
	HP  int64
	MP  int64
	ATK int64
	DEF int64
	MAT int64
	MDF int64
	SPD int64
	LUK int64
}

func (p *Param) Add(p2 *Param) Param {
	return Param{
		HP:  p.HP + p2.HP,
		MP:  p.MP + p2.MP,
		ATK: p.ATK + p2.ATK,
		DEF: p.DEF + p2.DEF,
		MAT: p.MAT + p2.MAT,
		MDF: p.MDF + p2.MDF,
		SPD: p.SPD + p2.SPD,
		LUK: p.LUK + p2.LUK,
	}
}

func (p *Param) Sub(p2 *Param) Param {
	return Param{
		HP:  p.HP - p2.HP,
		MP:  p.MP - p2.MP,
		ATK: p.ATK - p2.ATK,
		DEF: p.DEF - p2.DEF,
		MAT: p.MAT - p2.MAT,
		MDF: p.MDF - p2.MDF,
		SPD: p.SPD - p2.SPD,
		LUK: p.LUK - p2.LUK,
	}
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
	Weapon{ID: 0, Name: "なし"},
	Weapon{ID: 1, Name: "ひのきのぼう", Param: Param{ATK: 5}},
}

var armors = Armors{
	Armor{ID: 0, Name: "なし"},
	Armor{ID: 1, Name: "布の服", Param: Param{DEF: 3}},
}

func main() {
	actor := Actor{
		MHP: 100,
		MMP: 100,
		Param: Param{
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
		MHP: 100,
		MMP: 100,
		Param: Param{
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

	fmt.Println("actor:", actor)
	fmt.Println("enemy:", enemy)

	actor.EquipWeapon(weapons[1])
	actor.Attack(&enemy.Param)
	enemy.Attack(&actor.Param)

	fmt.Println("actor:", actor)
	fmt.Println("enemy:", enemy)
}

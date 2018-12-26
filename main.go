package main

import "fmt"

// Unit はキャラユニットを定義
type Unit struct {
	Param
	Weapon
	Armor
}

type Attacker interface {
	Attack(p *Param)
}

type Defender interface {
	Defend(p *Param)
}

type Skill int

// Actor/Enemyには固定値としてのパラメータと
// 戦闘の時に動的に変化するパラメータがある
// Paramで装備品と共有する能力値とすると
// 装備品にMHP, MMPは不要になる。
// どうする？

type Actor struct {
	Param
	Weapon
	Armor
}

func (a *Actor) Attack(p *Param) {
	dmg := a.ATK*4 - p.DEF*2
	p.HP -= dmg
}

type Enemy struct {
	Param
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

type Weapon struct {
	Param
}
type Armor struct {
	Param
}

func main() {
	fmt.Println("hello")
}

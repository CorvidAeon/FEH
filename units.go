package FEH

//Unit represents a character on the map with stats
type Unit struct {
	MaxHP       int
	HP          int
	Atk         int
	Speed       int
	Def         int
	Res         int
	Mobility    int //Armor 1, Infantry 2, Flier 2, Cavalry 3
	ActionTaken bool
	SkillSet    Skills
}

//Move operates on unit in map
func (a Unit) Move() bool {
	return false //TODO
}

//(Atk*Eff)+(Atk*Eff*Adv-Mit)*ClassMod
//((Atk*Eff+Atk*Eff*Adv+SpcStat*SpcMod-(Mit+Mit*MitMod))*(1+OffMult)+OffFlat)*(1-DefMult)-DefFlat
func (a Unit) Attack(d Unit) {
	dmg := (a.Atk - d.Def) //Fix later, too simple
	if d.MaxHP < dmg {
		d.MaxHP = 0
	} else {
		d.MaxHP -= dmg
	}
}

//Battle may have multiple attacks or just 1
func (a Unit) Battle(d Unit) {

}

func (a Unit) Assist() {

}

type Skills struct {
	Weap Weapon
	Ast  Assist
	Spec Special
}

type Weapon struct {
	Might int
	Range int
}

//weaponType
const (
	R = iota
	G
	B
	C
)

//unit type
const (
	Infantry = iota
	Cavalry
	Flier
	Armored
)

type Assist struct {
}

type Special struct {
}

//StatusEffect contains modifiers to stats
type StatusEffect struct {
	MaxHP       int
	HP          int
	Atk         int
	Speed       int
	Def         int
	Res         int
	ActionTaken bool
}

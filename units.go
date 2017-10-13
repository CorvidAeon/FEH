package FEH

//Unit represents a character on the map with stats
type Unit struct {
	Name        string
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
func (a *Unit) Move() bool {
	return false //TODO
}

//(Atk*Eff)+(Atk*Eff*Adv-Mit)*ClassMod
//((Atk*Eff+Atk*Eff*Adv+SpcStat*SpcMod-(Mit+Mit*MitMod))*(1+OffMult)+OffFlat)*(1-DefMult)-DefFlat
func (a *Unit) Attack(d *Unit) {
	dmg := (a.Atk - d.Def) //Fix later, too simple
	if d.HP < dmg {
		d.HP = 0
	} else {
		d.HP -= dmg
	}
}

//Battle may have multiple attacks or just 1
//need to account for certain skills
//oh and check if the unit can actually counter, more info
func (a *Unit) Battle(d *Unit) {
	var aTurns, dTurns = 1, 1
	if a.Speed >= d.Speed+5 {
		aTurns = 2
	} else if d.Speed >= a.Speed+5 {
		dTurns = 2
	}
	rounds := aTurns + dTurns
	for rounds > 0 {
		if aTurns > 0 {
			a.Attack(d)
			//Brave weapons only double for the attacker.
			if a.SkillSet.Weap.Brave {
				a.Attack(d)
			}
			aTurns--
			rounds--
		}
		if dTurns > 0 {
			d.Attack(a)
			dTurns--
			rounds--
		}
	}
}

//Later swap to work on multiple units, an effect can be attached to almost anything
func (a *Unit) PostBattle(d *Unit) {

}

func (a *Unit) Assist() {

}

//Skills set attached to character
type Skills struct {
	Weap Weapon
	Ast  Assist
	Spec Special
}

type Weapon struct {
	Might int
	Range int
	Brave bool
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
	Infantry = iota //iota starts at 0 and goes up
	Cavalry
	Flier
	Armored
)

//damage type
const (
	Magic = iota
	Physical
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

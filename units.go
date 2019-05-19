package FEH

import "fmt"

//Apparently these don't need to be exported to work in the same package... fix after finishing.

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
	unimplemented()
	return false //TODO
}

//(Atk*Eff)+(Atk*Eff*Adv-Mit)*ClassMod
//((Atk*Eff+Atk*Eff*Adv+SpcStat*SpcMod-(Mit+Mit*MitMod))*(1+OffMult)+OffFlat)*(1-DefMult)-DefFlat
//Special can activate during an attack or elsewhere depending on type
func (a *Unit) Attack(d *Unit) {
	var mit = 0 //Mitigation stat
	if a.SkillSet.Weap.DmgType == Magic {
		mit = d.Res
	} else {
		mit = d.Def
	}
	dmg := (a.Atk - mit) //Fix later, too simple
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
	totalTurns := aTurns + dTurns
	for totalTurns > 0 {
		if aTurns > 0 {
			a.Attack(d)
			//Brave weapons only double for the attacker.
			if a.SkillSet.Weap.Brave {
				a.Attack(d)
			}
			aTurns--
			totalTurns--
		}
		if dTurns > 0 {
			d.Attack(a)
			dTurns--
			totalTurns--
		}
	}
}

//ActivateSpecial activates units special when cooldown reaches 0
//Currently just one target but will be updated later
//Never call when a's HP is 0
func (a *Unit) ActivateSpecial(d *Unit) {
	if a.SkillSet.Spec.PostBattleEffect {
		return
	}
	//Damage section
	dmg := a.SkillSet.Spec.Dmg
	for _, target := range a.SkillSet.Spec.DmgTargets {
		switch target {
		case Self: //self targetting specials only ever reduce hp to 1
			a.HP -= dmg
			if a.HP <= 0 {
				a.HP = 1
			}
		case Enemy:
			if dmg > d.HP {
				d.HP = 0
			} else {
				d.HP -= dmg
			}
		case RangeAllies:
			unimplemented()
		case RangeEnemies:
			unimplemented()
		}
	}
	healing := a.SkillSet.Spec.Healing
	for _, target := range a.SkillSet.Spec.HealingTargets {
		switch target {
		case Self:
			if a.HP+healing > a.MaxHP {
				a.HP = a.MaxHP
			} else {
				a.HP += healing
			}
		case Enemy: //Doesn't make sense to heal enemies yet
			unimplemented()
		case RangeAllies: //Unimplemented
			unimplemented()
		case RangeEnemies:
			unimplemented()
		}
	}
	status := a.SkillSet.Spec.Status
	for _, target := range a.SkillSet.Spec.StatusTargets {
		switch target {
		case Self:
			fmt.Printf("%+v\n", status)
			unimplemented()
		case Enemy:
			unimplemented()
		case RangeAllies:
			unimplemented()
		case RangeEnemies:
			unimplemented()
		}
	}

}

//Later swap to work on multiple units, an effect can be attached to almost anything
func (a *Unit) PostBattle(d *Unit) {

}

func (a *Unit) Assist() {
	unimplemented()
}

//Skills set attached to character
type Skills struct {
	Weap Weapon
	Ast  Assist
	Spec Special
}

type Weapon struct {
	Might   int
	Range   int
	Brave   bool
	DmgType DamageType
}

//WeaponColor has possible values: R G B C
type WeaponColor int

//WeaponColor
const (
	R WeaponColor = iota
	G
	B
	C
)

//UnitType can be one of: Infantry, Cavalry, Flier, or Armored
type UnitType int

//unit types
const (
	Infantry UnitType = iota
	Cavalry
	Flier
	Armored
)

//DamageType specifies magic 0 or physical 1
type DamageType int

//damage type
const (
	Magic    = DamageType(0)
	Physical = DamageType(1)
)

//TargetType is used for determining activatables targets
type TargetType int

//target type
const (
	Self TargetType = iota
	Enemy
	Ally
	RangeAllies
	RangeEnemies
	RangeAll //certain items will use this probably
)

type Assist struct {
	Name            string
	Description     string
	Target          TargetType
	Move            int
	MoveTarget      TargetType
	SpecCooldownMod int //changes turns for special to activate
}

//Special can be many different types of effects such as healing, stat changes, or damage.
type Special struct {
	Name             string
	Description      string
	Cooldown         int
	Healing          int
	HealingTargets   []TargetType
	Dmg              int
	DmgTargets       []TargetType
	Status           StatusEffect
	StatusTargets    []TargetType
	PostBattleEffect bool
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

func unimplemented() {
	panic("Unimplemented")
}

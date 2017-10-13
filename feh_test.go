package FEH

import (
	"fmt"
	"testing"
)

func TestAttack(t *testing.T) {
	a := Unit{Name: "Ike", MaxHP: 20, HP: 20, Atk: 20, Speed: 20, Def: 10, Res: 10, Mobility: 2, ActionTaken: false}
	d := Unit{Name: "Nephenee", MaxHP: 20, HP: 20, Atk: 15, Speed: 20, Def: 10, Res: 10, Mobility: 2, ActionTaken: false}
	a.Attack(&d)
	fmt.Printf("%+v\n", d)
	if d.HP != 10 {
		t.Error("Expected HP: 10\nCurrent HP: ", d.HP)
	}
}

func TestBattle(t *testing.T) {
	a := Unit{Name: "Ike", MaxHP: 20, HP: 20, Atk: 20, Speed: 20, Def: 10, Res: 10, Mobility: 2, ActionTaken: false}
	d := Unit{Name: "Nephenee", MaxHP: 20, HP: 20, Atk: 15, Speed: 15, Def: 10, Res: 10, Mobility: 2, ActionTaken: false}
	a.Battle(&d)
	fmt.Println("Test battle")
	if d.HP != 0 {
		t.Error("Expected HP: 0\nCurrent HP: ", d.HP)
	}
}

package kobold

import "github.com/mphox-phoxdev/kobold-generator/utils"

type Weapon struct {
	Name   string  `json:"name"`
	Dam    int     `json:"dam"`
	Edges  []Edge  `json:"edges,omitempty"`
	Bogies []Bogie `json:"bogies,omitempty"`
}

func (kobold *Kobold) assignRandomWeapon() {
	if kobold.Skills["shoot"] {

	} else {
		//1/3 chance of rolling on the danger chart
		if utils.D6() <= 2 {
			kobold.assignDangerousRandomWeapon()
			kobold.DeathCheques++
		} else {
			kobold.assignBasicRandomWeapon()
		}
	}
}

func (kobold *Kobold) assignBasicRandomWeapon() {
	d6 := utils.D6()
	if kobold.Skills["heft"] {
		d6--
	}

	switch d6 {
	case 0:
		kobold.Weapon = Weapon{
			Name:   "Ham Hok, Rotting",
			Dam:    2,
			Bogies: []Bogie{Bogie{Name: "Foul Smelling"}},
		}
	case 1:
		kobold.Weapon = Weapon{
			Name: "Spork-Tipped Spear",
			Dam:  2,
		}
	case 2:
		kobold.Weapon = Weapon{
			Name:   "Kitchen Knife",
			Dam:    2,
			Bogies: []Bogie{Bogie{Name: "Razor"}},
		}
	case 3:
		kobold.Weapon = Weapon{
			Name:  "Hammer",
			Dam:   1,
			Edges: []Edge{Edge{Name: "Useful"}},
		}
	case 4:
		kobold.Weapon = Weapon{
			Name:  "Cooking Utensil",
			Dam:   1,
			Edges: []Edge{Edge{Name: "Cook"}},
		}
	case 5:
		kobold.Weapon = Weapon{
			Name:   "Dead Rat",
			Dam:    0,
			Bogies: []Bogie{Bogie{Name: "Foul Smelling"}},
		}
	case 6:
		kobold.Weapon = Weapon{
			Name: "Diddly Squat",
			Dam:  -1,
		}
	}
}

func (kobold *Kobold) assignDangerousRandomWeapon() {
	d6 := utils.D6()
	if kobold.Skills["duel"] {
		d6--
	}

	switch d6 {
	case 0:
		kobold.Weapon = Weapon{
			Name: "Axe, Big",
			Dam:  4,
			Bogies: []Bogie{
				Bogie{Name: "Big"},
				Bogie{Name: "Bulky"}},
		}
	case 1:
		kobold.Weapon = Weapon{
			Name:  "Iron Skillet",
			Dam:   3,
			Edges: []Edge{Edge{Name: "Cook"}},
		}
	case 2:
		kobold.Weapon = Weapon{
			Name:   "Bone, Large",
			Dam:    3,
			Bogies: []Bogie{Bogie{Name: "Big"}},
		}
	case 3:
		kobold.Weapon = Weapon{
			Name:  "Chain",
			Dam:   2,
			Edges: []Edge{Edge{Name: "Climb"}},
		}
	case 4:
		kobold.Weapon = Weapon{
			Name:  "Club, Heavy",
			Dam:   2,
			Edges: []Edge{Edge{Name: "Bash"}},
		}
	case 5:
		kobold.Weapon = Weapon{
			Name:  "Cleaver, Butcher",
			Dam:   2,
			Edges: []Edge{Edge{Name: "Cook"}},
		}
	case 6:
		kobold.Weapon = Weapon{
			Name:   "Flail",
			Dam:    2,
			Bogies: []Bogie{Bogie{Name: "Flail"}},
		}
	}
}

func (kobold *Kobold) assignBasicRangedWeapon() {
	d6 := utils.D6()

	switch d6 {

	case 1:
		kobold.Weapon = Weapon{
			Name: "A Rock",
			Dam:  1,
		}
	case 2, 3, 4:
		kobold.Weapon = Weapon{
			Name:   "Sling Shot",
			Dam:    0,
			Bogies: []Bogie{Bogie{Name: "Stones"}},
		}
	case 5, 6:
		kobold.Weapon = Weapon{
			Name:   "Sling",
			Dam:    0,
			Bogies: []Bogie{Bogie{Name: "Stones"}},
			Edges:  []Edge{Edge{Name: "Cook"}},
		}
	}
}

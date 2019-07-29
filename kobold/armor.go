package kobold

import "github.com/mphox-phoxdev/kobold-generator/utils"

type Armor struct {
	Name   string  `json:"name"`
	Hits   int     `json:"hits"`
	Edges  []Edge  `json:"edges,omitempty"`
	Bogies []Bogie `json:"bogies,omitempty"`
}

func (kobold *Kobold) assignRandomArmor() {
	//1/3 chance of rolling on the danger chart
	if utils.D6() <= 2 {
		kobold.assignDangerousRandomArmor()
		kobold.DeathCheques++
	} else {
		kobold.assignBasicRandomArmor()
	}
}

func (kobold *Kobold) assignBasicRandomArmor() {
	d6 := utils.D6()
	if kobold.Skills["sport"] {
		d6--
	}

	switch d6 {
	case 0:
		kobold.Armor = Armor{
			Name:  "Spare Tire",
			Hits:  3,
			Edges: []Edge{Edge{Name: "Bouncy"}},
		}
	case 1:
		kobold.Armor = Armor{
			Name:   "Small Shield",
			Hits:   4,
			Bogies: []Bogie{Bogie{Name: "Item"}},
		}
	case 2:
		kobold.Armor = Armor{
			Name: "Leather Vest",
			Hits: 3,
		}
	case 3:
		kobold.Armor = Armor{
			Name: "Hoodie",
			Hits: 2,
		}
	case 4:
		kobold.Armor = Armor{
			Name: "Kid's Clothes",
			Hits: 2,
		}
	case 5:
		kobold.Armor = Armor{
			Name: "Socks",
			Hits: 1,
		}
	case 6:
		kobold.Armor = Armor{
			Name: "Nekkid!",
			Hits: 0,
		}
	}
}

func (kobold *Kobold) assignDangerousRandomArmor() {
	d6 := utils.D6()
	if kobold.Skills["lift"] {
		d6--
	}

	switch d6 {
	case 0:
		kobold.Armor = Armor{
			Name: "Bone Mail",
			Hits: 8,
		}
	case 1:
		kobold.Armor = Armor{
			Name:   "Beer Barrel",
			Hits:   10,
			Bogies: []Bogie{Bogie{Name: "Item"}},
		}
	case 2:
		kobold.Armor = Armor{
			Name:   "Colander Helm",
			Hits:   8,
			Bogies: []Bogie{Bogie{Name: "Item"}},
		}
	case 3:
		kobold.Armor = Armor{
			Name:   "Chain Vest",
			Hits:   8,
			Bogies: []Bogie{Bogie{Name: "Jangly"}},
		}
	case 4:
		kobold.Armor = Armor{
			Name:  "Leather Jacket",
			Hits:  6,
			Edges: []Edge{Edge{Name: "Fonzie"}},
		}
	case 5:
		kobold.Armor = Armor{
			Name:  "Leather Apron",
			Hits:  6,
			Edges: []Edge{Edge{Name: "Backpack"}},
		}
	case 6:
		kobold.Armor = Armor{
			Name: "Kite Shield",
			Hits: 12,
			Bogies: []Bogie{
				Bogie{Name: "Big"},
				Bogie{Name: "Bulky"}},
		}
	}
}

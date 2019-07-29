package kobold

import "github.com/mphox-phoxdev/kobold-generator/utils"

type Gear struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (kobold *Kobold) assignRandomGear() {
	//1/3 chance of rolling on the danger chart
	if utils.D6() <= 2 {
		kobold.assignDangerousRandomGear()
		kobold.DeathCheques++
	} else {
		kobold.assignBasicRandomGear()
	}
}

func (kobold *Kobold) assignBasicRandomGear() {
	d6 := utils.D6()
	if kobold.Skills["nature"] {
		d6--
	}

	switch d6 {
	case 0:
		kobold.Gear = Gear{
			Name:        "Dowsing Rod",
			Description: "",
		}
	case 1:
		kobold.Gear = Gear{
			Name:        "Adventurer's Cast-Offs",
			Description: "pg. 42",
		}
	case 2:
		kobold.Gear = Gear{
			Name:        "Backpack",
			Description: "",
		}
	case 3:
		kobold.Gear = Gear{
			Name:        "Sack of Spices",
			Description: "",
		}
	case 4:
		kobold.Gear = Gear{
			Name:        "Rope, 25 Feet, Hemp",
			Description: "",
		}
	case 5:
		kobold.Gear = Gear{
			Name:        "9 Foot Pole",
			Description: "",
		}
	case 6:
		kobold.Gear = Gear{
			Name:        "Lint, Belly Button",
			Description: "",
		}
	}
}

func (kobold *Kobold) assignDangerousRandomGear() {
	d6 := utils.D6()
	if kobold.Skills["dungeon"] {
		d6--
	}

	switch d6 {
	case 0:
		kobold.Gear = Gear{
			Name:        "Booze",
			Description: "",
		}
	case 1:
		kobold.Gear = Gear{
			Name:        "Spell Pages",
			Description: "pg. 34",
		}
	case 2:
		kobold.Gear = Gear{
			Name:        "Circle of Sign Talking",
			Description: "",
		}
	case 3:
		kobold.Gear = Gear{
			Name:        "Bag of Holding: Chickens",
			Description: "",
		}
	case 4:
		kobold.Gear = Gear{
			Name:        "Bracers of Offense",
			Description: "",
		}
	case 5:
		kobold.Gear = Gear{
			Name:        "Thieves' Tool",
			Description: "",
		}
	case 6:
		kobold.Gear = Gear{
			Name:        "Cup of Elemental Summing: Milk",
			Description: "pg. 44",
		}
	}
}

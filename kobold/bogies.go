package kobold

import "github.com/mphox-phoxdev/kobold-generator/utils"

type Bogie struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func RandomBogie() Bogie {
	switch d6 := utils.D6(); d6 {
	case 1:
		return Bogie{Name: "Animul Foe"}
	case 2:
		return Bogie{Name: "Flammable"}
	case 3:
		return Bogie{Name: "Foul Smelling"}
	case 4:
		return Bogie{Name: "Hungry"}
	case 5:
		return Bogie{Name: "In Heat"}
	case 6:
		return Bogie{Name: "Tastes Like Baby"}
	}
	return Bogie{}
}

package kobold

import "github.com/mphox-phoxdev/kobold-generator/utils"

type Edge struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func RandomEdge() Edge {
	switch d6 := utils.D6(); d6 {
	case 1:
		return Edge{Name: "Animul Chum"}
		break
	case 2:
		return Edge{Name: "Bouncy"}
		break
	case 3:
		return Edge{Name: "Extra Padding"}
		break
	case 4:
		return Edge{Name: "Red Thumb"}
		break
	case 5:
		return Edge{Name: "Troll Blood"}
		break
	case 6:
		return Edge{Name: "Winnig Smile"}
		break
	}
	return Edge{}
}

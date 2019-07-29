package kobold

import (
	"fmt"

	"github.com/mphox-phoxdev/kobold-generator/utils"
)

type Kobold struct {
	Stats        Stats        `json:"stats"`
	HandyNumbers HandyNumbers `json:"handyNumbers"`

	DeathCheques  int             `json:"deathCheques"`
	Role          string          `json:"role"`
	Skills        map[string]bool `json:"skills"`
	playersChoice string
	Edges         []Edge  `json:"edges"`
	Bogies        []Bogie `json:"bogies"`
	Armor         Armor   `json:"armor"`
	Weapon        Weapon  `json:"weapon"`
	Gear          Gear    `json:"gear"`

	Hits int `json:"hits"`
	// Outfit?
}

type Stats struct {
	Brawn      int `json:"brawn"`
	Ego        int `json:"ego"`
	Extraneous int `json:"extraneous"`
	Reflexes   int `json:"reflexes"`
}

type HandyNumbers struct {
	Meat    int `json:"meat"`
	Cunning int `json:"cunning"`
	Luck    int `json:"luck"`
	Agility int `json:"agility"`
}

func (stats *Stats) GenerateHandyNumbersFromStats() HandyNumbers {
	return HandyNumbers{
		Meat:    (stats.Brawn-1)/4 + 1,
		Cunning: (stats.Ego-1)/4 + 1,
		Luck:    (stats.Extraneous-1)/4 + 1,
		Agility: (stats.Reflexes-1)/4 + 1,
	}
}

func GenerateKobold() (kobold Kobold) {
	stats := Stats{
		Brawn:      utils.D6() + utils.D6(),
		Ego:        utils.D6() + utils.D6(),
		Extraneous: utils.D6() + utils.D6(),
		Reflexes:   utils.D6() + utils.D6(),
	}

	kobold = Kobold{
		Stats:        stats,
		HandyNumbers: stats.GenerateHandyNumbersFromStats(),
	}

	kobold.Hits = kobold.Stats.Brawn

	// Skills will be assigned using the
	// Super Express Value Meal Alternate Alternate Kobold Kreation Method
	// from the Even More Things to Kill And Eat

	kobold.Skills = map[string]bool{"cook": true}

	// kobolds with 2 or more cunning get an additional random skill
	kobold.updateSkills(GetRandomSkill())

	// Randomly assign a role
	kobold.assignRandomRole()

	// Grant the skills from the role, offer them the same option
	// to choose if they role duplicates
	for _, roleSkill := range GetPotentialSkillsFromRole(kobold.Role) {
		kobold.updateSkills(roleSkill)
	}

	kobold.Bogies = []Bogie{RandomBogie()}
	kobold.Edges = []Edge{RandomEdge()}

	kobold.assignRandomArmor()
	kobold.assignRandomWeapon()
	kobold.assignRandomGear()

	return kobold
}

// Assign a skill to the kobold in question, if they have a duplicate
// skill, give them a "players choice" that they can redeem on their own
func (kobold *Kobold) updateSkills(skill Skill) {
	if kobold.Skills[skill.Name] {
		kobold.Skills[fmt.Sprintf("PLAYER CHOICE!%s", kobold.playersChoice)] = true
		kobold.playersChoice = fmt.Sprintf("!%s", kobold.playersChoice)
	} else {
		kobold.Skills[skill.Name] = true
	}
}

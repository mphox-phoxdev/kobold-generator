package kobold

import (
	"github.com/juju/errors"
	"github.com/mphox-phoxdev/kobold-generator/utils"
	"github.com/sirupsen/logrus"
)

type Skill struct {
	ID          int    `json:"-"`
	Name        string `json:"name"`
	Stat        string `json:"stat"`
	Description string `json:"description,omitempty"`
	Dangerous   bool   `json:"dangerous"`
	EveryKobold bool   `json:"everyKobold,omitempty"`
	Extra       string `json:"extra,omitempty"`
}

var idToSkillMap map[int]Skill
var roleSkillMap map[string]map[int][]Skill
var randomSkillMap map[int]map[int]Skill

func InitializeSkillMap(db Database) (err error) {
	idToSkillMap, err = db.GetSkillMap()
	if err != nil {
		err = errors.Trace(err)
		return
	}
	return err
}

func InitializeRoleSkillMap(db Database) (err error) {
	roleSkillMap, err = db.GetRoleSkillMap()
	if err != nil {
		err = errors.Trace(err)
		return
	}
	return err
}

func InitializeRandomSkillMap(db Database) (err error) {
	randomSkillMap, err = db.GetRandomSkillMap()
	if err != nil {
		err = errors.Trace(err)
		return
	}
	logrus.Info(randomSkillMap)
	return err
}

func GetSkillFromID(id int) Skill {
	return idToSkillMap[id]
}

func GetPotentialSkillsFromRole(role string) []Skill {
	return roleSkillMap[role][utils.D6()]
}

func GetRandomSkill() Skill {
	return randomSkillMap[utils.D6()][utils.D6()]
}

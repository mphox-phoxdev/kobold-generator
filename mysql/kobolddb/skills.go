package kobolddb

import (
	"database/sql"

	"github.com/juju/errors"
	"github.com/mphox-phoxdev/kobold-generator/kobold"
)

func (koboldDB *KoboldDB) GetSkillMap() (skillMap map[int]kobold.Skill, err error) {
	rows, err := koboldDB.Query(getAllSkillsQuery)
	skillMap = make(map[int]kobold.Skill)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.NotFoundf(err.Error())
			return skillMap, err
		}
		err = errors.Trace(err)
		return skillMap, err
	}

	defer rows.Close()

	for rows.Next() {
		var skill kobold.Skill

		err := rows.Scan(
			&skill.ID,
			&skill.Name,
			&skill.Stat,
			&skill.Description,
			&skill.Dangerous,
			&skill.EveryKobold,
			&skill.Extra,
		)

		if err != nil {
			err = errors.Trace(err)
			return skillMap, err
		}

		skillMap[skill.ID] = skill
	}

	err = rows.Err()
	if err != nil {
		err = errors.Trace(err)
		return skillMap, err
	}

	return skillMap, err
}

func (koboldDB *KoboldDB) GetRoleSkillMap() (roleSkillMap map[string]map[int][]kobold.Skill, err error) {
	rows, err := koboldDB.Query(getAllRoleSkillsQuery)
	roleSkillMap = make(map[string]map[int][]kobold.Skill)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.NotFoundf(err.Error())
			return roleSkillMap, err
		}
		err = errors.Trace(err)
		return roleSkillMap, err
	}

	defer rows.Close()

	for rows.Next() {
		var d6, brawnSkillID, egoSkillID, extraneousSkillID, reflexesSkillID int
		var role string
		err := rows.Scan(
			&d6,
			&role,
			&brawnSkillID,
			&egoSkillID,
			&extraneousSkillID,
			&reflexesSkillID,
		)

		if err != nil {
			err = errors.Trace(err)
			return roleSkillMap, err
		}

		if _, ok := roleSkillMap[role]; !ok {
			roleSkillMap[role] = make(map[int][]kobold.Skill)
		}

		roleSkillMap[role][d6] = []kobold.Skill{
			kobold.GetSkillFromID(brawnSkillID),
			kobold.GetSkillFromID(egoSkillID),
			kobold.GetSkillFromID(extraneousSkillID),
			kobold.GetSkillFromID(reflexesSkillID),
		}
	}

	err = rows.Err()
	if err != nil {
		err = errors.Trace(err)
		return roleSkillMap, err
	}

	return roleSkillMap, err
}

func (koboldDB *KoboldDB) GetRandomSkillMap() (randomSkillMap map[int]map[int]kobold.Skill, err error) {
	rows, err := koboldDB.Query(getAllRandomSkillsQuery)
	randomSkillMap = make(map[int]map[int]kobold.Skill)
	for i := 1; i <= 6; i++ {
		randomSkillMap[i] = make(map[int]kobold.Skill)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.NotFoundf(err.Error())
			return randomSkillMap, err
		}
		err = errors.Trace(err)
		return randomSkillMap, err
	}

	defer rows.Close()

	for rows.Next() {
		var firstD6, secondD6, skillID int
		err := rows.Scan(
			&firstD6,
			&secondD6,
			&skillID,
		)

		if err != nil {
			err = errors.Trace(err)
			return randomSkillMap, err
		}

		randomSkillMap[firstD6][secondD6] = kobold.GetSkillFromID(skillID)
	}

	err = rows.Err()
	if err != nil {
		err = errors.Trace(err)
		return randomSkillMap, err
	}

	return randomSkillMap, err
}

package kobold

import "github.com/mphox-phoxdev/kobold-generator/utils"

const ROLE_BLAZER = "blazer"
const ROLE_CASTER = "caster"
const ROLE_FRYER = "fryer"
const ROLE_SCRAPPER = "scrapper"
const ROLE_TAKER = "taker"
const ROLE_WEIRDER = "weirder"

var d6ToRole = map[int]string{
	1: ROLE_BLAZER,
	2: ROLE_CASTER,
	3: ROLE_FRYER,
	4: ROLE_SCRAPPER,
	5: ROLE_TAKER,
	6: ROLE_WEIRDER,
}

func (kobold *Kobold) assignRandomRole() {
	d6 := utils.D6()
	for d6 == 6 {
		d6 = utils.D6()
	}
	kobold.Role = d6ToRole[d6]
}

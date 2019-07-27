package kobold

// Database is the interface that covers the manipulation of kobold related objects in the database
//
// GetUserByID returns a User struct containing information about the user
// that has the given ID
//
// GetUserByEmail returns a User struct containing information about the user
// that has the given username
//
// AddUsertoDB adds the given user to the database and returns an error if that doesn't
// work
//
// GetSkillMap reads all of the entries in the skills table and returns a map of skill.IDs to Skills
//
// GetRandomSkillMap reads all of the entries in the random skills table and returns
// a 2D map of int, int, to skills skill.IDs to Skills
type Database interface {
	GetSkillMap() (skillMap map[int]Skill, err error)
	GetRoleSkillMap() (roleSkillMap map[string]map[int][]Skill, err error)
	GetRandomSkillMap() (randomSkillMap map[int]map[int]Skill, err error)
}

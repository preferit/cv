package main

type Company struct {
	Logo  string
	Phone string

	// private fields affect content
	hide bool
}

type CV struct {
	Person
	Languages
	Educations      []Education
	TechnicalSkills []TechSkill
	Experience      []Project
}

type Education struct {
	Subject string
	Grade   string
	Period
	Location string
}

type Project struct {
	Title    string
	Customer string
	Period
	Tags  []string
	Roles []string
	Short string
	More  string

	// private fields affect content
	hide      bool
	oneLiner  bool
	showShort bool
	showMore  bool
}

type Period struct {
	FromYear int
	ToYear   int
}

type TechSkill struct {
	Item  string
	E     float64
	Years int
}

type Languages []Language

type Language struct {
	Item  string
	E     float64 // experience 1.0 .. 5.0
	Level string
}

type Person struct {
	Name        string
	Image       string
	Description string
}

// ----------------------------------------

type TechSkillByName []TechSkill

func (me TechSkillByName) Len() int           { return len(me) }
func (me TechSkillByName) Swap(i, j int)      { me[j], me[i] = me[i], me[j] }
func (me TechSkillByName) Less(i, j int) bool { return me[i].Item < me[j].Item }

type TechSkillByE []TechSkill

func (me TechSkillByE) Len() int           { return len(me) }
func (me TechSkillByE) Swap(i, j int)      { me[j], me[i] = me[i], me[j] }
func (me TechSkillByE) Less(i, j int) bool { return me[i].E < me[j].E }

package model

func NewTeam() *Team {
	return &Team{Model: &Model{}}
}

type Team struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Leader   string `json:"leader" gorm:"column:leader"`
	Director string `json:"director" gorm:"column:director"`
	*Model
}

func (m *Team) TableName() string {
	return "team"
}

func (m *Team) IDField() string {
	return "id"
}
func (m *Team) NameField() string {
	return "name"
}
func (m *Team) LeaderField() string {
	return "leader"
}
func (m *Team) DirectorField() string {
	return "director"
}

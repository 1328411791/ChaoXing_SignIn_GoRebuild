package model

const (
	PeopleTableName = "test"
)

type Test struct {
	ID       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Password int64  `gorm:"column:password"`
}

func (p *Test) TableName() string {
	return PeopleTableName
}

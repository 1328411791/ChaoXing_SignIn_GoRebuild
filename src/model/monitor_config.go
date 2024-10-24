package model

// User
const (
	UserTableName = "test"
)

/*
  phone?: string;
  params?: UserParams;
  monitor?: MonitorConfig;
  mailing?: MailConfig;
  cqserver?: cqServerConfig;

*/

type MonitorConfig struct {
	ID     int64  `gorm:"column:id"`
	Phone  string `gorm:"column:phone"`
	Params int64  `gorm:"column:password"`
}

func (p *Test) UserTableName() string {
	return UserTableName
}

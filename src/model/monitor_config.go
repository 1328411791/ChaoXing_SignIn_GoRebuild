package model

// MonitorConfig
const (
	MonitorConfigTableName = "test"
)

/*
 delay: number;
  lon: number;
  lat: number;
  address: string;
*/

type MonitorConfig struct {
	ID       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Password int64  `gorm:"column:password"`
}

func (p *Test) MonitorConfigTableName() string {
	return MonitorConfigTableName
}

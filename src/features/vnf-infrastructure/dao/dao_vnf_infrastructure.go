package dao

type DaoVnfInfrastructure struct {
	Id                int    `gorm:"column:id;primaryKey"`
	Name              string `gorm:"column:name"`
	Description       string `gorm:"column:description"`
	ConfigurationFile string `gorm:"column:config_file"`
}

func (DaoVnfInfrastructure) TableName() string {
	return "agorangmanager.vnf_infrastructure"
}

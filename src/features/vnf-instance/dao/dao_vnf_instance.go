package dao

import (
	vnfInfrastructureDao "agora-vnf-manager/features/vnf-infrastructure/dao"
)

type DaoVnfInstance struct {
	Id                  int                                       `gorm:"column:id;primaryKey"`
	Name                string                                    `gorm:"column:name"`
	Description         string                                    `gorm:"column:description"`
	Type                string                                    `gorm:"column:type"`
	VnfInfraId          int                                       `gorm:"column:vnf_infra_id"`
	VnfInfrastructure   vnfInfrastructureDao.DaoVnfInfrastructure `gorm:"foreignKey:VnfInfraId;references:Id"`
	Discovered          bool                                      `gorm:"column:discovered"`
	ManagementInterface string                                    `gorm:"column:management_interface"`
	ControlInterface    string                                    `gorm:"column:control_interface"`
	Vendor              string                                    `gorm:"column:vendor"`
	Version             string                                    `gorm:"column:version"`
}

func (DaoVnfInstance) TableName() string {
	return "agorangmanager.vnf_instance"
}

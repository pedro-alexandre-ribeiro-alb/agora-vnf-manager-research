package dao

import (
	vnfInstanceDao "agora-vnf-manager/features/vnf-instance/dao"
)

type DaoVnfDeviceMapper struct {
	Id            int                           `gorm:"column:id;primaryKey"`
	DeviceId      string                        `gorm:"column:device_id"`
	VnfInstanceId int                           `gorm:"column:vnf_instance_id"`
	VnfInstance   vnfInstanceDao.DaoVnfInstance `gorm:"foreignKey:VnfInstanceId;references:Id"`
	ProxyId       int                           `gorm:"column:proxy_id"`
}

func (DaoVnfDeviceMapper) TableName() string {
	return "agorangmanager.vnf_device_mapper"
}

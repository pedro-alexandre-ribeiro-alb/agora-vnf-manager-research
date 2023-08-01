package vnf_device_mapper

import (
	app "agora-vnf-manager/core/application"
	log "agora-vnf-manager/core/log"
	vnfDeviceMapperDao "agora-vnf-manager/features/vnf-device-mapper/dao"
	types "agora-vnf-manager/features/vnf-device-mapper/types"
)

var application *app.Application

func InitVnfDeviceMapper(app *app.Application) {
	log.Info("[VnfDeviceMapperService - InitVnfDeviceMapper]: Initializing service...")
	application = app
}

func GetVnfDeviceMappers(vnf_device_mapper types.VnfDeviceMapper) ([]types.VnfDeviceMapper, error) {
	log.Info("[VnfDeviceMapperService - GetVnfDeviceMappers]")
	session := application.Db.NewSession()
	vnf_device_mapper_dao := vnfDeviceMapperDao.NewDao(session)
	session.Begin()
	vnf_device_mappers, err := vnf_device_mapper_dao.Find(vnf_device_mapper)
	if err != nil {
		log.Errorf("[VnfDeviceMapperService - GetVnfDeviceMappers]: %s", err.Error())
		return vnf_device_mappers, err
	}
	return vnf_device_mappers, nil
}

func GetVnfDeviceMapper(id int) (types.VnfDeviceMapper, bool, error) {
	log.Infof("[VnfDeviceMapperService - GetVnfDeviceMapper]: {id: %d}", id)
	session := application.Db.NewSession()
	vnf_device_mapper_dao := vnfDeviceMapperDao.NewDao(session)
	session.Begin()
	vnf_device_mapper, found, err := vnf_device_mapper_dao.Get(id)
	if err != nil || !found {
		if err != nil {
			log.Errorf("[VnfDeviceMapperService - GetVnfDeviceMapper]: %s", err.Error())
			return vnf_device_mapper, false, err
		}
		return vnf_device_mapper, false, err
	}
	return vnf_device_mapper, true, nil
}

func DeleteVnfDeviceMapper(id int) (bool, error) {
	log.Infof("[VnfDeviceMapperService - DeleteVnfDeviceMapper]: {id: %d}", id)
	session := application.Db.NewSession()
	vnf_device_mapper_dao := vnfDeviceMapperDao.NewDao(session)
	session.Begin()
	success, err := vnf_device_mapper_dao.Delete(id)
	if err != nil || !success {
		if err != nil {
			log.Errorf("[VnfDeviceMapperService - DeleteVnfDeviceMapper]: %s", err.Error())
			session.Rollback()
			return false, err
		}
		return false, err
	}
	session.Commit()
	return true, nil
}

func ExistsVnfDeviceMapper(id int) (bool, error) {
	log.Infof("[VnfDeviceMapperService - ExistsVnfDeviceMapper]: {id: %d}", id)
	session := application.Db.NewSession()
	vnf_device_mapper_dao := vnfDeviceMapperDao.NewDao(session)
	session.Begin()
	vnf_device_mapper_exists, err := vnf_device_mapper_dao.Exists(id)
	if err != nil {
		log.Errorf("[VnfDeviceMapperService - ExistsVnfDeviceMapper]: %s", err.Error())
		return vnf_device_mapper_exists, err
	}
	return vnf_device_mapper_exists, nil
}

func UpdateVnfDeviceMapper(vnf_device_mapper types.VnfDeviceMapper, id int) (types.VnfDeviceMapper, error) {
	log.Infof("[VnfDeviceMapperService - UpdateVnfDeviceMapper]: {vnf_device_mapper: %+v, id: %d}", vnf_device_mapper, id)
	session := application.Db.NewSession()
	vnf_device_mapper_dao := vnfDeviceMapperDao.NewDao(session)
	session.Begin()
	updated_vnf_device_mapper, err := vnf_device_mapper_dao.Update(vnf_device_mapper, id)
	if err != nil {
		log.Errorf("[VnfDeviceMapperService - UpdateVnfDeviceMapper]: %s", err.Error())
		session.Rollback()
		return updated_vnf_device_mapper, err
	}
	session.Commit()
	return updated_vnf_device_mapper, nil
}

func CreateVnfDeviceMapper(vnf_device_mapper types.VnfDeviceMapper) (types.VnfDeviceMapper, error) {
	log.Infof("[VnfDeviceMapperService - CreateVnfDeviceMapper]: {vnf_device_mapper: %+v}", vnf_device_mapper)
	session := application.Db.NewSession()
	vnf_device_mapper_dao := vnfDeviceMapperDao.NewDao(session)
	session.Begin()
	created_vnf_device_mapper, err := vnf_device_mapper_dao.Insert(vnf_device_mapper)
	if err != nil {
		log.Errorf("[VnfDeviceMapperService - CreateVnfDeviceMapper]: %s", err.Error())
		session.Rollback()
		return created_vnf_device_mapper, err
	}
	session.Commit()
	return created_vnf_device_mapper, nil
}

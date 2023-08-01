package vnf_instance

import (
	app "agora-vnf-manager/core/application"
	log "agora-vnf-manager/core/log"
	vnfInstanceDao "agora-vnf-manager/features/vnf-instance/dao"
	types "agora-vnf-manager/features/vnf-instance/types"
)

var application *app.Application

func InitVnfInstanceService(app *app.Application) {
	log.Info("[VnfInstanceService - InitVnfInstanceService]: Initializing service...")
	application = app
}

func GetVnfInstances(vnf_instance types.VnfInstance) ([]types.VnfInstance, error) {
	log.Info("[VnfInstanceService - GetVnfInstances]")
	session := application.Db.NewSession()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	session.Begin()
	vnf_instances, err := vnf_instance_dao.Find(vnf_instance)
	if err != nil {
		log.Errorf("[VnfInstanceService - GetVnfInstances]: %s", err.Error())
		return vnf_instances, err
	}
	return vnf_instances, nil
}

func GetVnfInstance(id int) (types.VnfInstance, bool, error) {
	log.Infof("[VnfInstanceService - GetVnfInstance]: {id: %d}", id)
	session := application.Db.NewSession()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	session.Begin()
	vnf_instance, found, err := vnf_instance_dao.Get(id)
	if err != nil || !found {
		if err != nil {
			log.Errorf("[VnfInstanceService - GetVnfInstance]: %s", err.Error())
			return vnf_instance, false, err
		}
		return vnf_instance, false, err
	}
	return vnf_instance, true, nil
}

func DeleteVnfInstance(id int) (bool, error) {
	log.Infof("[VnfInstanceService - DeleteVnfInstance]: {id: %d}", id)
	session := application.Db.NewSession()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	session.Begin()
	success, err := vnf_instance_dao.Delete(id)
	if err != nil || !success {
		if err != nil {
			log.Errorf("[VnfInstanceService - DeleteVnfInstance]: %s", err.Error())
			session.Rollback()
			return false, err
		}
		return false, err
	}
	session.Commit()
	return true, nil
}

func ExistsVnfInstance(id int) (bool, error) {
	log.Infof("[VnfInstanceService - ExistsVnfInstance]: {id: %d}", id)
	session := application.Db.NewSession()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	session.Begin()
	vnf_instance_exists, err := vnf_instance_dao.Exists(id)
	if err != nil {
		log.Errorf("[VnfInstanceService - ExistsVnfInstance]: %s", err.Error())
		return vnf_instance_exists, err
	}
	return vnf_instance_exists, nil
}

func UpdateVnfInstance(vnf_instance types.VnfInstance, id int) (types.VnfInstance, error) {
	log.Infof("[VnfInstanceService - UpdateVnfInstance]: {vnf_instance: %+v, id: %d}", vnf_instance, id)
	session := application.Db.NewSession()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	session.Begin()
	updated_vnf_instance, err := vnf_instance_dao.Update(vnf_instance, id)
	if err != nil {
		log.Errorf("[VnfInstanceService - UpdateVnfInstance]: %s", err.Error())
		session.Rollback()
		return updated_vnf_instance, err
	}
	session.Commit()
	return updated_vnf_instance, nil
}

func CreateVnfInstance(vnf_instance types.VnfInstance) (types.VnfInstance, error) {
	log.Infof("[VnfInstanceService - CreateVnfInstance]: {vnf_instance: %+v}", vnf_instance)
	session := application.Db.NewSession()
	vnf_instance_dao := vnfInstanceDao.NewDao(session)
	session.Begin()
	created_vnf_instance, err := vnf_instance_dao.Insert(vnf_instance)
	if err != nil {
		log.Errorf("[VnfInstanceService - CreateVnfInstance]: %s", err.Error())
		session.Rollback()
		return created_vnf_instance, err
	}
	session.Commit()
	return created_vnf_instance, nil
}

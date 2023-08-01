package vnf_infrastructure

import (
	app "agora-vnf-manager/core/application"
	log "agora-vnf-manager/core/log"
	vnfInfrastructureDao "agora-vnf-manager/features/vnf-infrastructure/dao"
	types "agora-vnf-manager/features/vnf-infrastructure/types"
)

var application *app.Application

func InitVnfInfrastructure(app *app.Application) {
	log.Info("[VnfInfrastructureService - InitVnfInfrastructure]: Initializing service...")
	application = app
}

func GetVnfInfrastructures(vnf_infrastructure types.VnfInfrastructure) ([]types.VnfInfrastructure, error) {
	log.Info("[VnfInfrastructureService - GetVnfInfrastructures]")
	session := application.Db.NewSession()
	vnf_infrastructure_dao := vnfInfrastructureDao.NewDao(session)
	session.Begin()
	vnf_infrastructures, err := vnf_infrastructure_dao.Find(vnf_infrastructure)
	if err != nil {
		log.Errorf("[VnfInfrastructureService - GetVnfInfrastructures]: %s", err.Error())
		return vnf_infrastructures, err
	}
	return vnf_infrastructures, nil
}

func GetVnfInfrastructure(id int) (types.VnfInfrastructure, bool, error) {
	log.Infof("[VnfInfrastructureService - GetVnfInfrastructure]: {id: %d}", id)
	session := application.Db.NewSession()
	vnf_infrastructure_dao := vnfInfrastructureDao.NewDao(session)
	session.Begin()
	vnf_infrastructure, found, err := vnf_infrastructure_dao.Get(id)
	if err != nil || !found {
		if err != nil {
			log.Errorf("[VnfInfrastructureService - GetVnfInfrastructure]: %s", err.Error())
			return vnf_infrastructure, false, err
		}
		return vnf_infrastructure, false, err
	}
	return vnf_infrastructure, true, nil
}

func DeleteVnfInfrastructure(id int) (bool, error) {
	log.Infof("[VnfInfrastructureService - DeleteVnfInfrastructure]: {id: %d}", id)
	session := application.Db.NewSession()
	vnf_infrastructure_dao := vnfInfrastructureDao.NewDao(session)
	session.Begin()
	success, err := vnf_infrastructure_dao.Delete(id)
	if err != nil || !success {
		if err != nil {
			log.Errorf("[VnfInfrastructureService - DeleteVnfInfrastructure]: %s", err.Error())
			session.Rollback()
			return false, err
		}
		return false, err
	}
	session.Commit()
	return true, nil
}

func ExistsVnfInfrastructure(id int) (bool, error) {
	log.Infof("[VnfInfrastructureService - ExistsVnfInfrastructure]: {id: %d}", id)
	session := application.Db.NewSession()
	vnf_infrastructure_dao := vnfInfrastructureDao.NewDao(session)
	session.Begin()
	vnf_infrastructure_exists, err := vnf_infrastructure_dao.Exists(id)
	if err != nil {
		log.Errorf("[VnfInfrastructureService - ExistsVnfInfrastructure]: %s", err.Error())
		return vnf_infrastructure_exists, err
	}
	return vnf_infrastructure_exists, nil
}

func UpdateVnfInfrastructure(vnf_infrastructure types.VnfInfrastructure, id int) (types.VnfInfrastructure, error) {
	log.Infof("[VnfInfrastructureService - UpdateVnfInfrastructure]: {vnf_infrastructure: %+v, id: %d}", vnf_infrastructure, id)
	session := application.Db.NewSession()
	vnf_infrastructure_dao := vnfInfrastructureDao.NewDao(session)
	session.Begin()
	update_vnf_infrastructure, err := vnf_infrastructure_dao.Update(vnf_infrastructure, id)
	if err != nil {
		log.Errorf("[VnfInfrastructureService - UpdateVnfInfrastructure]: %s", err.Error())
		session.Rollback()
		return update_vnf_infrastructure, err
	}
	session.Commit()
	return update_vnf_infrastructure, nil
}

func CreateVnfInfrastructure(vnf_infrastructure types.VnfInfrastructure) (types.VnfInfrastructure, error) {
	log.Infof("[VnfInfrastructureService - CreateVnfInfrastructure]: {vnf_infrastructure: %+v}", vnf_infrastructure)
	session := application.Db.NewSession()
	vnf_infrastructure_dao := vnfInfrastructureDao.NewDao(session)
	session.Begin()
	created_vnf_infrastructure, err := vnf_infrastructure_dao.Insert(vnf_infrastructure)
	if err != nil {
		log.Errorf("[VnfInfrastructureService - CreateVnfInfrastructure]: %s", err.Error())
		session.Rollback()
		return created_vnf_infrastructure, err
	}
	session.Commit()
	return created_vnf_infrastructure, nil
}

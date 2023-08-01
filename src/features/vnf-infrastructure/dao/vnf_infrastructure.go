package dao

import (
	helpers "agora-vnf-manager/core/helpers"
	log "agora-vnf-manager/core/log"
	db "agora-vnf-manager/db"
	types "agora-vnf-manager/features/vnf-infrastructure/types"
)

type VnfInfrastructureDao struct {
	session db.IDbSession
}

func NewDao(session db.IDbSession) *VnfInfrastructureDao {
	return &VnfInfrastructureDao{session: session}
}

func (vnf_infrastructure_dao VnfInfrastructureDao) Get(id int) (types.VnfInfrastructure, bool, error) {
	dao_vnf_infrastructure_search := DaoVnfInfrastructure{Id: id}
	dao_vnf_infrastructure_data, has, err := helpers.GetEntry(vnf_infrastructure_dao.session, dao_vnf_infrastructure_search)
	if err != nil || !has {
		if err != nil {
			log.Errorf("[VnfInfrastructureDao - Get]: %s", err.Error())
		}
		return types.VnfInfrastructure{}, false, err
	}
	retrieved_vnf_infrastructure := MapDaoVnfInfrastructureToVnfInfrastructure(dao_vnf_infrastructure_data)
	return retrieved_vnf_infrastructure, true, nil
}

func (vnf_infrastructure_dao VnfInfrastructureDao) Exists(id int) (bool, error) {
	dao_vnf_infrastructure_search := DaoVnfInfrastructure{Id: id}
	dao_vnf_infrastructure, err := helpers.ExistsEntry(vnf_infrastructure_dao.session, dao_vnf_infrastructure_search)
	if err != nil {
		log.Errorf("[VnfInfrastructureDao - Exists]: %s", err.Error())
		return false, err
	}
	return dao_vnf_infrastructure, nil
}

func (vnf_infrastructure_dao VnfInfrastructureDao) Find(vnf_infrastructure types.VnfInfrastructure) ([]types.VnfInfrastructure, error) {
	vnf_infrastructures := []types.VnfInfrastructure{}
	dao_vnf_infrastructures_search := MapVnfInfrastructureToDaoVnfInfrastructureFind(vnf_infrastructure)
	dao_vnf_infrastructures_data, err := helpers.FindEntry(vnf_infrastructure_dao.session, dao_vnf_infrastructures_search)
	if err != nil {
		log.Errorf("[VnfInfrastructureDao - Find]: %s", err.Error())
		return vnf_infrastructures, err
	}
	for _, dao_vnf_infrastructure_data := range dao_vnf_infrastructures_data {
		vnf_infrastructure := MapDaoVnfInfrastructureToVnfInfrastructure(dao_vnf_infrastructure_data)
		vnf_infrastructures = append(vnf_infrastructures, vnf_infrastructure)
	}
	return vnf_infrastructures, nil
}

func (vnf_infrastructure_dao VnfInfrastructureDao) Delete(id int) (bool, error) {
	dao_vnf_infrastructure_delete := DaoVnfInfrastructure{Id: id}
	_, err := helpers.DeleteEntry(vnf_infrastructure_dao.session, dao_vnf_infrastructure_delete)
	if err != nil {
		log.Errorf("[VnfInfrastructureDao - Delete]: %s", err.Error())
		return false, err
	}
	return true, nil
}

func (vnf_infrastructure_dao VnfInfrastructureDao) Insert(vnf_infrastructure types.VnfInfrastructure) (types.VnfInfrastructure, error) {
	dao_vnf_infrastructure_create, err := MapVnfInfrastructureToDaoVnfInfrastructure(vnf_infrastructure, vnf_infrastructure_dao)
	if err != nil {
		log.Errorf("[VnfInfrastructureDao - Insert]: %s", err.Error())
		return types.VnfInfrastructure{}, err
	}
	dao_vnf_infrastructure_data, err := helpers.CreateEntry(vnf_infrastructure_dao.session, dao_vnf_infrastructure_create)
	if err != nil {
		log.Errorf("[VnfInfrastructureDao - Insert]: %s", err.Error())
		return types.VnfInfrastructure{}, err
	}
	created_vnf_infrastructure := MapDaoVnfInfrastructureToVnfInfrastructure(dao_vnf_infrastructure_data)
	return created_vnf_infrastructure, nil
}

func (vnf_infrastructure_dao VnfInfrastructureDao) Update(vnf_infrastructure types.VnfInfrastructure, id int) (types.VnfInfrastructure, error) {
	dao_vnf_infrastructure_replace, err := MapVnfInfrastructureToDaoVnfInfrastructureUpdate(vnf_infrastructure, id, vnf_infrastructure_dao)
	if err != nil {
		log.Errorf("[VnfInfrastructureDao - Update]: %s", err.Error())
		return types.VnfInfrastructure{}, err
	}
	dao_vnf_infrastructure_data, err := helpers.ReplaceEntry(vnf_infrastructure_dao.session, dao_vnf_infrastructure_replace)
	if err != nil {
		log.Errorf("[VnfInfrastructureDao - Update]: %s", err.Error())
		return types.VnfInfrastructure{}, err
	}
	updated_vnf_infrastructure := MapDaoVnfInfrastructureToVnfInfrastructure(dao_vnf_infrastructure_data)
	return updated_vnf_infrastructure, nil
}

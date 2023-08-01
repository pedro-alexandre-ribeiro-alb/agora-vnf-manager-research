package dao

import (
	helpers "agora-vnf-manager/core/helpers"
	log "agora-vnf-manager/core/log"
	db "agora-vnf-manager/db"
	types "agora-vnf-manager/features/vnf-instance/types"
)

type VnfInstanceDao struct {
	session db.IDbSession
}

func NewDao(session db.IDbSession) *VnfInstanceDao {
	return &VnfInstanceDao{session: session}
}

func (vnf_instance_dao VnfInstanceDao) Get(id int) (types.VnfInstance, bool, error) {
	dao_vnf_instance_search := DaoVnfInstance{Id: id}
	dao_vnf_instance_data, has, err := helpers.GetEntry(vnf_instance_dao.session, dao_vnf_instance_search)
	if err != nil || !has {
		if err != nil {
			log.Errorf("[VnfInstanceDao - Get]: %s", err.Error())
		}
		return types.VnfInstance{}, false, err
	}
	retrieved_vnf_instance := MapDaoVnfInstanceToVnfInstance(dao_vnf_instance_data)
	return retrieved_vnf_instance, true, nil
}

func (vnf_instance_dao VnfInstanceDao) Exists(id int) (bool, error) {
	dao_vnf_instance_search := DaoVnfInstance{Id: id}
	dao_vnf_instance, err := helpers.ExistsEntry(vnf_instance_dao.session, dao_vnf_instance_search)
	if err != nil {
		log.Errorf("[VnfInstanceDao - Exists]: %s", err.Error())
		return false, err
	}
	return dao_vnf_instance, nil
}

func (vnf_instance_dao VnfInstanceDao) Find(vnf_instance types.VnfInstance) ([]types.VnfInstance, error) {
	vnf_instances := []types.VnfInstance{}
	dao_vnf_instances_search := MapVnfInstanceToDaoVnfInstanceFind(vnf_instance)
	dao_vnf_instances_data, err := helpers.FindEntry(vnf_instance_dao.session, dao_vnf_instances_search)
	if err != nil {
		log.Errorf("[VnfInstanceDao - Find]: %s", err.Error())
		return vnf_instances, err
	}
	for _, dao_vnf_instance_data := range dao_vnf_instances_data {
		vnf_instance := MapDaoVnfInstanceToVnfInstance(dao_vnf_instance_data)
		vnf_instances = append(vnf_instances, vnf_instance)
	}
	return vnf_instances, nil
}

func (vnf_instance_dao VnfInstanceDao) Delete(id int) (bool, error) {
	dao_vnf_instance_delete := DaoVnfInstance{Id: id}
	_, err := helpers.DeleteEntry(vnf_instance_dao.session, dao_vnf_instance_delete)
	if err != nil {
		log.Errorf("[VnfInstanceDao - Delete]: %s", err.Error())
		return false, err
	}
	return true, nil
}

func (vnf_instance_dao VnfInstanceDao) Insert(vnf_instance types.VnfInstance) (types.VnfInstance, error) {
	dao_vnf_instance_create, err := MapVnfInstanceToDaoVnfInstance(vnf_instance, vnf_instance_dao)
	if err != nil {
		log.Errorf("[VnfInstanceDao - Insert]: %s", err.Error())
		return types.VnfInstance{}, err
	}
	dao_vnf_instance_data, err := helpers.CreateEntry(vnf_instance_dao.session, dao_vnf_instance_create)
	if err != nil {
		log.Errorf("[VnfInstanceDao - Insert]: %s", err.Error())
		return types.VnfInstance{}, err
	}
	created_vnf_instance := MapDaoVnfInstanceToVnfInstance(dao_vnf_instance_data)
	return created_vnf_instance, nil
}

func (vnf_instance_dao VnfInstanceDao) Update(vnf_instance types.VnfInstance, id int) (types.VnfInstance, error) {
	dao_vnf_instance_replace, err := MapVnfInstanceToDaoVnfInstanceUpdate(vnf_instance, id, vnf_instance_dao)
	if err != nil {
		log.Errorf("[VnfInstanceDao - Update]: %s", err.Error())
		return types.VnfInstance{}, err
	}
	dao_vnf_instance_data, err := helpers.ReplaceEntry(vnf_instance_dao.session, dao_vnf_instance_replace)
	if err != nil {
		log.Errorf("[VnfInstanceDao - Update]: %s", err.Error())
		return types.VnfInstance{}, err
	}
	updated_vnf_instance := MapDaoVnfInstanceToVnfInstance(dao_vnf_instance_data)
	return updated_vnf_instance, nil
}

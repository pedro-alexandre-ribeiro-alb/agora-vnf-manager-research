package dao

import (
	helpers "agora-vnf-manager/core/helpers"
	log "agora-vnf-manager/core/log"
	db "agora-vnf-manager/db"
	types "agora-vnf-manager/features/vnf-device-mapper/types"
)

type VnfDeviceMapperDao struct {
	session db.IDbSession
}

func NewDao(session db.IDbSession) *VnfDeviceMapperDao {
	return &VnfDeviceMapperDao{session: session}
}

func (vnf_device_mapper_dao VnfDeviceMapperDao) Get(id int) (types.VnfDeviceMapper, bool, error) {
	dao_vnf_device_mapper_search := DaoVnfDeviceMapper{Id: id}
	dao_vnf_device_mapper_data, has, err := helpers.GetEntry(vnf_device_mapper_dao.session, dao_vnf_device_mapper_search)
	if err != nil || !has {
		if err != nil {
			log.Errorf("[VnfDeviceMapperDao - Get]: %s", err.Error())
		}
		return types.VnfDeviceMapper{}, false, err
	}
	retrieved_vnf_device_mapper := MapDaoVnfDeviceMapperToVnfDeviceMapper(dao_vnf_device_mapper_data)
	return retrieved_vnf_device_mapper, true, nil
}

func (vnf_device_mapper_dao VnfDeviceMapperDao) Exists(id int) (bool, error) {
	dao_vnf_device_mapper_search := DaoVnfDeviceMapper{Id: id}
	dao_vnf_device_mapper, err := helpers.ExistsEntry(vnf_device_mapper_dao.session, dao_vnf_device_mapper_search)
	if err != nil {
		log.Errorf("[VnfDeviceMapperDao - Exists]: %s", err.Error())
		return false, err
	}
	return dao_vnf_device_mapper, nil
}

func (vnf_device_mapper_dao VnfDeviceMapperDao) Find(vnf_device_mapper types.VnfDeviceMapper) ([]types.VnfDeviceMapper, error) {
	vnf_device_mappers := []types.VnfDeviceMapper{}
	dao_vnf_device_mappers_search := MapVnfDeviceMapperToDaoVnfDeviceMapperFind(vnf_device_mapper)
	dao_vnf_device_mappers_data, err := helpers.FindEntry(vnf_device_mapper_dao.session, dao_vnf_device_mappers_search)
	if err != nil {
		log.Errorf("[VnfDeviceMapperDao - Find]: %s", err.Error())
		return vnf_device_mappers, err
	}
	for _, dao_vnf_device_mapper_data := range dao_vnf_device_mappers_data {
		vnf_device_mapper := MapDaoVnfDeviceMapperToVnfDeviceMapper(dao_vnf_device_mapper_data)
		vnf_device_mappers = append(vnf_device_mappers, vnf_device_mapper)
	}
	return vnf_device_mappers, nil
}

func (vnf_device_mapper_dao VnfDeviceMapperDao) Delete(id int) (bool, error) {
	dao_vnf_device_mapper_delete := DaoVnfDeviceMapper{Id: id}
	_, err := helpers.DeleteEntry(vnf_device_mapper_dao.session, dao_vnf_device_mapper_delete)
	if err != nil {
		log.Errorf("[VnfDeviceMapperDao - Delete]: %s", err.Error())
		return false, err
	}
	return true, nil
}

func (vnf_device_mapper_dao VnfDeviceMapperDao) Insert(vnf_device_mapper types.VnfDeviceMapper) (types.VnfDeviceMapper, error) {
	dao_vnf_device_mapper_create, err := MapVnfDeviceMapperToDaoVnfDeviceMapper(vnf_device_mapper, vnf_device_mapper_dao)
	if err != nil {
		log.Errorf("[VnfDeviceMapperDao - Insert]: %s", err.Error())
		return types.VnfDeviceMapper{}, err
	}
	dao_vnf_device_mapper_data, err := helpers.CreateEntry(vnf_device_mapper_dao.session, dao_vnf_device_mapper_create)
	if err != nil {
		log.Errorf("[VnfDeviceMapperDao - Insert]: %s", err.Error())
		return types.VnfDeviceMapper{}, err
	}
	created_vnf_device_mapper := MapDaoVnfDeviceMapperToVnfDeviceMapper(dao_vnf_device_mapper_data)
	return created_vnf_device_mapper, nil
}

func (vnf_device_mapper_dao VnfDeviceMapperDao) Update(vnf_device_mapper types.VnfDeviceMapper, id int) (types.VnfDeviceMapper, error) {
	dao_vnf_device_mapper_replace, err := MapVnfDeviceMapperToDaoVnfDeviceMapperUpdate(vnf_device_mapper, id, vnf_device_mapper_dao)
	if err != nil {
		log.Errorf("[VnfDeviceMapperDao - Update]: %s", err.Error())
		return types.VnfDeviceMapper{}, err
	}
	dao_vnf_device_mapper_data, err := helpers.ReplaceEntry(vnf_device_mapper_dao.session, dao_vnf_device_mapper_replace)
	if err != nil {
		log.Errorf("[VnfDeviceMapperDao - Update]: %s", err.Error())
		return types.VnfDeviceMapper{}, err
	}
	updated_vnf_device_mapper := MapDaoVnfDeviceMapperToVnfDeviceMapper(dao_vnf_device_mapper_data)
	return updated_vnf_device_mapper, nil
}

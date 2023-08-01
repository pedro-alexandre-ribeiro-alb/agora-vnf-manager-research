package dao

import (
	helpers "agora-vnf-manager/core/helpers"
	log "agora-vnf-manager/core/log"
	utils "agora-vnf-manager/core/utils"
	types "agora-vnf-manager/features/vnf-device-mapper/types"
	"fmt"
)

func MapDaoVnfDeviceMapperToVnfDeviceMapper(dao_vnf_device_mapper DaoVnfDeviceMapper) types.VnfDeviceMapper {
	vnf_device_mapper := types.VnfDeviceMapper{}
	vnf_device_mapper.Id = dao_vnf_device_mapper.Id
	vnf_device_mapper.DeviceId.Scan(dao_vnf_device_mapper.DeviceId)
	vnf_device_mapper.VnfInstanceId.Scan(dao_vnf_device_mapper.VnfInstanceId)
	vnf_device_mapper.ProxyId.Scan(dao_vnf_device_mapper.ProxyId)
	return vnf_device_mapper
}

func MapVnfDeviceMapperToDaoVnfDeviceMapper(vnf_device_mapper types.VnfDeviceMapper, vnf_device_mapper_dao VnfDeviceMapperDao) (DaoVnfDeviceMapper, error) {
	dao_vnf_device_mapper := DaoVnfDeviceMapper{}
	dao_vnf_device_mapper_id, err := getDaoVnfDeviceMapperId(vnf_device_mapper_dao)
	if err != nil {
		log.Errorf("[MapVnfDeviceMapperToDaoVnfDeviceMapper]: %s", err.Error())
		return dao_vnf_device_mapper, err
	}
	dao_vnf_device_mapper.Id = dao_vnf_device_mapper_id
	dao_vnf_device_mapper.DeviceId = string(utils.First(vnf_device_mapper.DeviceId.Get()))
	dao_vnf_device_mapper.VnfInstanceId = int(utils.First(vnf_device_mapper.VnfInstanceId.Get()))
	dao_vnf_device_mapper.ProxyId = int(utils.First(vnf_device_mapper.ProxyId.Get()))
	return dao_vnf_device_mapper, nil
}

func MapVnfDeviceMapperToDaoVnfDeviceMapperFind(vnf_device_mapper types.VnfDeviceMapper) DaoVnfDeviceMapper {
	dao_vnf_device_mapper := DaoVnfDeviceMapper{}
	if vnf_device_mapper.DeviceId.Present() {
		dao_vnf_device_mapper.DeviceId = string(utils.First(vnf_device_mapper.DeviceId.Get()))
	}
	if vnf_device_mapper.VnfInstanceId.Present() {
		dao_vnf_device_mapper.VnfInstanceId = int(utils.First(vnf_device_mapper.VnfInstanceId.Get()))
	}
	if vnf_device_mapper.ProxyId.Present() {
		dao_vnf_device_mapper.ProxyId = int(utils.First(vnf_device_mapper.ProxyId.Get()))
	}
	return dao_vnf_device_mapper
}

func MapVnfDeviceMapperToDaoVnfDeviceMapperUpdate(vnf_device_mapper types.VnfDeviceMapper, id int, vnf_device_mapper_dao VnfDeviceMapperDao) (DaoVnfDeviceMapper, error) {
	dao_vnf_device_mapper_search := DaoVnfDeviceMapper{Id: id}
	dao_vnf_device_mapper_data, has, err := helpers.GetEntry(vnf_device_mapper_dao.session, dao_vnf_device_mapper_search)
	if err != nil || !has {
		if err != nil {
			log.Errorf("[MapVnfDeviceMapperToDaoVnfDeviceMapperUpdate]: %s", err.Error())
			return DaoVnfDeviceMapper{}, err
		}
		return DaoVnfDeviceMapper{}, fmt.Errorf("Could not find referenced vnf_device_mapper with id: %s", err.Error())
	}
	if vnf_device_mapper.DeviceId.Present() {
		dao_vnf_device_mapper_data.DeviceId = string(utils.First(vnf_device_mapper.DeviceId.Get()))
	}
	if vnf_device_mapper.VnfInstanceId.Present() {
		dao_vnf_device_mapper_data.VnfInstanceId = int(utils.First(vnf_device_mapper.VnfInstanceId.Get()))
	}
	if vnf_device_mapper.ProxyId.Present() {
		dao_vnf_device_mapper_data.ProxyId = int(utils.First(vnf_device_mapper.ProxyId.Get()))
	}
	return dao_vnf_device_mapper_data, nil
}

func getDaoVnfDeviceMapperId(vnf_device_mapper_dao VnfDeviceMapperDao) (int, error) {
	var vnf_device_mapper_dao_id int
	err := vnf_device_mapper_dao.session.Query("SELECT nextval('agorangmanager.vnf_device_mapper_id_seq');", &vnf_device_mapper_dao_id)
	if err != nil {
		log.Errorf("[VnfDeviceMapperDao - getDaoVnfDeviceMapperId]: %s", err.Error())
		return -1, err
	}
	return vnf_device_mapper_dao_id, nil
}

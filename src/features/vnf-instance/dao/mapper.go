package dao

import (
	helpers "agora-vnf-manager/core/helpers"
	log "agora-vnf-manager/core/log"
	utils "agora-vnf-manager/core/utils"
	types "agora-vnf-manager/features/vnf-instance/types"
	"fmt"
)

func MapDaoVnfInstanceToVnfInstance(dao_vnf_instance DaoVnfInstance) types.VnfInstance {
	vnf_instance := types.VnfInstance{}
	vnf_instance.Id = dao_vnf_instance.Id
	vnf_instance.Name.Scan(dao_vnf_instance.Name)
	vnf_instance.Description.Scan(dao_vnf_instance.Description)
	vnf_instance.Type.Scan(dao_vnf_instance.Type)
	vnf_instance.VnfInfraId.Scan(dao_vnf_instance.VnfInfraId)
	vnf_instance.Discovered.Scan(dao_vnf_instance.Discovered)
	vnf_instance.ManagementInterface.Scan(dao_vnf_instance.ManagementInterface)
	vnf_instance.ControlInterface.Scan(dao_vnf_instance.ControlInterface)
	vnf_instance.Vendor.Scan(dao_vnf_instance.Vendor)
	vnf_instance.Version.Scan(dao_vnf_instance.Version)
	return vnf_instance
}

func MapVnfInstanceToDaoVnfInstance(vnf_instance types.VnfInstance, vnf_instance_dao VnfInstanceDao) (DaoVnfInstance, error) {
	dao_vnf_instance := DaoVnfInstance{}
	dao_vnf_instance_id, err := getDaoVnfInstanceId(vnf_instance_dao)
	if err != nil {
		log.Errorf("[MapVnfInstanceToDaoVnfInstance]: %s", err.Error())
		return dao_vnf_instance, err
	}
	dao_vnf_instance.Id = dao_vnf_instance_id
	dao_vnf_instance.Name = string(utils.First(vnf_instance.Name.Get()))
	dao_vnf_instance.Description = string(utils.First(vnf_instance.Description.Get()))
	dao_vnf_instance.Type = string(utils.First(vnf_instance.Type.Get()))
	dao_vnf_instance.VnfInfraId = int(utils.First(vnf_instance.VnfInfraId.Get()))
	dao_vnf_instance.Discovered = bool(utils.First(vnf_instance.Discovered.Get()))
	dao_vnf_instance.ManagementInterface = string(utils.First(vnf_instance.ManagementInterface.Get()))
	dao_vnf_instance.ControlInterface = string(utils.First(vnf_instance.ControlInterface.Get()))
	dao_vnf_instance.Vendor = string(utils.First(vnf_instance.Vendor.Get()))
	dao_vnf_instance.Version = string(utils.First(vnf_instance.Version.Get()))
	return dao_vnf_instance, nil
}

func MapVnfInstanceToDaoVnfInstanceFind(vnf_instance types.VnfInstance) DaoVnfInstance {
	dao_vnf_instance := DaoVnfInstance{}
	if vnf_instance.Name.Present() {
		dao_vnf_instance.Name = string(utils.First(vnf_instance.Name.Get()))
	}
	if vnf_instance.Description.Present() {
		dao_vnf_instance.Description = string(utils.First(vnf_instance.Description.Get()))
	}
	if vnf_instance.Type.Present() {
		dao_vnf_instance.Type = string(utils.First(vnf_instance.Type.Get()))
	}
	if vnf_instance.VnfInfraId.Present() {
		dao_vnf_instance.VnfInfraId = int(utils.First(vnf_instance.VnfInfraId.Get()))
	}
	if vnf_instance.Discovered.Present() {
		dao_vnf_instance.Discovered = bool(utils.First(vnf_instance.Discovered.Get()))
	}
	if vnf_instance.ManagementInterface.Present() {
		dao_vnf_instance.ManagementInterface = string(utils.First(vnf_instance.ManagementInterface.Get()))
	}
	if vnf_instance.ControlInterface.Present() {
		dao_vnf_instance.ControlInterface = string(utils.First(vnf_instance.ControlInterface.Get()))
	}
	if vnf_instance.Vendor.Present() {
		dao_vnf_instance.Vendor = string(utils.First(vnf_instance.Vendor.Get()))
	}
	if vnf_instance.Version.Present() {
		dao_vnf_instance.Version = string(utils.First(vnf_instance.Version.Get()))
	}
	return dao_vnf_instance
}

func MapVnfInstanceToDaoVnfInstanceUpdate(vnf_instance types.VnfInstance, id int, vnf_instance_dao VnfInstanceDao) (DaoVnfInstance, error) {
	dao_vnf_instance_search := DaoVnfInstance{Id: id}
	dao_vnf_instance_data, has, err := helpers.GetEntry(vnf_instance_dao.session, dao_vnf_instance_search)
	if err != nil || !has {
		if err != nil {
			log.Errorf("[MapVnfInstanceToDaoVnfInstanceUpdate]: %s", err.Error())
			return DaoVnfInstance{}, err
		}
		return DaoVnfInstance{}, fmt.Errorf("Could not find referenced vnf_instance with id: %d", id)
	}
	if vnf_instance.Name.Present() {
		dao_vnf_instance_data.Name = string(utils.First(vnf_instance.Name.Get()))
	}
	if vnf_instance.Description.Present() {
		dao_vnf_instance_data.Description = string(utils.First(vnf_instance.Description.Get()))
	}
	if vnf_instance.Type.Present() {
		dao_vnf_instance_data.Type = string(utils.First(vnf_instance.Type.Get()))
	}
	if vnf_instance.VnfInfraId.Present() {
		dao_vnf_instance_data.VnfInfraId = int(utils.First(vnf_instance.VnfInfraId.Get()))
	}
	if vnf_instance.Discovered.Present() {
		dao_vnf_instance_data.Discovered = bool(utils.First(vnf_instance.Discovered.Get()))
	}
	if vnf_instance.ManagementInterface.Present() {
		dao_vnf_instance_data.ManagementInterface = string(utils.First(vnf_instance.ManagementInterface.Get()))
	}
	if vnf_instance.ControlInterface.Present() {
		dao_vnf_instance_data.ControlInterface = string(utils.First(vnf_instance.ControlInterface.Get()))
	}
	if vnf_instance.Vendor.Present() {
		dao_vnf_instance_data.Vendor = string(utils.First(vnf_instance.Vendor.Get()))
	}
	if vnf_instance.Version.Present() {
		dao_vnf_instance_data.Version = string(utils.First(vnf_instance.Version.Get()))
	}
	return dao_vnf_instance_data, nil
}

func getDaoVnfInstanceId(vnf_instance_dao VnfInstanceDao) (int, error) {
	var vnf_instance_dao_id int
	err := vnf_instance_dao.session.Query("SELECT nextval('agorangmanager.vnf_instance_id_seq');", &vnf_instance_dao_id)
	if err != nil {
		log.Errorf("[VnfInstanceDao - getDaoVnfInstanceId]: %s", err.Error())
		return -1, err
	}
	return vnf_instance_dao_id, nil
}

package dao

import (
	helpers "agora-vnf-manager/core/helpers"
	log "agora-vnf-manager/core/log"
	utils "agora-vnf-manager/core/utils"
	types "agora-vnf-manager/features/vnf-infrastructure/types"
	"fmt"
)

func MapDaoVnfInfrastructureToVnfInfrastructure(dao_vnf_infrastructure DaoVnfInfrastructure) types.VnfInfrastructure {
	vnf_infrastructure := types.VnfInfrastructure{}
	vnf_infrastructure.Id = dao_vnf_infrastructure.Id
	vnf_infrastructure.Name.Scan(dao_vnf_infrastructure.Name)
	vnf_infrastructure.Description.Scan(dao_vnf_infrastructure.Description)
	vnf_infrastructure.ConfigurationFile.Scan(dao_vnf_infrastructure.ConfigurationFile)
	return vnf_infrastructure
}

func MapVnfInfrastructureToDaoVnfInfrastructure(vnf_infrastructure types.VnfInfrastructure, vnf_infrastructure_dao VnfInfrastructureDao) (DaoVnfInfrastructure, error) {
	dao_vnf_infrastructure := DaoVnfInfrastructure{}
	dao_vnf_infrastructure_id, err := getDaoVnfInfrastructureId(vnf_infrastructure_dao)
	if err != nil {
		log.Errorf("[MapVnfInfrastructureToDaoVnfInfrastructure]: %s", err.Error())
		return dao_vnf_infrastructure, err
	}
	dao_vnf_infrastructure.Id = dao_vnf_infrastructure_id
	dao_vnf_infrastructure.Name = string(utils.First(vnf_infrastructure.Name.Get()))
	dao_vnf_infrastructure.Description = string(utils.First(vnf_infrastructure.Description.Get()))
	dao_vnf_infrastructure.ConfigurationFile = string(utils.First(vnf_infrastructure.ConfigurationFile.Get()))
	return dao_vnf_infrastructure, nil
}

func MapVnfInfrastructureToDaoVnfInfrastructureFind(vnf_infrastructure types.VnfInfrastructure) DaoVnfInfrastructure {
	dao_vnf_infrastructure := DaoVnfInfrastructure{}
	if vnf_infrastructure.Name.Present() {
		dao_vnf_infrastructure.Name = string(utils.First(vnf_infrastructure.Name.Get()))
	}
	if vnf_infrastructure.Description.Present() {
		dao_vnf_infrastructure.Description = string(utils.First(vnf_infrastructure.Description.Get()))
	}
	if vnf_infrastructure.ConfigurationFile.Present() {
		dao_vnf_infrastructure.ConfigurationFile = string(utils.First(vnf_infrastructure.ConfigurationFile.Get()))
	}
	return dao_vnf_infrastructure
}

func MapVnfInfrastructureToDaoVnfInfrastructureUpdate(vnf_infrastructure types.VnfInfrastructure, id int, vnf_infrastructure_dao VnfInfrastructureDao) (DaoVnfInfrastructure, error) {
	dao_vnf_infrastructure_search := DaoVnfInfrastructure{Id: id}
	dao_vnf_infrastructure_data, has, err := helpers.GetEntry(vnf_infrastructure_dao.session, dao_vnf_infrastructure_search)
	if err != nil || !has {
		if err != nil {
			log.Errorf("[MapVnfInfrastructureToDaoVnfInfrastructureUpdate]: %s", err.Error())
			return DaoVnfInfrastructure{}, err
		}
		return DaoVnfInfrastructure{}, fmt.Errorf("Could not find referenced vnf_infrastructure with id: %s", err.Error())
	}
	if vnf_infrastructure.Name.Present() {
		dao_vnf_infrastructure_data.Name = string(utils.First(vnf_infrastructure.Name.Get()))
	}
	if vnf_infrastructure.Description.Present() {
		dao_vnf_infrastructure_data.Description = string(utils.First(vnf_infrastructure.Description.Get()))
	}
	if vnf_infrastructure.ConfigurationFile.Present() {
		dao_vnf_infrastructure_data.ConfigurationFile = string(utils.First(vnf_infrastructure.ConfigurationFile.Get()))
	}
	return dao_vnf_infrastructure_data, nil
}

func getDaoVnfInfrastructureId(vnf_infrastructure_dao VnfInfrastructureDao) (int, error) {
	var vnf_infrastructure_dao_id int
	err := vnf_infrastructure_dao.session.Query("SELECT nextval('agorangmanager.vnf_infrastructure_id_seq');", &vnf_infrastructure_dao_id)
	if err != nil {
		log.Errorf("[VnfInfrastructureDao - getDaoVnfInfrastructureId]: %s", err.Error())
		return -1, err
	}
	return vnf_infrastructure_dao_id, nil
}

package consul

import (
	log "agora-vnf-manager/core/log"

	consul "github.com/hashicorp/consul/api"
)

var config *consul.Config
var client *consul.Client

func InitConsulService() (err error) {
	config = consul.DefaultConfig()
	client, err = consul.NewClient(config)
	if err != nil {
		log.Errorf("[ConsulService - InitConsulService]: %s", err.Error())
		return err
	}
	return nil
}

func DiscoverServices() (services map[string][]string, err error) {
	services, _, err = client.Catalog().Services(nil)
	if err != nil {
		log.Errorf("[ConsulService - DiscoverServices]: %s", err.Error())
		return services, err
	}
	return services, nil
}

func DiscoverNodes() (nodes []*consul.Node, err error) {
	query_options := &consul.QueryOptions{}
	nodes, _, err = client.Catalog().Nodes(query_options)
	if err != nil {
		log.Errorf("[ConsulService - DiscoverNodes]: %s", err.Error())
		return nodes, err
	}
	return nodes, err
}

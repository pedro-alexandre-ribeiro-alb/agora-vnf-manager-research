package helm

import (
	log "agora-vnf-manager/core/log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

func ReadValuesFromYamlFile(file_path string) (values map[string]interface{}, err error) {
	content, err := os.ReadFile(file_path)
	if err != nil {
		log.Errorf("[HelmUtils - ReadValuesFromYamlFile]: %s", err.Error())
		return nil, err
	}
	err = yaml.Unmarshal(content, &values)
	if err != nil {
		log.Errorf("[HelmUtils - ReadValuesFromYamlFile]: %s", err.Error())
		return nil, err
	}
	return values, nil
}

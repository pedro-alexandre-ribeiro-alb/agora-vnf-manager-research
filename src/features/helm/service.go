package helm

import (
	"agora-vnf-manager/core/log"
	"agora-vnf-manager/core/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	types "agora-vnf-manager/features/helm/types"

	"gopkg.in/yaml.v3"
	action "helm.sh/helm/v3/pkg/action"
	loader "helm.sh/helm/v3/pkg/chart/loader"
	cli "helm.sh/helm/v3/pkg/cli"
	release "helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
)

func ListEnrolledRepositories() ([]*repo.Entry, error) {
	workdir, err := os.Getwd()
	if err != nil {
		log.Errorf("[HelmService - ListEnrolledRepositories]: %s", err.Error())
		return []*repo.Entry{}, err
	}
	repositories_file := filepath.Join(workdir, "../config/repositories.yaml")
	parsed_repositories_information, err := repo.LoadFile(repositories_file)
	if err != nil {
		log.Errorf("[HelmService - ListEnrolledRepositories]: %s", err.Error())
		return []*repo.Entry{}, err
	}
	return parsed_repositories_information.Repositories, nil
}

func ListRepositoryCharts(repository_name string) (map[string][]types.Chart, error) {
	workdir, err := os.Getwd()
	if err != nil {
		log.Errorf("[HelmService - ListRepositoryCharts]: %s", err.Error())
		return map[string][]types.Chart{}, err
	}
	repositories_file := filepath.Join(workdir, "../config/repositories.yaml")
	parsed_repositories_information, err := repo.LoadFile(repositories_file)
	if err != nil {
		log.Errorf("[HelmService - ListRepositoryCharts]: %s", err.Error())
		return map[string][]types.Chart{}, err
	}
	for _, repository := range parsed_repositories_information.Repositories {
		if repository.Name == repository_name {
			repo_url := repository.URL
			response, err := http.Get(repo_url + "/index.yaml")
			if err != nil {
				log.Errorf("[HelmService - ListRepositoryCharts]: %s", err.Error())
				return map[string][]types.Chart{}, err
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				log.Errorf("[HelmService - ListRepositoryCharts]: %s", err.Error())
				return map[string][]types.Chart{}, err
			}
			var index types.Index
			if err := yaml.Unmarshal(body, &index); err != nil {
				log.Errorf("[HelmService - ListRepositoryCharts]: %s", err.Error())
				return map[string][]types.Chart{}, err
			}
			return index.Entries, nil
		}
	}
	return map[string][]types.Chart{}, fmt.Errorf("Could not find the helm chart repository with the provided name")
}

func DeployHelmChart(specification types.Specification) (release *release.Release, err error) {
	action_config := new(action.Configuration)
	settings := cli.New()
	err = action_config.Init(settings.RESTClientGetter(), specification.Namespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		log.Infof("[HelmService - DeployHelmChart]: Starting - %s - %+v", format, v)
	})
	if err != nil {
		log.Errorf("[HelmService - DeployHelmChart]: %s", err.Error())
		return nil, err
	}
	client := action.NewInstall(action_config)
	client.Namespace = specification.Namespace
	client.ReleaseName = specification.ReleaseName
	cp, err := client.ChartPathOptions.LocateChart(string(utils.First(specification.ChartPath.Get())), settings)
	if err != nil {
		log.Errorf("[HelmService - DeployHelmChart]: %s", err.Error())
		return nil, err
	}
	chart_req, err := loader.Load(cp)
	values, err := ReadValuesFromYamlFile(string(utils.First(specification.ValuesPath.Get())))
	if err != nil {
		log.Errorf("[HelmService - DeployHelmChart]: %s", err.Error())
		return nil, err
	}
	release, err = client.Run(chart_req, values)
	if err != nil {
		log.Errorf("[HelmService - DeployHelmChart]: %s", err.Error())
		return nil, err
	}
	return release, nil
}

func UndeployHelmChart(specification types.Specification) (response *release.UninstallReleaseResponse, err error) {
	action_config := new(action.Configuration)
	settings := cli.New()
	err = action_config.Init(settings.RESTClientGetter(), specification.Namespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		log.Infof("[HelmService - UndeployHelmChart]: Starting - %s - %+v", format, v)
	})
	if err != nil {
		log.Errorf("[HelmService - UndeployHelmChart]: %s", err.Error())
		return nil, err
	}
	client := action.NewUninstall(action_config)
	response, err = client.Run(specification.ReleaseName)
	if err != nil {
		log.Errorf("[HelmService - UndeployHelmChart]: %s", err.Error())
		return nil, err
	}
	return response, nil
}

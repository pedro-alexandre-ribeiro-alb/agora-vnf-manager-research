package helm

import (
	"agora-vnf-manager/core/log"
	"agora-vnf-manager/core/utils"
	"os"

	types "agora-vnf-manager/features/helm/types"

	action "helm.sh/helm/v3/pkg/action"
	loader "helm.sh/helm/v3/pkg/chart/loader"
	cli "helm.sh/helm/v3/pkg/cli"
	release "helm.sh/helm/v3/pkg/release"
)

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

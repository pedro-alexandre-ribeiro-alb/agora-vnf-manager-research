package types

import optional "agora-vnf-manager/core/optional"

type Specification struct {
	ReleaseName       string                             `json:"releaseName"`
	Namespace         string                             `json:"namespace"`
	ChartPath         optional.Optional[optional.String] `json:"chartPath"`
	HelmRepositoryUrl optional.Optional[optional.String] `json:"helmRepositoryUrl"`
	ChartName         optional.Optional[optional.String] `json:"chartName"`
	ValuesPath        optional.Optional[optional.String] `json:"valuesPath"`
}

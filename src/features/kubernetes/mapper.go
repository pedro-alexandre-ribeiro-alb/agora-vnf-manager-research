package kubernetes

import (
	"agora-vnf-manager/core/utils"
	types "agora-vnf-manager/features/kubernetes/types"

	api_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MapSpecificationsToListOptions(specification types.Specifications) api_v1.ListOptions {
	list_options := api_v1.ListOptions{}
	if specification.Labels.Present() {
		labels := utils.First(specification.Labels.Get())
		if len(labels) > 0 {
			label_selectors := map[string]string{}
			for _, label := range utils.First(specification.Labels.Get()) {
				label_selectors[label.Name] = label.Value
			}
			label_selector := api_v1.LabelSelector{MatchLabels: label_selectors}
			list_options.LabelSelector = api_v1.FormatLabelSelector(&label_selector)
		}
	}
	return list_options
}

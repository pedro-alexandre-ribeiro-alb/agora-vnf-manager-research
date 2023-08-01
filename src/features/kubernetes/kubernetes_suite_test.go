package kubernetes

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	types "agora-vnf-manager/features/kubernetes/types"
)

func TestKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kubernetes Suite")
}

var kubernetes_mapper_test = Describe("Kubernetes Mapper Test", Ordered, func() {

	Context("Test MapSpecificationsToListOptions() - provided valid specification - no labels", func() {
		It("should return a valid list of api_v1.ListOptions", func() {
			specification := types.Specifications{ConfigurationFile: "/path/to/configuration/file.yaml"}
			list_options := MapSpecificationsToListOptions(specification)
			Expect(list_options).ToNot(BeNil())
			Expect(list_options.LabelSelector).To(BeEquivalentTo(""))
		})
	})

})

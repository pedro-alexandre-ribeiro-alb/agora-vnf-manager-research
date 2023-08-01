package helm

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHelm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helm Suite")
}

var helm_utils_test = Describe("Helm Utils Test", Ordered, func() {

	Context("Test ReadValuesFromYamlFile() - the file does not exists", func() {
		It("should return nil and an error", func() {
			file_path := "/home/pedro/Code/alticelabs/agora-vnf-manager/helm/dhcp/values.yaml"
			values, err := ReadValuesFromYamlFile(file_path)
			Expect(values).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})

	Context("Test ReadValuesFromYamlFile() - the file exists", func() {
		It("should return a filled map and no error", func() {
			file_path := "/home/pedro/Code/alticelabs/agora-vnf-manager/helm/dhcp/values-inst1.yaml"
			values, err := ReadValuesFromYamlFile(file_path)
			Expect(values).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
	})
})

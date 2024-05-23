package v1alpha1_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"os"
	"path"

	"sigs.k8s.io/yaml"

	"github.com/kumahq/kuma/pkg/core/resources/apis/meshexternalservice/api/v1alpha1"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/test/matchers"
)

var _ = Describe("MeshExternalService", func() {
	Describe("Validate()", func() {
		type testCase struct {
			file string
		}

		DescribeTable("should validate all fields and return as much individual errors as possible",
			func(given testCase) {
				// setup
				MeshExternalService := v1alpha1.NewMeshExternalServiceResource()

				// when
				contents, err := os.ReadFile(path.Join("testdata", given.file+".input.yaml"))
				Expect(err).ToNot(HaveOccurred())
				err = core_model.FromYAML(contents, &MeshExternalService.Spec)
				Expect(err).ToNot(HaveOccurred())
				// and
				verr := MeshExternalService.Validate()
				actual, err := yaml.Marshal(verr)
				if string(actual) == "null\n" {
					actual = []byte{}
				}
				Expect(err).ToNot(HaveOccurred())

				// then
				Expect(actual).To(matchers.MatchGoldenYAML(path.Join("testdata", given.file+".output.yaml")))
			},
			Entry("minimal passing example", testCase{
				file: "minimal-valid",
			}),
			Entry("full example without extension", testCase{
				file: "full-without-extension-valid",
			}),
			Entry("minimal failing example with unknown port", testCase{
				file: "minimal-invalid",
			}),
		)
	})
})

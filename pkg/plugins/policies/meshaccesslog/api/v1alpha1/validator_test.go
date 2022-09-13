package v1alpha1_test

import (
	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	meshaccesslog_proto "github.com/kumahq/kuma/pkg/plugins/policies/meshaccesslog/api/v1alpha1"
	util_proto "github.com/kumahq/kuma/pkg/util/proto"
)

var _ = Describe("MeshAccessLog", func() {
	Describe("Validate()", func() {
		DescribeTable("should pass validation",
			func(mtpYAML string) {
				// setup
				meshAccessLog := meshaccesslog_proto.NewMeshAccessLogResource()

				// when
				err := util_proto.FromYAML([]byte(mtpYAML), meshAccessLog.Spec)
				Expect(err).ToNot(HaveOccurred())
				// and
				verr := meshAccessLog.Validate()

				// then
				Expect(verr).To(BeNil())
			},
            Entry("mesh from/to example", `
targetRef:
  kind: MeshService
  name: web-frontend
from:
  - targetRef:
      kind: Mesh
      name: default
    default:
      backends:
        - tcp:
            format:
              json:
                - key: "start_time"
                  value: "%START_TIME%"
            address: 127.0.0.1:5000
        - reference:
            kind: MeshAccessLogBackend
            name: file-backend
to:
  - targetRef:
      kind: MeshService
      name: web-backend
    default:
      backends:
        - file:
           format:
             plain: '{"start_time": "%START_TIME%"}'
           path: '/tmp/logs.txt'
`),
			Entry("empty format", `
targetRef:
  kind: MeshService
  name: web-frontend
from:
  - targetRef:
      kind: Mesh
      name: default
    default:
      backends:
        - file:
           path: '/tmp/logs.txt'
`),
        )

		type testCase struct {
			inputYaml string
			expected  string
		}

		DescribeTable("should validate all fields and return as much individual errors as possible",
			func(given testCase) {
				// setup
				meshAccessLog := meshaccesslog_proto.NewMeshAccessLogResource()

				// when
				err := util_proto.FromYAML([]byte(given.inputYaml), meshAccessLog.Spec)
				Expect(err).ToNot(HaveOccurred())
				// and
				verr := meshAccessLog.Validate()
				actual, err := yaml.Marshal(verr)
				Expect(err).ToNot(HaveOccurred())

				// then
				Expect(actual).To(MatchYAML(given.expected))
			},
			Entry("empty 'from' and 'to' array", testCase{
				inputYaml: `
targetRef:
  kind: Mesh
  name: default
`,
				expected: `
violations:
  - field: spec
    message: at lest one of "from", "to" has to be defined`,
			}),
			Entry("empty 'path'", testCase{
				inputYaml: `
targetRef:
  kind: MeshService
  name: web-frontend
from:
  - targetRef:
      kind: Mesh
      name: default
    default:
      backends:
        - file:
           format:
             plain: '{"start_time": "%START_TIME%"}'
`,
				expected: `
violations:
  - field: spec.from[0].default.backend[0].file.path
    message: file backend requires a path`,
			}),
			Entry("empty 'key'", testCase{
				inputYaml: `
targetRef:
  kind: MeshService
  name: web-frontend
from:
  - targetRef:
      kind: Mesh
      name: default
    default:
      backends:
        - file:
           path: '/tmp/logs.txt'
           format:
             json:
                - value: "%START_TIME%"
`,
				expected: `
violations:
  - field: spec.from[0].default.backend[0].json[0].key
    message: key cannot be empty`,
			}),
			Entry("empty 'value'", testCase{
				inputYaml: `
targetRef:
  kind: MeshService
  name: web-frontend
from:
  - targetRef:
      kind: Mesh
      name: default
    default:
      backends:
        - file:
           path: '/tmp/logs.txt'
           format:
             json:
                - key: "start_time"
`,
				expected: `
violations:
  - field: spec.from[0].default.backend[0].json[0].value
    message: value cannot be empty`,
			}),
			Entry("both 'plain' and 'json' defined", testCase{
				inputYaml: `
targetRef:
  kind: MeshService
  name: web-frontend
from:
  - targetRef:
      kind: Mesh
      name: default
    default:
      backends:
        - tcp:
            address: 127.0.0.1:5000
            format:
              plain: '{"start_time": "%START_TIME%"}'
              json:
                - key: "start_time"
                  value: "%START_TIME%"
`,
				expected: `
violations:
- field: spec.from[0].default.backend[0]
  message: 'format can only have one type defined: plain, json'`,
			}),
			Entry("both 'tcp' and 'reference' defined", testCase{
				inputYaml: `
targetRef:
  kind: MeshService
  name: web-frontend
from:
  - targetRef:
      kind: Mesh
      name: default
    default:
      backends:
        - tcp:
            address: 127.0.0.1:5000
            format:
              json:
                - key: "start_time"
                  value: "%START_TIME%"
          reference:
            kind: MeshAccessLogBackend
            name: file-backend
`,
				expected: `
violations:
- field: spec.from[0].default.backend[0]
  message: 'backend can have only one type type defined: tcp, file, reference'`,
			}),

			Entry("both 'tcp' and 'reference' defined", testCase{
				inputYaml: `
targetRef:
  kind: MeshService
  name: web-frontend
from:
  - targetRef:
      kind: Mesh
      name: default
    default:
      backends:
        - tcp:
            address: 127.0.0.1:5000
            format:
              json:
                - key: "start_time"
                  value: "%START_TIME%"
          reference:
            kind: MeshAccessLogBackend
            name: file-backend
`,
				expected: `
violations:
- field: spec.from[0].default.backend[0]
  message: 'backend can have only one type type defined: tcp, file, reference'`,
			}),
		)
	})
})

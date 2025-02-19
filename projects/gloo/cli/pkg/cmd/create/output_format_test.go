package create_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/testutils"
)

var _ = Describe("Upstream", func() {

	BeforeEach(func() {
		helpers.UseMemoryClients()
	})

	Context("--dry-run should override -o table", func() {
		kubeYamlOutput := `apiVersion: gloo.solo.io/v1
kind: Upstream
metadata:
  creationTimestamp: null
  name: jsonplaceholder-80
  namespace: gloo-system
spec:
  upstreamSpec:
    static:
      hosts:
      - addr: jsonplaceholder.typicode.com
        port: 80
status: {}
`

		yamlOutput := `---
metadata:
  name: jsonplaceholder-80
  namespace: gloo-system
status: {}
upstreamSpec:
  static:
    hosts:
    - addr: jsonplaceholder.typicode.com
      port: 80
`

		It("--dry-run should override -o table and replace with kube-yaml", func() {
			By("should use kube-yaml format by default")
			output, err := testutils.GlooctlOut("create upstream static jsonplaceholder-80 --static-hosts jsonplaceholder.typicode.com:80 --dry-run")
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(kubeYamlOutput))

			By("should override -o table flag")
			output, err = testutils.GlooctlOut("create upstream static jsonplaceholder-80 --static-hosts jsonplaceholder.typicode.com:80 --dry-run -o table")
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(kubeYamlOutput))

			By("should respect -o kube-yaml flag")
			output, err = testutils.GlooctlOut("create upstream static jsonplaceholder-80 --static-hosts jsonplaceholder.typicode.com:80 --dry-run -o kube-yaml")
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(kubeYamlOutput))

			By("should respect -o yaml flag")
			output, err = testutils.GlooctlOut("create upstream static jsonplaceholder-80 --static-hosts jsonplaceholder.typicode.com:80 --dry-run -o yaml")
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(Equal(yamlOutput))
		})
	})
})

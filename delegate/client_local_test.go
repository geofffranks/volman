package delegate_test

import (
	"github.com/cloudfoundry-incubator/volman/delegate"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-golang/lager/lagertest"
)

var _ = Describe("ListDrivers", func() {

	Context("volman has no drivers", func() {
		BeforeEach(func() {
			client = &delegate.LocalClient{}
		})

		It("should report empty list of drivers", func() {
			testLogger := lagertest.NewTestLogger("ClientTest")
			drivers, err := client.ListDrivers(testLogger)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(drivers.Drivers)).To(Equal(0))
		})
	})

	Context("volman has drivers", func() {
		BeforeEach(func() {
			client = &delegate.LocalClient{fakeDriverPath}
		})

		It("should report list of drivers", func() {
			testLogger := lagertest.NewTestLogger("ClientTest")
			drivers, err := client.ListDrivers(testLogger)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(drivers.Drivers)).To(Equal(1))
		})
	})
})

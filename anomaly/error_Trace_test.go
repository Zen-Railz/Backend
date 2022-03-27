package anomaly

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tracing Error", func() {
	It("should populate the history of the ServiceError", func() {
		err := (&ServiceError{}).Trace()
		Expect(len(err.history)).To(Equal(1))
	})
})

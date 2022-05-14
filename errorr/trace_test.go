package errorr

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tracing Error", func() {
	var err *entity

	Context("Private trace method", func() {
		BeforeEach(func() {
			err = (&entity{}).trace(1)
		})

		It("should populate the history of the error entity", func() {
			Expect(len(err.stackTrace)).To(Equal(1))
		})

		It("should not have a history that contains the trace method", func() {
			Expect(err.stackTrace[0].FunctionName).NotTo(ContainSubstring("trace"))
		})
	})

	Context("Public Trace method", func() {
		When("only a single Trace is called", func() {
			BeforeEach(func() {
				e := (&entity{}).Trace()
				err = convertToEntityToStruct(e)
			})

			It("should populate the history of the error entity", func() {
				Expect(len(err.stackTrace)).To(Equal(1))
			})

			It("shoud not have a history that contains the Trace method", func() {
				Expect(err.stackTrace[0].FunctionName).NotTo(ContainSubstring("Trace"))
			})
		})

		When("only multiple Trace are called", func() {
			BeforeEach(func() {
				e := outerTraceCall()
				err = convertToEntityToStruct(e)
			})

			It("should populate the history of the error entity", func() {
				Expect(len(err.stackTrace)).To(Equal(3))
			})

			It("shoud contain stack trace with function name", func() {
				Expect(err.stackTrace[0].FunctionName).To(ContainSubstring("errorOrigin"))
				Expect(err.stackTrace[1].FunctionName).To(ContainSubstring("innerTraceCall"))
				Expect(err.stackTrace[2].FunctionName).To(ContainSubstring("outerTraceCall"))
			})
		})
	})
})

func convertToEntityToStruct(e Entity) *entity {
	return e.(*entity)
}

func outerTraceCall() Entity {
	return innerTraceCall().Trace()
}

func innerTraceCall() Entity {
	return errorOrigin().Trace()
}

func errorOrigin() Entity {
	return (&entity{}).Trace()
}

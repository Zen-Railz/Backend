package anomaly

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elaborating the details of ServiceError", func() {
	var err *ServiceError

	BeforeEach(func() {
		err = &ServiceError{
			Code:    mockServiceErrorCode,
			Message: mockServiceErrorMessage,
		}
	})

	Context("has trace history", func() {
		When("only one record and no annex", func() {
			It("should output code, message and the single line of history", func() {
				errHistory := errorHistory{
					functionName: mockServiceErrorFunctionName,
					lineNumber:   mockServiceErrorLineNumber,
				}
				err.history = append(err.history, errHistory)
				expectedOutput := fmt.Sprintf("(%s) %s\n1. %s | Line %d\n", err.Code, err.Message, errHistory.functionName, errHistory.lineNumber)
				output := err.Elaborate()
				Expect(output).To(Equal(expectedOutput))
			})
		})

		When("multiple records and no annex", func() {
			It("should output code, message and some lines of history", func() {
				firstHistory := errorHistory{
					functionName: mockServiceErrorFunctionName,
					lineNumber:   mockServiceErrorLineNumber,
				}
				secondHistory := errorHistory{
					functionName: mockServiceErrorFunctionName,
					lineNumber:   mockServiceErrorLineNumber,
				}
				err.history = append(err.history, firstHistory, secondHistory)
				expectedOutput := fmt.Sprintf("(%s) %s\n1. %s | Line %d\n2. %s | Line %d\n", err.Code, err.Message, firstHistory.functionName, firstHistory.lineNumber, secondHistory.functionName, secondHistory.lineNumber)
				output := err.Elaborate()
				Expect(output).To(Equal(expectedOutput))
			})
		})

		When("there is annex", func() {
			It("should output code, message, some lines of history and annex", func() {
				errHistory := errorHistory{
					functionName: mockServiceErrorFunctionName,
					lineNumber:   mockServiceErrorLineNumber,
				}
				errAnnex := struct {
					text string
				}{
					text: "mock_annex",
				}
				err.history = append(err.history, errHistory)
				err.Annex = errAnnex
				expectedOutput := fmt.Sprintf("(%s) %s\n1. %s | Line %d\n%+v", err.Code, err.Message, errHistory.functionName, errHistory.lineNumber, errAnnex)
				output := err.Elaborate()
				Expect(output).To(Equal(expectedOutput))
			})
		})
	})
})

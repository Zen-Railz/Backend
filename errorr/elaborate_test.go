package errorr

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elaborating the details of the error", func() {
	var err *entity

	BeforeEach(func() {
		err = &entity{
			code:    mockErrorCode,
			message: mockErrorMessage,
		}
	})

	Context("has trace history", func() {
		When("only one record and no annex", func() {
			It("should output code, message and the single line of history", func() {
				errHistory := StackTrace{
					FunctionName: mockErrorFunctionName,
					LineNumber:   mockErrorLineNumber,
				}
				err.stackTrace = append(err.stackTrace, errHistory)
				expectedOutput := fmt.Sprintf("(%s) %s\n1. %s | Line %d\n", err.code, err.message, errHistory.FunctionName, errHistory.LineNumber)
				output := err.Elaborate()
				Expect(output).To(Equal(expectedOutput))
			})
		})

		When("multiple records and no annex", func() {
			It("should output code, message and some lines of history", func() {
				firstHistory := StackTrace{
					FunctionName: mockErrorFunctionName,
					LineNumber:   mockErrorLineNumber,
				}
				secondHistory := StackTrace{
					FunctionName: mockErrorFunctionName,
					LineNumber:   mockErrorLineNumber,
				}
				err.stackTrace = append(err.stackTrace, firstHistory, secondHistory)
				expectedOutput := fmt.Sprintf("(%s) %s\n1. %s | Line %d\n2. %s | Line %d\n", err.code, err.message, firstHistory.FunctionName, firstHistory.LineNumber, secondHistory.FunctionName, secondHistory.LineNumber)
				output := err.Elaborate()
				Expect(output).To(Equal(expectedOutput))
			})
		})

		When("there is annex", func() {
			It("should output code, message, some lines of history and annex", func() {
				errHistory := StackTrace{
					FunctionName: mockErrorFunctionName,
					LineNumber:   mockErrorLineNumber,
				}
				errAnnex := struct {
					text string
				}{
					text: "mock_annex",
				}
				err.stackTrace = append(err.stackTrace, errHistory)
				err.annex = errAnnex
				expectedOutput := fmt.Sprintf("(%s) %s\n1. %s | Line %d\n%+v", err.code, err.message, errHistory.FunctionName, errHistory.LineNumber, errAnnex)
				output := err.Elaborate()
				Expect(output).To(Equal(expectedOutput))
			})
		})
	})
})

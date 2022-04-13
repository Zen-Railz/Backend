package native

import (
	"fmt"
	"reflect"
	"zenrailz/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Checking the health of the system", func() {
	var serviceUnderTest *Service

	When("all sub-systems are healthy", func() {
		BeforeEach(func() {
			mockLogger := mock.NewLogger()
			mockDbRepo := mock.NewDatabaseRepository()
			serviceUnderTest = NewService(mockLogger, mockDbRepo)
		})

		It("should return overall status as healthy", func() {
			systemState := serviceUnderTest.SystemHealth()
			Expect(systemState.Status).To(Equal(Healthy))
		})

		It("should return healthy for all sub-system status", func() {
			systemState := serviceUnderTest.SystemHealth()
			properties := reflect.ValueOf(*systemState)
			for i := 0; i < properties.NumField(); i++ {
				fieldName := properties.Type().Field(i).Name
				fieldValue := properties.Field(i)

				if fieldName != "Status" {
					subSystemHealth := fmt.Sprintf("%s", fieldValue)
					Expect(subSystemHealth).To(Equal(Healthy))
				}
			}
		})
	})

	When("any sub-system is unhealthy", func() {
		BeforeEach(func() {
			mockLogger := mock.NewLogger()
			mockDbRepo := mock.NewDatabaseRepository().
				SetPingError()
			serviceUnderTest = NewService(mockLogger, mockDbRepo)
		})

		It("should return overall status as unhealthy", func() {
			systemState := serviceUnderTest.SystemHealth()
			Expect(systemState.Status).To(Equal(Unhealthy))
		})

		It("should return unhealthy for some sub-system status", func() {
			systemState := serviceUnderTest.SystemHealth()
			properties := reflect.ValueOf(*systemState)

			errorCount := 0
			for i := 0; i < properties.NumField(); i++ {
				fieldName := properties.Type().Field(i).Name
				fieldValue := properties.Field(i)

				if fieldName != "Status" {
					subSystemHealth := fmt.Sprintf("%s", fieldValue)
					if subSystemHealth == Unhealthy {
						errorCount++
					}
				}
			}

			Expect(errorCount).To(BeNumerically(">", 0))
		})
	})
})

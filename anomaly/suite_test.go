package anomaly

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAnomaly(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Anomaly Test Suite")
}

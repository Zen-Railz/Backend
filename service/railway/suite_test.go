package railway

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRailwayService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Railway Service Test Suite")
}

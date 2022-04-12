package native

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestNativeService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Native Service Test Suite")
}

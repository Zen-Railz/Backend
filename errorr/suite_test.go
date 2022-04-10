package errorr

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestErrorr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Errorr Test Suite")
}

package facebook

import (
	"testing"

	"github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OAuth Facebook Suite")
}

var _ = BeforeSuite(func() {
	logrus.Info("OAuth Facebook Suite starts")
})

var _ = AfterSuite(func() {
	logrus.Printf("OAuth Facebook Suite ends")
})

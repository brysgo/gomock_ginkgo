package gomock_ginkgo

import (
  gomock "code.google.com/p/gomock/gomock"
  "fmt"
  ginkgo "github.com/onsi/ginkgo"
)

type GinkgoTestReporter struct{}

func (g GinkgoTestReporter) Errorf(format string, args ...interface{}) {
  ginkgo.Fail(fmt.Sprintf(format, args))
}

func (g GinkgoTestReporter) Fatalf(format string, args ...interface{}) {
  ginkgo.Fail(fmt.Sprintf(format, args))
}

var t GinkgoTestReporter

func NewController() *gomock.Controller {
  return gomock.NewController(t)
}

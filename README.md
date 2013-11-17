Gomock-Ginkgo
=============

Gomock needs a reference to `testing.T`

This is a workaround adopted from [this thread.](https://github.com/onsi/ginkgo/issues/9)

```
import (
  gomock "code.google.com/p/gomock/gomock"
  gomock_ginkgo "github.com/brysgo/gomock_ginkgo"
)

var _ = Describe("Handler", func() {
    var (
        t GinkgoTestReporter
        mockCtrl *gomock.Controller
        mockDatabase *models.MockDatabase
        handler *Handler
    )

    BeforeEach(func() {
        mockCtrl = gomock_ginkgo.NewController() // This line gives you a mock controller without `testing.T`
        mockDatabaseModel = models.NewMockDatabase(mockCtrl)
        handler = NewHandler(mockDatabase)
    })

    AfterEach(func() {
        mockCtrl.Finish()
    })

    // It() calls here
})
```

<p xmlns:dct="http://purl.org/dc/terms/" xmlns:vcard="http://www.w3.org/2001/vcard-rdf/3.0#">
  <a rel="license"
     href="http://creativecommons.org/publicdomain/zero/1.0/">
    <img src="http://i.creativecommons.org/p/zero/1.0/88x31.png" style="border-style: none;" alt="CC0" />
  </a>
</p>

[![GoDoc](https://godoc.org/github.com/xaionaro-go/multierror?status.svg)](https://pkg.go.dev/github.com/xaionaro-go/multierror?tab=doc)
[![go report](https://goreportcard.com/badge/github.com/xaionaro-go/multierror)](https://goreportcard.com/report/github.com/xaionaro-go/multierror)

---

Sometimes it is required to return multiple errors from one function. This
package implements exactly such cases:

```go
package mylib

import (
    "sync"

    "github.com/xaionaro-go/multierror"
)

func SomeConcurrentFunc() error {
    var err multierror.SyncSlice
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        err.Add(someFunc1())
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        err.Add(someFunc2())
    }()

    wg.Wait()
    return err.ReturnValue()
}
```
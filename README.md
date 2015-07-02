
# aeimagesflags

This project lets you use the undocumented flags of the App Engine Images service URLs to serve transformed images.

## Documentation

[![GoDoc](https://godoc.org/github.com/ernestoalejo/aeimagesflags?status.svg)](https://godoc.org/github.com/ernestoalejo/aeimagesflags)

## Usage example

```go
package example

import (
  "fmt"

  "github.com/ernestoalejo/aeimagesflags"
  "google.golang.org/appengine"
  "google.golang.org/appengine/images"
)

func demo(blobKey appengine.BlobKey) error {
  url, err := images.CreateServingURL(blobKey, nil)
  if err != nil {
    return err
  }

  flags := aeimagesflags.Flags{
    Width: 300,
    Height: 100,
    Crop: true,
  }
  fmt.Println(aeimagesflags.Apply(url.String(), flags))

  return nil
}
```

## License

```
Copyright (c) 2015, Ernesto Alejo
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

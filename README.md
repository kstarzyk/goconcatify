# goconcatify [development]
Golang [image] concat library

## Example
```go
package main

import "github.com/kstarzyk/concatify"

func main() {
  concated, err := concatify.NewConcatedImage([]string{"/path/to/image1", "path/to/image2"}) 
  if err != nil {
    ...
  }
  concated.Draw("./output-vertical.png")
  concatedHorizontal, err := concatify.NewConcatedImage([]string{"/path/to/image1", "path/to/image2", ConcatedImageOptions{HORIZONTAL, false, false})
    if err != nil {
    ...
  }
  concatedHorizontal.Draw("./output-horizontal.png")
}
```

## Todo 
- Support different image formats (*.jpg, *.png)
- Docs


## Tests
```bash
go test -v
```


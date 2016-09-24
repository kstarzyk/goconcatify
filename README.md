# goconcatify [development]
Golang [image] concat library

## Example
```go
package main

import "github.com/kstarzyk/concatify"

func main() {
  concated := concatify.NewConcatedImage([]string{"/path/to/image1", "path/to/image2"}) 
  concated.Draw("./output-vertical.png")
  concatedHorizontal := concatify.NewConcatedImage([]string{"/path/to/image2", "path/to/image3", ConcatedImageOptions{HORIZONTAL, false, false})
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


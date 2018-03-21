# goconcatify [![Build Status](https://travis-ci.org/kstarzyk/goconcatify.svg?branch=tests%2Frefactor)](https://travis-ci.org/kstarzyk/goconcatify)
Golang [image] concat library

## Example
```go
package main

import "github.com/kstarzyk/concatify"

func main() {
  paths := []string{"alpha.png", "omega.png"}
  concated, err := concatify.NewVertical(paths) 
  // concated, err := concatify.NewHorizontal(paths)  
  // concated, err := concatify.NewGrid([]string{"alpha.png", "omega.png", "omega.png", "alpha.png"}, 2, 2)  

  if err != nil {
    ...
  }
  concated.Save("./output.png")
}
```

## Todo 
- [x] Horizontal/Vertical
- [x] Grid 
- [ ] Resizng
- [ ] Support different image formats (*.jpg, *.png)
- [ ] Docs


## Tests
```bash
go test -v
```


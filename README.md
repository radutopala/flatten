# Go Flatten

[![Github Actions](https://github.com/radutopala/flatten/workflows/Test/badge.svg)](https://github.com/radutopala/flatten/actions)

Go lib for flattening Go array and slices.

## Installation
```
go get github.com/radutopala/flatten
```

## Usage

```
package main

import (
  "fmt"
  "github.com/radutopala/flatten"
)

func main() {
  fmt.Println(flatten.It([]interface{}{[]interface{}{1, 2, []int{3}}, 4}))
}

```

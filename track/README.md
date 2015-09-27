# ace-cors

Track middleware for [ACE](https://github.com/plimble/ace)

## Installation

``` bash
$ go get github.com/sbusso/ace-contrib/track
```

## Usage

``` go
import (
    "github.com/plimble/ace"
    "github.com/sbusso/ace-contrib/track"
)

func main(){
    a := ace.New()
    a.Use(cors.Cors(cors.Options{}))
}
```

[ACE]: https://github.com/plimble/ace

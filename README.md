[![Go Report Card](https://goreportcard.com/badge/github.com/dreddsa5dies/goStackQuery)](https://goreportcard.com/report/github.com/dreddsa5dies/goStackQuery) [![GORef](https://godoc.org/github.com/dreddsa5dies/goStackQuery?status.svg)](https://godoc.org/github.com/dreddsa5dies/goStackQuery) ![License](https://img.shields.io/badge/License-MIT-blue.svg)

![IMAGE](img/logo.png)

### Go (golang) package for search error on [Stackoverflow](https://stackoverflow.com)

## Install
```bash
go get -v github.com/dreddsa5dies/goStackQuery
```

## Package used
```bash
go get -v -u github.com/toqueteos/webbrowser
```

## Example
Examples of use, see [examples](https://github.com/dreddsa5dies/goStackQuery/tree/master/_examples)
```Go
package main

import (
	"os"

	goStackQuery "github.com/dreddsa5dies/goStackQuery"
)

func main() {
	file, err := os.Open("file")
	if err != nil {
		goStackQuery.Query(err, 1)
	}
	defer file.Close()

	file.WriteString("Aloha")
}
```

## Switchers:
0 - Query ERROR and paste link on STDERR  
![IMAGE](img/switcher_0.png)  
1 - Query ERROR on web browser and automate open link  
![IMAGE](img/switcher_1.png)

## The code contains comments in Russian

## License
This project is licensed under MIT license. Please read the LICENSE file.

## Contribute
Welcomes any kind of contribution. Please read the CONTRIBUTING and CODE_OF_CONDUCT file.
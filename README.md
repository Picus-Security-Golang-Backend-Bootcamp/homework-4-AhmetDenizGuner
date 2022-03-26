# Golang Library Command Line Project

## About the project

This application provides 9 featured library appliation from rest API. This app take 6 commands such as **list** , **search**, **buy**, **delete**, **add** and **update** for books. **get** , **search** , **add** for authors.

paths:

/book/list

/book/search/{search_item}

/book/buy --> query params id,amount

/book/delete/{id:[0-9]+}

/book/update --> query params id,amount

/book/add

/author/get{id:[0-9]+}

/author/search/{name}

/author/add



## Notes

*   You can put the searchelper directory in src directory that is in `GOPATH` . In this case you need to update import statment.
```go
package main


import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	data "workspace/data"
	"searchhelper"
	model "workspace/models"
)
```


* If there is any problem accsesing package or GOPATH you can set again GOPATH and you can use this command below and go visit this [link](https://stackoverflow.com/questions/68693154/package-is-not-in-goroot).
```bash
foo@bar:~$ go env -w GO111MODULE=off
```

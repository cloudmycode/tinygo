package main

import (
	"tinygo/tinygo"
)

func main() {
	r := tinygo.New()
	r.Static("/assets", "/Users/wangyang/Documents/Code/gopath/src/tinygo/static")
	r.Run(":9999")
}

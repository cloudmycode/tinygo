package main

import "tinygo/tiny"

func main() {
	r := tiny.New()
	r.Static("/assets", "/Users/wangyang/Documents/Code/gopath/src/tinygo/static")
	r.Run(":9999")
}

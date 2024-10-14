package main

import (
	"time"
	"xddq/cmd/xddq/app"
)

func init() {
	time.Local = time.FixedZone("CST", 8*3600)
}

func main() {
	if err := app.Main(); err != nil {
		panic(err)
	}
}

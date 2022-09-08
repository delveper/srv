package main

import (
	"log"

	"github.com/delveper/srv/core"
)

func main() {
	r := new(core.HttpRouter)

	if err := core.Run(r); err != nil {
		log.Fatal(err)
	}
}

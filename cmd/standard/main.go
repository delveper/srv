package main

import (
	"log"
)

func main() {
	r := new(core.StdMux)

	if err := core.Run(r); err != nil {
		log.Fatal(err)
	}
}

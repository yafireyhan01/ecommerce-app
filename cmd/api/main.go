package main

import (
	"github.com/yafireyhan01/synapsis-test/internal/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

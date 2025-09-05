package main

import (
	"fmt"
	"log"

	"github.com/Cheemx/gator/internal/config"
)

func main() {
	conf, err := config.Read()
	if err != nil {
		log.Panic(err)
	}
	err = conf.SetUser("Cheems")
	if err != nil {
		log.Panic(err)
	}
	conf, err = config.Read()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(conf)
}

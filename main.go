package main

import (
	"fmt"
	"log"

	"github.com/mrcordova/gator/internal/config"
)


func main()  {
cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("noah")
	if err != nil {
		log.Fatalf("Set User failed: %v", err )
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
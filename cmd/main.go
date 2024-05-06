package main

import (
	"log"
	"os"
	"push_swap/internal/ps"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: push_swap [numbers]")
	}
	ps.Run(os.Args[1:])
	// err := ps.Run([]string{"1", "4", "5", "7", "6"})
	// if err != nil {
	// 	panic(err)
	// }
}

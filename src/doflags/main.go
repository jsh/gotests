package main

import (
	"fmt"
	"flag"
)

func main() {
	hPtr := flag.Bool("help", false, "help")
	vPtr := flag.Bool("version", false, "version")

	flag.Parse()
	fmt.Println("help: ", *hPtr)
	fmt.Println("version: ", *vPtr)

	if *vPtr == false {
		fmt.Println("eew")
	} else {
		fmt.Println("Okay!")
	}
}

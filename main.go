package main

import (
	"fmt"
	"github.com/adamyordan/go-die/die"
	"log"
	"os"
)

func main()  {
	if len(os.Args) < 2 {
		log.Fatalf("usage: go-die.exe [file_path]")
	}
	res, err := die.DIEScan(os.Args[1], die.DIE_SHOWERRORS)
	if err != nil {
		log.Fatalf("error running die scan: %v", err)
	}
	fmt.Printf("DIE Scan successful.\n\nResult:\n%s\n", res)
}

package main

import (
	"log"
	"os"
)

func main() {
	if rcErr != nil {
		log.Println(rcErr.Error())
		os.Exit(1)
		return
	}
	log.Println(rc.GetBaseURL())
	return
}

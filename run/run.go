package main

import (
	"log"
	"os"

	"github.com/gauravsarma1992/gopyast/gopyast"
)

func main() {

	var (
		err error
		mgr *gopyast.Manager
	)

	log.Println("Starting gopyast program")

	if mgr, err = gopyast.New(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	if err = mgr.Run(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	return
}

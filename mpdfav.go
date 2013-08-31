package main

import (
	"flag"
	. "github.com/vincent-petithory/mpdclient"
	. "github.com/vincent-petithory/mpdfav"
	"log"
	"sync"
)

var noRatings = flag.Bool("no-ratings", false, "Disable ratings service")
var noPlaycounts = flag.Bool("no-playcounts", false, "Disable playcounts service")

func startMpdService(mpdc *MPDClient, service func(*MPDClient), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		service(mpdc)
	}()
}

func main() {
	var wg sync.WaitGroup

	mpdc, err := Connect("localhost", 6600)
	if err != nil {
		panic(err)
	}
	defer mpdc.Close()

	if !*noPlaycounts {
		startMpdService(mpdc, RecordPlayCounts, &wg)
		log.Println("Started Playcounts service... ")
	}
	if !*noRatings {
		startMpdService(mpdc, ListenRatings, &wg)
		log.Println("Started Ratings service... ")
	}

	wg.Wait()
}

package main

import (
	"searcher/internal/fetcher"
	"searcher/internal/processor"
	"fmt"
	"time"
)

func main() {
	priceChannel := make(chan float64)
	done := make(chan bool)
	start := time.Now()


	go fetcher.FetchPrices(priceChannel)

	go func(){
		processor.ShowPriceAVG(priceChannel, done)
	}()

	<- done
	
	fmt.Printf("%s total time for search\n", time.Since(start))
	fmt.Printf("----------\n")
}

package main

import (
	"searcher/internal/fetcher"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	price1 := fetcher.FetchPriceFromSite1()
	price2 := fetcher.FetchPriceFromSite2()
	price3 := fetcher.FetchPriceFromSite3()

	fmt.Printf("----------\n")
	fmt.Printf("$%.2f for site1\n", price1)
	fmt.Printf("$%.2f for site2\n", price2)
	fmt.Printf("$%.2f for site3\n", price3)

	fmt.Printf("%s total time for search\n", time.Since(start))
	fmt.Printf("----------\n")
}

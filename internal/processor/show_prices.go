package processor

import (
	"searcher/internal/models"
	"fmt"
)

func ShowPriceAVG(priceChannel <- chan models.PriceDetail, done chan <- bool) {
	var totalPrice float64
	countPrices := 0.0

	for price := range priceChannel {
		totalPrice += price.Value
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("$%.2f for site %s | avarage is: $%.2f\n", price.Value, price.StoreName, avgPrice)
	}

	done <- true
}

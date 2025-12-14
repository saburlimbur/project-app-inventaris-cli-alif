package utils

import "time"

const DepreciationRate = 0.2

func YearsUsed(purchaseDate time.Time) int {
	years := int(time.Since(purchaseDate).Hours() / 24 / 365)
	if years < 0 {
		return 0
	}
	return years
}

func DecliningBalance(price float64, years int) float64 {
	value := price
	for i := 0; i < years; i++ {
		value *= (1 - DepreciationRate)
	}
	return value
}

package main

import "fmt"

func DiscountForTotal(total float64) float64 {
	switch {
	case total >= 10000:
		return 20
	case total >= 5000:
		return 10
	default:
		return 0
	}
}

func GetCustomerStatus(total float64) string {
	switch {
	case total >= 10000:
		return "VIP-покупатель"
	case total >= 5000:
		return "Постоянный покупатель"
	default:
		return "Обычный покупатель"
	}
}

func CalculateDiscount(price, discount float64) float64 {
	if discount < 0 || discount > 100 {
		return price
	}
	return price - (price * discount / 100)
}

func main() {
	totals := []float64{320, 6000, 12500}

	for _, total := range totals {
		percent := DiscountForTotal(total)
		status := GetCustomerStatus(total)
		final := CalculateDiscount(total, percent)

		fmt.Printf("Сумма: %.0f ₽ | Статус: %s | Скидка: %.0f%% | Итого: %.0f ₽\n",
			total, status, percent, final)
	}
}
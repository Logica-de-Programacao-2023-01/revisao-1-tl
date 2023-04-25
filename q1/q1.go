package q1

import "fmt"

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	// Implemente sua solução aqui
	if currentPurchase <= 0 {
		return 0, fmt.Errorf("O valor da compra é invalido")

	}
	var totalPurchases float64 = 0
	for _, purchase := range purchaseHistory {
		totalPurchases += purchase
	}
	var discount float64 = 0
	if len(purchaseHistory) == 0 {
		discount = 0.1 * currentPurchase
	} else if totalPurchases > 1000 {
		discount = 0.1 * currentPurchase
	} else if totalPurchases <= 1000 {
		discount = 0.05 * currentPurchase
	}
	if totalPurchases <= 500 {
		discount = 0.2 * currentPurchase
	}
	if totalPurchases/float64(len(purchaseHistory)) > 1000 {
		if 0.2*currentPurchase > discount {
			discount = 0.2 * currentPurchase
		}
	}

	return discount, nil
}
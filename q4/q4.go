package q4

import "errors"

func CalculateFinalPrice(basePrice float64, state string, taxCode int) (float64, error) {
	if basePrice <= 0 {
		return 0, errors.New("preço base inválido")
	}
	var imposto, frete float64

	if taxCode == 1 {
		imposto = 0.05
	} else if taxCode == 2 {
		imposto = 0.1
	} else if taxCode = 3 {
		imposto = 0.15
	} else {
		return 0, errors.New("imposto não encontrado")
	}

	if state == "SP" {
		frete = 0.1
	} else if state == "RJ" {
		frete = 0.2
	} else if state == "MG" {
		frete = 0.2
	} else if state == "ES" {
		frete = 0.25
	} else {
		frete = 0.3
	}

	precoFinal := basePrice + basePrice*imposto + basePrice*frete

	return precoFinal, nil
}

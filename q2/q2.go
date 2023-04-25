package q2

import (
	"fmt"
	"strings"
)

func AverageLettersPerWord(text string) (float64, error) {
	palavras := strings.Fields(text)
	if len(palavras) == 0 {
		return 0, fmt.Errorf("texto vazio")
	}
	var numLetras int
	for _, palavra := range palavras {
		numLetras += len(palavra)
	}
	return float64(numLetras) / float64(len(palavras)), nil
}

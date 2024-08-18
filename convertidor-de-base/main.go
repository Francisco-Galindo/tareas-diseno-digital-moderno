package main

// Un programa que convierte un número escrito de una base a otra

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strings"
)

// HashMap bidireccional para pasar de símbolos a valores numéricos
type symbolMap struct {
	forward  map[string]int
	backward map[int]string
}

func main() {
	var s string
	var inBase int
	var outBase int

	flag.StringVar(&s, "n", "", "Número a procesar")
	flag.IntVar(&inBase, "i", 2, "Base de origen")
	flag.IntVar(&outBase, "o", 10, "Base de salida")
	flag.Parse()

	sMap := new(symbolMap)
	sMap.forward = map[string]int{}
	sMap.backward = map[int]string{}

	// Llena los primeros diez dígitos
	for i := 0; i < 10; i++ {
		sMap.forward[fmt.Sprintf("%d", i)] = i
		sMap.backward[i] = fmt.Sprintf("%d", i)
	}

	// Llena con las letras
	for i := 0; i < 26; i++ {
		sMap.forward[string(i+'A')] = 10 + i
		sMap.backward[10+i] = string(i + 'A')
	}

	fmt.Println(baseToBase(s, inBase, outBase, *sMap))
}

func stringToInt(s string, base int, symbols map[string]int) (res int) {
	for i, c := range s {
		val := symbols[string(c)]
		if val > base {
			log.Panic("La cadena de entrada usa símbolos que no corresponden con la base")
		}
		res += val * int(math.Pow(float64(base), float64(len(s)-i-1)))
	}

	return
}

func intToString(n int, base int, symbols map[int]string) (s string) {
	for n > 0 {
		s = symbols[n%base] + s
		n /= base
	}

	return
}

func baseToBase(in string, inBase, outBase int, sMap symbolMap) string {
	tokens := strings.Split(in, ".")

	integerPart := stringToInt(tokens[0], inBase, sMap.forward)
	integerPartStr := intToString(integerPart, outBase, sMap.backward)

	mantissaPartStr := ""
	if len(tokens) > 1 {
		mantissaPart := stringToMant(tokens[1], inBase, sMap.forward)
		mantissaPartStr = "." + mantToString(mantissaPart, outBase, sMap.backward)
	}

	return integerPartStr + mantissaPartStr
}

func mantToString(m float64, base int, symbols map[int]string) (s string) {
	for i := 0; m > math.Pow10(-8.0) && i < 8; i++ {
		m *= float64(base)
		s += symbols[int(math.Floor(m))]
		m -= math.Floor(m)
	}

	return
}

func stringToMant(s string, base int, symbols map[string]int) (m float64) {
	for i, c := range s {
		m += float64(symbols[string(c)]) * math.Pow(float64(base), -float64(i+1))
	}

	return
}

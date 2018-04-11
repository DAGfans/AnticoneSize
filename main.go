package main

import (
	"math"
	"log"
)

func main() {
	delay := 1.0
	rate := 10.00
	threshold := 0.5

	antiConeSize(delay, rate, threshold)

}

func antiConeSize(_delay, _rate, _threshold float64) {

	factor := 2 * _delay * _rate
	coef := 1 / (math.Pow(math.E, factor) - 1)

	log.Println("factor:", factor, "coef:", coef)

	sum := 0.0

	for k := 1; k < 1000; k++ {

		sum = 0.0
		for j := k + 1; j < k+100; j++ {
			xx := 1.0
			for jj := 1; jj <= j; jj++ {
				xx *= factor / float64(jj)
			}
			sum += xx
		}
		sum *= coef

		if sum < _threshold {
			log.Println(k, sum)
			break
		}
	}
}

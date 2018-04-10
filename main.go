package main

import (
	"math"
	"log"
)

func main() {
	delay := 40.0
	rate := 1/60.0
	threshold := 0.5

	antiConeSize(delay, rate, threshold)

}

func antiConeSize(_delay, _rate, _threshold float64) {

	factor := 2 * _delay * _rate
	coef := 1 / (1/math.Pow(math.E, -1*factor) - 1)

	sum := 0.0

	for k := 1; k < 100; k++ {

		sum = 0.0
		for j := k + 1; j < k+10; j++ {
			xx := math.Pow(factor, float64(j))
			for jj := 2; jj <= j; jj++ {
				xx /= float64(jj)
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

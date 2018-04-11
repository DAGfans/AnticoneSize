package main

import (
	"math"
	"log"
	"flag"
)

func main() {

	delay := flag.Float64("delay", 40.0, "Max propagation delay, unit is second")
	rate := flag.Float64("rate", 1.0, "Block rate, unit is blocks/second")
	threshold := flag.Float64("threshold", 0.5, "Percentage of honest nodes, value between 0 and 1")

	flag.Parse()

	if *delay < 0 || *rate < 0 || *threshold < 0 {
		flag.Usage()
		log.Fatal("Keep parameters above zero!", *delay, *rate, *threshold)
	}

	if *threshold > 1.0 {
		flag.Usage()
		log.Fatal("Keep threshold under one!", *threshold)
	}

	antiConeSize(*delay, *rate, *threshold)

}

func antiConeSize(_delay, _rate, _threshold float64) {

	factor := 2 * _delay * _rate
	coef := 1 / (math.Pow(math.E, factor) - 1)

	log.Printf("_delay:%v _rate:%v _threshold:%v factor:%v coef:%v", _delay, _rate, _threshold, factor, coef)

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

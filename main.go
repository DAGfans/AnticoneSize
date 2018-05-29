package main

import (
	"math"
	"flag"
	"log"
	"os"
)

var myLog = log.New(os.Stdout, "", 0)

func main() {

	delay := flag.Float64("delay", 15, "Max propagation delay, unit is second")
	rate := flag.Float64("rate", 2.0, "Block rate, unit is blocks/second")
	level := flag.Float64("level", 0.01, "Security level, the probability of an honest block being marked red")

	flag.Parse()

	if *delay < 0 || *rate < 0 || *level < 0 {
		flag.Usage()
		myLog.Fatalf("keep parameters above zero! delay:%v rate:%v level:%v", *delay, *rate, *level)
	}

	if *level > 1.0 {
		flag.Usage()
		myLog.Fatalf("keep level under one! level:%v", *level)
	}

	antiConeSize(*delay, *rate, *level)

}

func antiConeSize(_delay, _rate, _level float64) {

	factor := 2 * _delay * _rate
	if factor > 10000 {
		myLog.Fatalf("keep factor:%v = 2 * _delay:%v * _rate:%v under 1000!", factor, _delay, _rate)
	}

	coef := math.Pow(math.E, factor)

	myLog.Printf("_delay:%v _rate:%v  level:%v factor:%v coef:%v\n\n", _delay, _rate, _level, factor, coef)

	sum := 0.0

	outLen := 10
	kQueue := make([]float64, 0)

	end := 1000

	k := -1

	for kk := 1; kk < end; kk++ {
		sum = coef

		sigma := 1.0
		for j := 1; j <= kk; j++ {
			n := 1.0
			for jj := 1; jj <= j; jj++ {
				n *= factor / float64(jj)
			}
			sigma += n
		}
		sum -= sigma
		sum /= coef

		if k < 0 {
			if sum < _level {
				for i := 0; i < len(kQueue); i++ {
					leftBound := outLen
					if kk <= leftBound {
						leftBound = kk - 1
					}
					myLog.Printf("kk=%v sum=%v", kk-(leftBound-i), kQueue[i])
				}
				myLog.Printf("\n[MIN]kk=%v sum=%v\n\n", kk, sum)
				k = kk
				end = kk + outLen + 1
			}

			kQueue = append(kQueue, sum)
			if kk > outLen {
				kQueue = kQueue[1:]
			}

		} else {
			myLog.Printf("kk=%v sum=%v", kk, sum)
		}
	}
}
